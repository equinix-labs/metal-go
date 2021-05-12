# \SelfServiceReservationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectIdSelfServiceReservationsGet**](SelfServiceReservationsApi.md#ProjectsProjectIdSelfServiceReservationsGet) | **Get** /projects/{project_id}/self-service/reservations | Retrieve all reservations
[**ProjectsProjectIdSelfServiceReservationsIdGet**](SelfServiceReservationsApi.md#ProjectsProjectIdSelfServiceReservationsIdGet) | **Get** /projects/{project_id}/self-service/reservations/{id} | Retrieve a reservation
[**ProjectsProjectIdSelfServiceReservationsPost**](SelfServiceReservationsApi.md#ProjectsProjectIdSelfServiceReservationsPost) | **Post** /projects/{project_id}/self-service/reservations | Create a reservation



## ProjectsProjectIdSelfServiceReservationsGet

> SelfServiceReservationList ProjectsProjectIdSelfServiceReservationsGet(ctx, projectId).Page(page).PerPage(perPage).Execute()

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
    projectId := TODO // string | Project UUID
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsGet(context.Background(), projectId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsProjectIdSelfServiceReservationsGet`: SelfServiceReservationList
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdSelfServiceReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**SelfServiceReservationList**](SelfServiceReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdSelfServiceReservationsIdGet

> SelfServiceReservationResponse ProjectsProjectIdSelfServiceReservationsIdGet(ctx, id, projectId).Execute()

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
    id := TODO // string | Reservation short_id
    projectId := TODO // string | Project UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsIdGet(context.Background(), id, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsProjectIdSelfServiceReservationsIdGet`: SelfServiceReservationResponse
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Reservation short_id | 
**projectId** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdSelfServiceReservationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SelfServiceReservationResponse**](SelfServiceReservationResponse.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdSelfServiceReservationsPost

> SelfServiceReservationResponse ProjectsProjectIdSelfServiceReservationsPost(ctx, projectId).Reservation(reservation).Execute()

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
    projectId := TODO // string | Project UUID
    reservation := *openapiclient.NewCreateSelfServiceReservationRequest() // CreateSelfServiceReservationRequest | reservation to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsPost(context.Background(), projectId).Reservation(reservation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsProjectIdSelfServiceReservationsPost`: SelfServiceReservationResponse
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceReservationsApi.ProjectsProjectIdSelfServiceReservationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdSelfServiceReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reservation** | [**CreateSelfServiceReservationRequest**](CreateSelfServiceReservationRequest.md) | reservation to create | 

### Return type

[**SelfServiceReservationResponse**](SelfServiceReservationResponse.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

