#!/bin/sh
curl -o equinix-metal.swagger.json https://api.equinix.com/metal/v1/api-docs
for a in *patch; do patch -p0 < "$a"; done

SWAGGER="docker run --rm -it --env GOPATH=/go -v `pwd`:/go/src -w /go/src quay.io/goswagger/swagger"
# binary available at https://github.com/go-swagger/go-swagger/releases/latest
# we use Docker to keep builds consistent between users
$SWAGGER generate client \
	--model-package types \
	--additional-initialism bgp \
	--additional-initialism vpn \
	--additional-initialism vlan \
	--additional-initialism vlans \
	-f equinix-metal.swagger.json

