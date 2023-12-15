# Golang client for Equinix Metal

[![Deprecated](https://img.shields.io/badge/Stability-Deprecated-red.svg)](https://github.com/equinix-labs/equinix-labs/blob/main/uniform-standards.md)

> **[Deprecated](https://github.com/equinix-labs/equinix-labs/blob/main/deprecated-statement.md)** This repository is Deprecated, meaning that this software is only supported or maintained by Equinix Metal and its community in a case-by-case basis.  It has been replaced by [`equinix-sdk-go`](https://github.com/equinix/equinix-sdk-go), which is generated using a similar process to this SDK and is intended to support Equinix services beyond only Metal.  Version 0.30.0 of the new SDK is a drop-in replacement for version 0.29.0 of this SDK.

## Contents

- `Makefile` includes tasks to fetch the API spec, apply patches, and generate a client
- `spec/oas3.fetched` a directory of the latest fetched OpenAPI spec
- `spec/oas3.patched` a directory of the latest patched OpenAPI spec
- `patches/spec.fetched.json/*.patch` patch files to apply against the fetched OpenAPI spec
- `patches/post/*.patch` patch files to apply against the generated Go code
- `examples/` hand crafted examples to demonstrate usage
- `metal/v1/` generated Go client

## Generated Client

See [API.md](API.md)

## Build

To build the client, run `make`.

## Examples

You can see usage of the generated code in the [`examples` directory](https://github.com/equinix-labs/metal-go/tree/main/examples). In order to try, export `METAL_AUTH_TOKEN` token and execute the code, e.g.

## Patches

1. Make changes in ``spec/oas3.patched/`` dir.
2. Create a patch file in metal-go:
   ```
   git diff spec/oas3.patched/ > ../patches/spec.fetched.json/<patchfilename>
   cd ..
   ```
3. ``patchfilename`` should be in format: ``<patch_index>-<short_patch_decription_or_identifier>.patch``
4. Run ``make`` to reapply the changes to the fetched OpenAPI spec and confirm that the patched spec includes the expected changes


