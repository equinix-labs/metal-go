# 0.7.0 (22 Mar, 2023)

FEATURES:

- Download split spec from https://api.equinix.com/metal/v1/api-docs for more straightforward patching [#69](https://github.com/equinix-labs/metal-go/pull/69)
- Enable the use of `additionalProperties` to access fields that are returned by the API but are not defined in the API spec [#85](https://github.com/equinix-labs/metal-go/pull/85)
- Update the default `User-Agent` for improved usage tracking [#86](https://github.com/equinix-labs/metal-go/pull/86)
- Regenerated code from latest API spec [#87](https://github.com/equinix-labs/metal-go/pull/87)

BUG FIXES:

- Patch `Device` schema so that `project` and `hardware_reservation` properties are defined with the correct schema components [#80](https://github.com/equinix-labs/metal-go/pull/80)
