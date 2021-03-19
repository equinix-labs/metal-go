# \InternetGatewaysApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternetGateway**](InternetGatewaysApi.md#CreateInternetGateway) | **Post** /virtual-networks/{id}/internet-gateways | Create an internet gateway



## CreateInternetGateway

> InternetGateway CreateInternetGateway(ctx, id).Length(length).Execute()

Create an internet gateway



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
    id := TODO // string | Virtual Network UUID
    length := "length_example" // string | IP Reservation length

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetGatewaysApi.CreateInternetGateway(context.Background(), id).Length(length).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewaysApi.CreateInternetGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInternetGateway`: InternetGateway
    fmt.Fprintf(os.Stdout, "Response from `InternetGatewaysApi.CreateInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **length** | **string** | IP Reservation length | 

### Return type

[**InternetGateway**](InternetGateway.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

