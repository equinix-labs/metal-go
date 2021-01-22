#!/bin/sh
curl -o equinix-metal.swagger.json https://api.equinix.com/metal/v1/api-docs
for a in *patch; do patch -p0 < "$a"; done

# https://github.com/go-swagger/go-swagger/releases/latest
swagger generate client \
	--model-package types \
	--additional-initialism bgp \
	--additional-initialism vpn \
	--additional-initialism vlan \
	--additional-initialism vlans \
	-f equinix-metal.swagger.json

