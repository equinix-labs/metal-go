# \MetrosApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindMetros**](MetrosApi.md#FindMetros) | **Get** /locations/metros | Retrieve all metros
[**GetMetro**](MetrosApi.md#GetMetro) | **Get** /locations/metros/{id} | Retrieve a specific Metro&#39;s details



## FindMetros

> FindMetros200Response FindMetros(ctx).Execute()

Retrieve all metros



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetrosApi.FindMetros(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetrosApi.FindMetros``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetros`: FindMetros200Response
    fmt.Fprintf(os.Stdout, "Response from `MetrosApi.FindMetros`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindMetrosRequest struct via the builder pattern


### Return type

[**FindMetros200Response**](FindMetros200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetro

> GetInterconnection200ResponseMetroAllOf GetMetro(ctx, id).Execute()

Retrieve a specific Metro's details



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metro UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetrosApi.GetMetro(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetrosApi.GetMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetro`: GetInterconnection200ResponseMetroAllOf
    fmt.Fprintf(os.Stdout, "Response from `MetrosApi.GetMetro`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metro UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInterconnection200ResponseMetroAllOf**](GetInterconnection200ResponseMetroAllOf.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

