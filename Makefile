.PHONY: all gen patch fetch

# https://github.com/go-swagger/go-swagger/releases/latest
SWAGGER=docker run --rm -it --env GOPATH=/go -v $(CURDIR):/go/src -w /go/src quay.io/goswagger/swagger

all: fetch patch gen

fetch:
	curl -o equinix-metal.swagger.json https://api.equinix.com/metal/v1/api-docs

patch:
	for a in patches/*patch; do patch -p0 < "$$a"; done

gen:
	${SWAGGER} generate client \
		--model-package types \
		--additional-initialism bgp \
		--additional-initialism vpn \
		--additional-initialism vlan \
		--additional-initialism vlans \
		-f equinix-metal.swagger.json
