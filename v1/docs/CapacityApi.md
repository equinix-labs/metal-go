# \CapacityApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapacityForFacility**](CapacityApi.md#CheckCapacityForFacility) | **Post** /capacity | Check capacity
[**CheckCapacityForMetro**](CapacityApi.md#CheckCapacityForMetro) | **Post** /capacity/metros | Check capacity for a metro
[**FindCapacityForFacility**](CapacityApi.md#FindCapacityForFacility) | **Get** /capacity | View capacity
[**FindCapacityForMetro**](CapacityApi.md#FindCapacityForMetro) | **Get** /capacity/metros | View capacity for metros



## CheckCapacityForFacility

> CapacityCheckPerFacilityList CheckCapacityForFacility(ctx).Facility(facility).Execute()

Check capacity



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
    facility := *openapiclient.NewCapacityInput() // CapacityInput | Facility to check capacity in

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityApi.CheckCapacityForFacility(context.Background()).Facility(facility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.CheckCapacityForFacility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCapacityForFacility`: CapacityCheckPerFacilityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.CheckCapacityForFacility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCapacityForFacilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | [**CapacityInput**](CapacityInput.md) | Facility to check capacity in | 

### Return type

[**CapacityCheckPerFacilityList**](CapacityCheckPerFacilityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckCapacityForMetro

> CapacityCheckPerMetroList CheckCapacityForMetro(ctx).Servers(servers).Execute()

Check capacity for a metro



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
    servers := *openapiclient.NewCapacityPerMetroInput() // CapacityPerMetroInput | Metro to check capacity in

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityApi.CheckCapacityForMetro(context.Background()).Servers(servers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.CheckCapacityForMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCapacityForMetro`: CapacityCheckPerMetroList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.CheckCapacityForMetro`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCapacityForMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servers** | [**CapacityPerMetroInput**](CapacityPerMetroInput.md) | Metro to check capacity in | 

### Return type

[**CapacityCheckPerMetroList**](CapacityCheckPerMetroList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCapacityForFacility

> CapacityList FindCapacityForFacility(ctx).Execute()

View capacity



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
    resp, r, err := api_client.CapacityApi.FindCapacityForFacility(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindCapacityForFacility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCapacityForFacility`: CapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindCapacityForFacility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCapacityForFacilityRequest struct via the builder pattern


### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCapacityForMetro

> MetroCapacityList FindCapacityForMetro(ctx).Execute()

View capacity for metros



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
    resp, r, err := api_client.CapacityApi.FindCapacityForMetro(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindCapacityForMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCapacityForMetro`: MetroCapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindCapacityForMetro`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCapacityForMetroRequest struct via the builder pattern


### Return type

[**MetroCapacityList**](MetroCapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

