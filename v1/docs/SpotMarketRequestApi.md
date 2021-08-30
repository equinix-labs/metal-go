# \SpotMarketRequestApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpotMarketRequest**](SpotMarketRequestApi.md#DeleteSpotMarketRequest) | **Delete** /spot-market-requests/{id} | Delete the spot market request
[**FindSpotMarketRequestById**](SpotMarketRequestApi.md#FindSpotMarketRequestById) | **Get** /spot-market-requests/{id} | Retrieve a spot market request



## DeleteSpotMarketRequest

> DeleteSpotMarketRequest(ctx, id).ForceTermination(forceTermination).Execute()

Delete the spot market request



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
    id := TODO // string | SpotMarketRequest UUID
    forceTermination := true // bool | Terminate associated spot instances (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SpotMarketRequestApi.DeleteSpotMarketRequest(context.Background(), id).ForceTermination(forceTermination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketRequestApi.DeleteSpotMarketRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | SpotMarketRequest UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpotMarketRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceTermination** | **bool** | Terminate associated spot instances | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketRequestById

> SpotMarketRequest FindSpotMarketRequestById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a spot market request



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
    id := TODO // string | SpotMarketRequest UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SpotMarketRequestApi.FindSpotMarketRequestById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketRequestApi.FindSpotMarketRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketRequestById`: SpotMarketRequest
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketRequestApi.FindSpotMarketRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | SpotMarketRequest UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**SpotMarketRequest**](SpotMarketRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

