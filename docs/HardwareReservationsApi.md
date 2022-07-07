# \HardwareReservationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindHardwareReservationById**](HardwareReservationsApi.md#FindHardwareReservationById) | **Get** /hardware-reservations/{id} | Retrieve a hardware reservation
[**FindProjectHardwareReservations**](HardwareReservationsApi.md#FindProjectHardwareReservations) | **Get** /projects/{id}/hardware-reservations | Retrieve all hardware reservations for a given project
[**MoveHardwareReservation**](HardwareReservationsApi.md#MoveHardwareReservation) | **Post** /hardware-reservations/{id}/move | Move a hardware reservation



## FindHardwareReservationById

> FindDeviceById200Response FindHardwareReservationById(ctx, id).Include(include).Exclude(exclude).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | HardwareReservation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.FindHardwareReservationById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.FindHardwareReservationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindHardwareReservationById`: FindDeviceById200Response
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.FindHardwareReservationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HardwareReservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindHardwareReservationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindDeviceById200Response**](FindDeviceById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectHardwareReservations

> FindProjectHardwareReservations200Response FindProjectHardwareReservations(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all hardware reservations for a given project



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.FindProjectHardwareReservations(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.FindProjectHardwareReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectHardwareReservations`: FindProjectHardwareReservations200Response
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.FindProjectHardwareReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectHardwareReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindProjectHardwareReservations200Response**](FindProjectHardwareReservations200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveHardwareReservation

> MoveHardwareReservation201Response MoveHardwareReservation(ctx, id).Body(body).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Hardware Reservation UUID
    body := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.MoveHardwareReservation(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.MoveHardwareReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveHardwareReservation`: MoveHardwareReservation201Response
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.MoveHardwareReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hardware Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveHardwareReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Project UUID | 

### Return type

[**MoveHardwareReservation201Response**](MoveHardwareReservation201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

