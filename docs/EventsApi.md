# \EventsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindConnectionEvents**](EventsApi.md#FindConnectionEvents) | **Get** /connections/{connection_id}/events | Retrieve connection events
[**FindConnectionPortEvents**](EventsApi.md#FindConnectionPortEvents) | **Get** /connections/{connection_id}/ports/{id}/events | Retrieve connection port events
[**FindDeviceEvents**](EventsApi.md#FindDeviceEvents) | **Get** /devices/{id}/events | Retrieve device&#39;s events
[**FindEventById**](EventsApi.md#FindEventById) | **Get** /events/{id} | Retrieve an event
[**FindEvents**](EventsApi.md#FindEvents) | **Get** /events | Retrieve current user&#39;s events
[**FindOrganizationEvents**](EventsApi.md#FindOrganizationEvents) | **Get** /organizations/{id}/events | Retrieve organization&#39;s events
[**FindProjectEvents**](EventsApi.md#FindProjectEvents) | **Get** /projects/{id}/events | Retrieve project&#39;s events
[**FindVirtualCircuitEvents**](EventsApi.md#FindVirtualCircuitEvents) | **Get** /virtual-circuit/{id}/events | Retrieve connection events



## FindConnectionEvents

> FindConnectionEvents200Response FindConnectionEvents(ctx, connectionId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve connection events



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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindConnectionEvents(context.Background(), connectionId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindConnectionEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindConnectionEvents`: FindConnectionEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindConnectionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindConnectionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindConnectionEvents200Response**](FindConnectionEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindConnectionPortEvents

> FindConnectionEvents200Response FindConnectionPortEvents(ctx, connectionId, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve connection port events



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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection Port UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindConnectionPortEvents(context.Background(), connectionId, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindConnectionPortEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindConnectionPortEvents`: FindConnectionEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindConnectionPortEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 
**id** | **string** | Connection Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindConnectionPortEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindConnectionEvents200Response**](FindConnectionEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDeviceEvents

> FindDeviceEvents200Response FindDeviceEvents(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve device's events



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Device UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindDeviceEvents(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindDeviceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceEvents`: FindDeviceEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindDeviceEvents200Response**](FindDeviceEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEventById

> FindConnectionEvents200Response FindEventById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve an event



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Event UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindEventById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindEventById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindEventById`: FindConnectionEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Event UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindConnectionEvents200Response**](FindConnectionEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEvents

> FindDeviceEvents200Response FindEvents(ctx).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve current user's events



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
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindEvents(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindEvents`: FindDeviceEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindDeviceEvents200Response**](FindDeviceEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOrganizationEvents

> FindDeviceEvents200Response FindOrganizationEvents(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve organization's events



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindOrganizationEvents(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindOrganizationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOrganizationEvents`: FindDeviceEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindOrganizationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOrganizationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindDeviceEvents200Response**](FindDeviceEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectEvents

> FindDeviceEvents200Response FindProjectEvents(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve project's events



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
    resp, r, err := apiClient.EventsApi.FindProjectEvents(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindProjectEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectEvents`: FindDeviceEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindProjectEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindDeviceEvents200Response**](FindDeviceEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVirtualCircuitEvents

> FindConnectionEvents200Response FindVirtualCircuitEvents(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve connection events



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.FindVirtualCircuitEvents(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.FindVirtualCircuitEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVirtualCircuitEvents`: FindConnectionEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.FindVirtualCircuitEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVirtualCircuitEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindConnectionEvents200Response**](FindConnectionEvents200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

