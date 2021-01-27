.PHONY: all gen patch fetch

all: fetch patch gen

fetch:
	curl -o equinix-metal.swagger.json https://api.equinix.com/metal/v1/api-docs

patch:
	for a in *patch; do patch -p0 < "$$a"; done

gen:
	# https://github.com/go-swagger/go-swagger/releases/latest
	swagger generate client \
		--model-package types \
		--additional-initialism bgp \
		--additional-initialism vpn \
		--additional-initialism vlan \
		--additional-initialism vlans \
		-f equinix-metal.swagger.json
