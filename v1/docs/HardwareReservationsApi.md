# \HardwareReservationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindHardwareReservationById**](HardwareReservationsApi.md#FindHardwareReservationById) | **Get** /hardware-reservations/{id} | Retrieve a hardware reservation
[**MoveHardwareReservation**](HardwareReservationsApi.md#MoveHardwareReservation) | **Post** /hardware-reservations/{id}/move | Move a hardware reservation



## FindHardwareReservationById

> Device FindHardwareReservationById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a hardware reservation



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
    id := TODO // string | HardwareReservation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HardwareReservationsApi.FindHardwareReservationById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.FindHardwareReservationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindHardwareReservationById`: Device
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.FindHardwareReservationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | HardwareReservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindHardwareReservationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveHardwareReservation

> HardwareReservation MoveHardwareReservation(ctx, id).ProjectId(projectId).Execute()

Move a hardware reservation



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
    id := TODO // string | Hardware Reservation UUID
    projectId := string(38400000-8cf0-11bd-b23e-10b96e4ef00d) // string | Project UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HardwareReservationsApi.MoveHardwareReservation(context.Background(), id).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.MoveHardwareReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveHardwareReservation`: HardwareReservation
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.MoveHardwareReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Hardware Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveHardwareReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | Project UUID | 

### Return type

[**HardwareReservation**](HardwareReservation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

