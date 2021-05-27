# \GlobalBgpRangesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindGlobalBgpRanges**](GlobalBgpRangesApi.md#FindGlobalBgpRanges) | **Get** /projects/{id}/global-bgp-ranges | Retrieve all global bgp ranges



## FindGlobalBgpRanges

> GlobalBgpRangeList FindGlobalBgpRanges(ctx, id).Execute()

Retrieve all global bgp ranges



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
    resp, r, err := api_client.GlobalBgpRangesApi.FindGlobalBgpRanges(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalBgpRangesApi.FindGlobalBgpRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindGlobalBgpRanges`: GlobalBgpRangeList
    fmt.Fprintf(os.Stdout, "Response from `GlobalBgpRangesApi.FindGlobalBgpRanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindGlobalBgpRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalBgpRangeList**](GlobalBgpRangeList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

