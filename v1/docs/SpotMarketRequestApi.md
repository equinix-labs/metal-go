# \SpotMarketRequestApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpotMarketRequest**](SpotMarketRequestApi.md#CreateSpotMarketRequest) | **Post** /projects/{id}/spot-market-requests | Create a spot market request
[**DeleteSpotMarketRequest**](SpotMarketRequestApi.md#DeleteSpotMarketRequest) | **Delete** /spot-market-requests/{id} | Delete the spot market request
[**FindSpotMarketRequestById**](SpotMarketRequestApi.md#FindSpotMarketRequestById) | **Get** /spot-market-requests/{id} | Retrieve a spot market request
[**ListSpotMarketRequests**](SpotMarketRequestApi.md#ListSpotMarketRequests) | **Get** /projects/{id}/spot-market-requests | List spot market requests



## CreateSpotMarketRequest

> SpotMarketRequest CreateSpotMarketRequest(ctx, id).SpotMarketRequest(spotMarketRequest).Execute()

Create a spot market request



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
    id := TODO // string | Project UUID
    spotMarketRequest := *openapiclient.NewSpotMarketRequestCreateInput() // SpotMarketRequestCreateInput | Spot Market Request to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SpotMarketRequestApi.CreateSpotMarketRequest(context.Background(), id).SpotMarketRequest(spotMarketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketRequestApi.CreateSpotMarketRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpotMarketRequest`: SpotMarketRequest
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketRequestApi.CreateSpotMarketRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpotMarketRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spotMarketRequest** | [**SpotMarketRequestCreateInput**](SpotMarketRequestCreateInput.md) | Spot Market Request to create | 

### Return type

[**SpotMarketRequest**](SpotMarketRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListSpotMarketRequests

> SpotMarketRequestList ListSpotMarketRequests(ctx, id).Execute()

List spot market requests



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
    id := TODO // string | Project UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SpotMarketRequestApi.ListSpotMarketRequests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketRequestApi.ListSpotMarketRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpotMarketRequests`: SpotMarketRequestList
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketRequestApi.ListSpotMarketRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpotMarketRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpotMarketRequestList**](SpotMarketRequestList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

