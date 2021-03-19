# \VPNApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindCurrentUserVpnConfig**](VPNApi.md#FindCurrentUserVpnConfig) | **Get** /user/vpn | Retrieve the client vpn config for current user
[**TurnOffCurrentUserVpn**](VPNApi.md#TurnOffCurrentUserVpn) | **Delete** /user/vpn | Turn off vpn for the current user
[**TurnOnCurrentUserVpn**](VPNApi.md#TurnOnCurrentUserVpn) | **Post** /user/vpn | Turn on vpn for the current user



## FindCurrentUserVpnConfig

> VPNConfig FindCurrentUserVpnConfig(ctx).Code(code).Execute()

Retrieve the client vpn config for current user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | Facility code

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VPNApi.FindCurrentUserVpnConfig(context.Background()).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.FindCurrentUserVpnConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCurrentUserVpnConfig`: VPNConfig
    fmt.Fprintf(os.Stdout, "Response from `VPNApi.FindCurrentUserVpnConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCurrentUserVpnConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Facility code | 

### Return type

[**VPNConfig**](VPNConfig.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TurnOffCurrentUserVpn

> TurnOffCurrentUserVpn(ctx).Execute()

Turn off vpn for the current user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VPNApi.TurnOffCurrentUserVpn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.TurnOffCurrentUserVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTurnOffCurrentUserVpnRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TurnOnCurrentUserVpn

> TurnOnCurrentUserVpn(ctx).Execute()

Turn on vpn for the current user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VPNApi.TurnOnCurrentUserVpn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.TurnOnCurrentUserVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTurnOnCurrentUserVpnRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

