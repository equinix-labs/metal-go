.PHONY: all gen patch fetch

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

# https://github.com/OpenAPITools/openapi-generator-cli

SPEC_URL:=https://api.equinix.com/metal/v1/api-docs

SPEC_FETCHED_FILE:=spec.fetched.json
SPEC_PATCHED_FILE:=spec.patched.json
IMAGE=openapitools/openapi-generator-cli
GIT_ORG=equinix-labs
GIT_REPO=metal-go
PACKAGE_PREFIX=metal
PACKAGE_MAJOR=v1

# Pull in custom-built generator jar so we can use unmerged bugfixes
# Custom generator is built from https://github.com/ctreatma/openapi-generator/tree/local-generator-testing
SWAGGER=docker run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local -v $(CURDIR)/lib/openapi-generator-cli.jar:/opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar ${IMAGE}
GOLANGCI_LINT=golangci-lint

all: pull fetch patch clean gen mod docs move-other patch-post fmt test stage

pull:
	docker pull ${IMAGE}

fetch:
	curl ${SPEC_URL} | jq . > ${SPEC_FETCHED_FILE}

fix-tags:
	- jq '. | select(((.paths[][].tags| type=="array"), length) > 1).paths[][].tags |= [.[0]]' ${SPEC_FETCHED_FILE} | diff -d -U6 ${SPEC_FETCHED_FILE} - >  patches/01-tag-from-last-in-path.patch

patch:
	# patch is idempotent, always starting with the fetched
	# fetched file to create the patched file.
	cp ${SPEC_FETCHED_FILE} ${SPEC_PATCHED_FILE}
	ARGS="-o ${SPEC_PATCHED_FILE} ${SPEC_FETCHED_FILE}"; \
	for diff in $(shell find patches/${SPEC_FETCHED_FILE} -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t $$ARGS $$diff; \
		ARGS=${SPEC_PATCHED_FILE}; \
	done

patch-post:
	# patch is idempotent, always starting with the generated files
	for diff in $(shell find patches/post -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done


clean:
	rm -rf $(PACKAGE_PREFIX)

gen:
	${SWAGGER} generate -g go \
		--package-name ${PACKAGE_MAJOR} \
		-p isGoSubmodule=true \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO}/${PACKAGE_PREFIX} \
		-o /local/${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i /local/${SPEC_PATCHED_FILE}
    # generated code is missing some types; hack 'em in
	cat missing_types.go.part >> metal/v1/utils.go

validate:
	${SWAGGER} validate \
		--recommend \
		-i /local/${SPEC_PATCHED_FILE}

mod:
	rm -f go.mod go.sum ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.mod ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.sum
	go mod init github.com/${GIT_ORG}/${GIT_REPO}
	go mod tidy

test:
	go test -v ./...

clean-docs:
	rm -rf docs API.md

move-docs:
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/README.md API.md
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/docs .

docs: clean-docs move-docs

move-other:
	rm -rf api .travis.yml git_push.sh
	rm -f ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/.travis.yml
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/api .
	rm ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/git_push.sh

# https://github.com/OpenAPITools/openapi-generator/issues/741#issuecomment-569791780
remove-dupe-requests: ## Removes duplicate Request structs from the generated code
	@for struct in $$(grep -h 'type .\{1,\} struct' $(PACKAGE_PREFIX)/$(PACKAGE_MAJOR)/*.go | grep Request  | sort | uniq -c | grep -v '^      1' | awk '{print $$3}'); do \
	  for f in $$(/bin/ls $(PACKAGE_PREFIX)/$(PACKAGE_MAJOR)/*.go); do \
	    if grep -qF "type $${struct} struct" "$${f}"; then \
	      if eval "test -z \$${$${struct}}"; then \
	        echo "skipping first appearance of $${struct} in file $${f}"; \
	        eval "export $${struct}=1"; \
	      else \
	        echo "removing dupe $${struct} from file $${f}"; \
	        tr '\n' '\r' <"$${f}" | sed 's~// '"$${struct}"'.\{1,\}type '"$${struct}"' struct {[^}]\{1,\}}~~' | tr '\r' '\n' >"$${f}.tmp"; \
	        mv -f "$${f}.tmp" "$${f}"; \
	      fi; \
	    fi \
	  done \
	done
lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w $(PACKAGE_PREFIX)

stage:
	test -d .git && git add --intent-to-add API.md docs ${PACKAGE_PREFIX} go.mod go.sum
