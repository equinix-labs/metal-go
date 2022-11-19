# \PaymentMethodsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePaymentMethod**](PaymentMethodsApi.md#DeletePaymentMethod) | **Delete** /payment-methods/{id} | Delete the payment method
[**FindPaymentMethodById**](PaymentMethodsApi.md#FindPaymentMethodById) | **Get** /payment-methods/{id} | Retrieve a payment method
[**UpdatePaymentMethod**](PaymentMethodsApi.md#UpdatePaymentMethod) | **Put** /payment-methods/{id} | Update the payment method



## DeletePaymentMethod

> DeletePaymentMethod(ctx, id).Execute()

Delete the payment method



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment Method UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsApi.DeletePaymentMethod(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.DeletePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment Method UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## FindPaymentMethodById

> FindOrganizationPaymentMethods200ResponsePaymentMethodsInner FindPaymentMethodById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a payment method



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment Method UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsApi.FindPaymentMethodById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.FindPaymentMethodById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPaymentMethodById`: FindOrganizationPaymentMethods200ResponsePaymentMethodsInner
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.FindPaymentMethodById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment Method UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPaymentMethodByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindOrganizationPaymentMethods200ResponsePaymentMethodsInner**](FindOrganizationPaymentMethods200ResponsePaymentMethodsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentMethod

> FindOrganizationPaymentMethods200ResponsePaymentMethodsInner UpdatePaymentMethod(ctx, id).UpdatePaymentMethodRequest(updatePaymentMethodRequest).Execute()

Update the payment method



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment Method UUID
    updatePaymentMethodRequest := *openapiclient.NewUpdatePaymentMethodRequest() // UpdatePaymentMethodRequest | Payment Method to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsApi.UpdatePaymentMethod(context.Background(), id).UpdatePaymentMethodRequest(updatePaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.UpdatePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentMethod`: FindOrganizationPaymentMethods200ResponsePaymentMethodsInner
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.UpdatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment Method UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaymentMethodRequest** | [**UpdatePaymentMethodRequest**](UpdatePaymentMethodRequest.md) | Payment Method to update | 

### Return type

[**FindOrganizationPaymentMethods200ResponsePaymentMethodsInner**](FindOrganizationPaymentMethods200ResponsePaymentMethodsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

