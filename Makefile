.PHONY: all gen patch fetch

# https://github.com/go-swagger/go-swagger/releases/latest
SPEC_URL:=https://api.equinix.com/metal/v1/api-docs
SPEC_FETCHED_FILE:=equinix-metal.fetched.json
SPEC_PATCHED_FILE:=equinix-metal.patched.json
IMAGE=quay.io/goswagger/swagger

SWAGGER=docker run --rm --env GOPATH=/go -v $(CURDIR):/go/src -w /go/src ${IMAGE}

all: pull fetch patch gen

pull:
	docker pull ${IMAGE}

fetch:
	curl -o ${SPEC_FETCHED_FILE} ${SPEC_URL}

patch:
	# patch is idempotent, always starting with the fetched
	# fetched file to create the patched file.
	ARGS="-o ${SPEC_PATCHED_FILE} ${SPEC_FETCHED_FILE}"; \
	for diff in $(shell find patches/*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t $$ARGS $$diff; \
		ARGS=${SPEC_PATCHED_FILE}; \
	done

gen:
	${SWAGGER} generate client \
		--model-package types \
		--additional-initialism bgp \
		--additional-initialism vpn \
		--additional-initialism vlan \
		--additional-initialism vlans \
		-f ${SPEC_PATCHED_FILE}
