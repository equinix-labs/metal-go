# \DevicesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpSession**](DevicesApi.md#CreateBgpSession) | **Post** /devices/{id}/bgp/sessions | Create a BGP session
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /projects/{id}/devices | Create a device
[**CreateIPAssignment**](DevicesApi.md#CreateIPAssignment) | **Post** /devices/{id}/ips | Create an ip assignment
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{id} | Delete the device
[**FindBgpSessions**](DevicesApi.md#FindBgpSessions) | **Get** /devices/{id}/bgp/sessions | Retrieve all BGP sessions
[**FindDeviceById**](DevicesApi.md#FindDeviceById) | **Get** /devices/{id} | Retrieve a device
[**FindDeviceCustomdata**](DevicesApi.md#FindDeviceCustomdata) | **Get** /devices/{id}/customdata | Retrieve the custom metadata of an instance
[**FindIPAssignmentCustomdata**](DevicesApi.md#FindIPAssignmentCustomdata) | **Get** /devices/{instance_id}/ips/{id}/customdata | Retrieve the custom metadata of an IP Assignment
[**FindIPAssignments**](DevicesApi.md#FindIPAssignments) | **Get** /devices/{id}/ips | Retrieve all ip assignments
[**FindInstanceBandwidth**](DevicesApi.md#FindInstanceBandwidth) | **Get** /devices/{id}/bandwidth | Retrieve an instance bandwidth
[**FindOrganizationDevices**](DevicesApi.md#FindOrganizationDevices) | **Get** /organizations/{id}/devices | Retrieve all devices of an organization
[**FindProjectDevices**](DevicesApi.md#FindProjectDevices) | **Get** /projects/{id}/devices | Retrieve all devices of a project
[**FindTraffic**](DevicesApi.md#FindTraffic) | **Get** /devices/{id}/traffic | Retrieve device traffic
[**GetBgpNeighborData**](DevicesApi.md#GetBgpNeighborData) | **Get** /devices/{id}/bgp/neighbors | Retrieve BGP neighbor data for this device
[**PerformAction**](DevicesApi.md#PerformAction) | **Post** /devices/{id}/actions | Perform an action
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{id} | Update the device



## CreateBgpSession

> FindBgpSessionById200Response CreateBgpSession(ctx, id).Body(body).Execute()

Create a BGP session



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
    body := *openapiclient.NewCreateBgpSessionRequest() // CreateBgpSessionRequest | BGP session to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateBgpSession(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateBgpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpSession`: FindBgpSessionById200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateBgpSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateBgpSessionRequest**](CreateBgpSessionRequest.md) | BGP session to create | 

### Return type

[**FindBgpSessionById200Response**](FindBgpSessionById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> FindDeviceById200Response CreateDevice(ctx, id).Body(body).Execute()

Create a device



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
    body := *openapiclient.NewCreateDeviceRequest("OperatingSystem_example", "c3.large.x86") // CreateDeviceRequest | Device to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDevice(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: FindDeviceById200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateDeviceRequest**](CreateDeviceRequest.md) | Device to create | 

### Return type

[**FindDeviceById200Response**](FindDeviceById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIPAssignment

> FindDeviceById200ResponseIpAddressesInner CreateIPAssignment(ctx, id).Body(body).Execute()

Create an ip assignment



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
    body := *openapiclient.NewCreateIPAssignmentRequest("Address_example") // CreateIPAssignmentRequest | IPAssignment to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateIPAssignment(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateIPAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPAssignment`: FindDeviceById200ResponseIpAddressesInner
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateIPAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateIPAssignmentRequest**](CreateIPAssignmentRequest.md) | IPAssignment to create | 

### Return type

[**FindDeviceById200ResponseIpAddressesInner**](FindDeviceById200ResponseIpAddressesInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, id).ForceDelete(forceDelete).Execute()

Delete the device



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
    forceDelete := true // bool | Force the deletion of the device, by detaching any storage volume still active. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteDevice(context.Background(), id).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Force the deletion of the device, by detaching any storage volume still active. | 

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


## FindBgpSessions

> FindBgpSessions200Response FindBgpSessions(ctx, id).Execute()

Retrieve all BGP sessions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindBgpSessions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindBgpSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBgpSessions`: FindBgpSessions200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindBgpSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBgpSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindBgpSessions200Response**](FindBgpSessions200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDeviceById

> FindDeviceById200Response FindDeviceById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a device



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindDeviceById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceById`: FindDeviceById200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindDeviceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceByIdRequest struct via the builder pattern


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


## FindDeviceCustomdata

> FindDeviceCustomdata(ctx, id).Execute()

Retrieve the custom metadata of an instance



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Instance UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindDeviceCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceCustomdataRequest struct via the builder pattern


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


## FindIPAssignmentCustomdata

> FindIPAssignmentCustomdata(ctx, instanceId, id).Execute()

Retrieve the custom metadata of an IP Assignment



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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Instance UUID
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Ip Assignment UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindIPAssignmentCustomdata(context.Background(), instanceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindIPAssignmentCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance UUID | 
**id** | **string** | Ip Assignment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAssignmentCustomdataRequest struct via the builder pattern


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


## FindIPAssignments

> FindIPAssignments200Response FindIPAssignments(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve all ip assignments



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindIPAssignments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindIPAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAssignments`: FindIPAssignments200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindIPAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindIPAssignments200Response**](FindIPAssignments200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInstanceBandwidth

> FindInstanceBandwidth(ctx, id).From(from).Until(until).Execute()

Retrieve an instance bandwidth



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
    from := "from_example" // string | Timestamp from range
    until := "until_example" // string | Timestamp to range

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindInstanceBandwidth(context.Background(), id).From(from).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindInstanceBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInstanceBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Timestamp from range | 
 **until** | **string** | Timestamp to range | 

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


## FindOrganizationDevices

> FindOrganizationDevices200Response FindOrganizationDevices(ctx, id).Facility(facility).Hostname(hostname).Reserved(reserved).Tag(tag).Type_(type_).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all devices of an organization



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
    facility := "facility_example" // string | Filter by device facility (optional)
    hostname := "hostname_example" // string | Filter by partial hostname (optional)
    reserved := true // bool | Filter only reserved instances (optional)
    tag := "tag_example" // string | Filter by device tag (optional)
    type_ := "type__example" // string | Filter by instance type (ondemand,spot,reserved) (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindOrganizationDevices(context.Background(), id).Facility(facility).Hostname(hostname).Reserved(reserved).Tag(tag).Type_(type_).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindOrganizationDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOrganizationDevices`: FindOrganizationDevices200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindOrganizationDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOrganizationDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **facility** | **string** | Filter by device facility | 
 **hostname** | **string** | Filter by partial hostname | 
 **reserved** | **bool** | Filter only reserved instances | 
 **tag** | **string** | Filter by device tag | 
 **type_** | **string** | Filter by instance type (ondemand,spot,reserved) | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindOrganizationDevices200Response**](FindOrganizationDevices200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectDevices

> FindOrganizationDevices200Response FindProjectDevices(ctx, id).Facility(facility).Hostname(hostname).Reserved(reserved).Tag(tag).Type_(type_).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all devices of a project



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
    facility := "facility_example" // string | Filter by device facility (optional)
    hostname := "hostname_example" // string | Filter by partial hostname (optional)
    reserved := true // bool | Filter only reserved instances (optional)
    tag := "tag_example" // string | Filter by device tag (optional)
    type_ := "type__example" // string | Filter by instance type (ondemand,spot,reserved) (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindProjectDevices(context.Background(), id).Facility(facility).Hostname(hostname).Reserved(reserved).Tag(tag).Type_(type_).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindProjectDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectDevices`: FindOrganizationDevices200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindProjectDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **facility** | **string** | Filter by device facility | 
 **hostname** | **string** | Filter by partial hostname | 
 **reserved** | **bool** | Filter only reserved instances | 
 **tag** | **string** | Filter by device tag | 
 **type_** | **string** | Filter by instance type (ondemand,spot,reserved) | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindOrganizationDevices200Response**](FindOrganizationDevices200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTraffic

> FindTraffic(ctx, id).Direction(direction).Body(body).Interval(interval).Bucket(bucket).Execute()

Retrieve device traffic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Device UUID
    direction := "direction_example" // string | Traffic direction
    body := *openapiclient.NewFindTrafficRequest(time.Now(), time.Now()) // FindTrafficRequest | Traffic timeframe
    interval := "interval_example" // string | Traffic interval (optional)
    bucket := "bucket_example" // string | Traffic bucket (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.FindTraffic(context.Background(), id).Direction(direction).Body(body).Interval(interval).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** | Traffic direction | 
 **body** | [**FindTrafficRequest**](FindTrafficRequest.md) | Traffic timeframe | 
 **interval** | **string** | Traffic interval | 
 **bucket** | **string** | Traffic bucket | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpNeighborData

> GetBgpNeighborData200Response GetBgpNeighborData(ctx, id).Execute()

Retrieve BGP neighbor data for this device



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetBgpNeighborData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetBgpNeighborData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBgpNeighborData`: GetBgpNeighborData200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetBgpNeighborData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpNeighborDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBgpNeighborData200Response**](GetBgpNeighborData200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformAction

> PerformAction(ctx, id).Type_(type_).Execute()

Perform an action



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
    type_ := "type__example" // string | Action to perform

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.PerformAction(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.PerformAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Action to perform | 

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


## UpdateDevice

> FindDeviceById200Response UpdateDevice(ctx, id).Body(body).Execute()

Update the device



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
    body := *openapiclient.NewUpdateDeviceRequest() // UpdateDeviceRequest | Facility to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateDevice(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: FindDeviceById200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md) | Facility to update | 

### Return type

[**FindDeviceById200Response**](FindDeviceById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

