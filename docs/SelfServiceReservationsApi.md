# \SelfServiceReservationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSelfServiceReservation**](SelfServiceReservationsApi.md#CreateSelfServiceReservation) | **Post** /projects/{project_id}/self-service/reservations | Create a reservation
[**FindSelfServiceReservation**](SelfServiceReservationsApi.md#FindSelfServiceReservation) | **Get** /projects/{project_id}/self-service/reservations/{id} | Retrieve a reservation
[**FindSelfServiceReservations**](SelfServiceReservationsApi.md#FindSelfServiceReservations) | **Get** /projects/{project_id}/self-service/reservations | Retrieve all reservations



## CreateSelfServiceReservation

> FindSelfServiceReservations200ResponseReservationsInner CreateSelfServiceReservation(ctx, projectId).CreateSelfServiceReservationRequest(createSelfServiceReservationRequest).Execute()

Create a reservation



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    createSelfServiceReservationRequest := *openapiclient.NewCreateSelfServiceReservationRequest() // CreateSelfServiceReservationRequest | reservation to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceReservationsApi.CreateSelfServiceReservation(context.Background(), projectId).CreateSelfServiceReservationRequest(createSelfServiceReservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.CreateSelfServiceReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSelfServiceReservation`: FindSelfServiceReservations200ResponseReservationsInner
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.CreateSelfServiceReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSelfServiceReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSelfServiceReservationRequest** | [**CreateSelfServiceReservationRequest**](CreateSelfServiceReservationRequest.md) | reservation to create | 

### Return type

[**FindSelfServiceReservations200ResponseReservationsInner**](FindSelfServiceReservations200ResponseReservationsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSelfServiceReservation

> FindSelfServiceReservations200ResponseReservationsInner FindSelfServiceReservation(ctx, id, projectId).Execute()

Retrieve a reservation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Reservation short_id
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceReservationsApi.FindSelfServiceReservation(context.Background(), id, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.FindSelfServiceReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSelfServiceReservation`: FindSelfServiceReservations200ResponseReservationsInner
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.FindSelfServiceReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Reservation short_id | 
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSelfServiceReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FindSelfServiceReservations200ResponseReservationsInner**](FindSelfServiceReservations200ResponseReservationsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSelfServiceReservations

> FindSelfServiceReservations200Response FindSelfServiceReservations(ctx, projectId).Page(page).PerPage(perPage).Execute()

Retrieve all reservations



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceReservationsApi.FindSelfServiceReservations(context.Background(), projectId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.FindSelfServiceReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSelfServiceReservations`: FindSelfServiceReservations200Response
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.FindSelfServiceReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSelfServiceReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindSelfServiceReservations200Response**](FindSelfServiceReservations200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

