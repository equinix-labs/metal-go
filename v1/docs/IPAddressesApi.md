# \IPAddressesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPAssignment**](IPAddressesApi.md#CreateIPAssignment) | **Post** /devices/{id}/ips | Create an ip assignment
[**DeleteIPAddress**](IPAddressesApi.md#DeleteIPAddress) | **Delete** /ips/{id} | Unassign an ip address
[**FindIPAddressById**](IPAddressesApi.md#FindIPAddressById) | **Get** /ips/{id} | Retrieve an ip address
[**FindIPAddressCustomdata**](IPAddressesApi.md#FindIPAddressCustomdata) | **Get** /ips/{id}/customdata | Retrieve the custom metadata of an IP Reservation or IP Assignment
[**FindIPAssignments**](IPAddressesApi.md#FindIPAssignments) | **Get** /devices/{id}/ips | Retrieve all ip assignments
[**FindIPAvailabilities**](IPAddressesApi.md#FindIPAvailabilities) | **Get** /ips/{id}/available | Retrieve all available subnets of a particular reservation
[**FindIPReservations**](IPAddressesApi.md#FindIPReservations) | **Get** /projects/{id}/ips | Retrieve all ip reservations
[**RequestIPReservation**](IPAddressesApi.md#RequestIPReservation) | **Post** /projects/{id}/ips | Requesting IP reservations
[**UpdateIPAddress**](IPAddressesApi.md#UpdateIPAddress) | **Patch** /ips/{id} | Update an ip address



## CreateIPAssignment

> IPAssignment CreateIPAssignment(ctx, id).IpAssignment(ipAssignment).Execute()

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
    id := TODO // string | Device UUID
    ipAssignment := *openapiclient.NewIPAssignmentInput("Address_example") // IPAssignmentInput | IPAssignment to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.CreateIPAssignment(context.Background(), id).IpAssignment(ipAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.CreateIPAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPAssignment`: IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.CreateIPAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAssignment** | [**IPAssignmentInput**](IPAssignmentInput.md) | IPAssignment to create | 

### Return type

[**IPAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPAddress

> DeleteIPAddress(ctx, id).Execute()

Unassign an ip address



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
    id := TODO // string | IP Address UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.DeleteIPAddress(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.DeleteIPAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPAddressRequest struct via the builder pattern


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


## FindIPAddressById

> IPAssignment FindIPAddressById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve an ip address



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
    id := TODO // string | IP Address UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.FindIPAddressById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAddressById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAddressById`: IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPAddressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAddressCustomdata

> FindIPAddressCustomdata(ctx, id).Execute()

Retrieve the custom metadata of an IP Reservation or IP Assignment



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
    id := TODO // string | Ip Reservation UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.FindIPAddressCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAddressCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Ip Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAddressCustomdataRequest struct via the builder pattern


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

> IPAssignmentList FindIPAssignments(ctx, id).Include(include).Exclude(exclude).Execute()

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
    id := TODO // string | Device UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.FindIPAssignments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAssignments`: IPAssignmentList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPAssignmentList**](IPAssignmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAvailabilities

> IPAvailabilitiesList FindIPAvailabilities(ctx, id).Cidr(cidr).Execute()

Retrieve all available subnets of a particular reservation



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
    id := TODO // string | IP Reservation UUID
    cidr := "cidr_example" // string | Size of subnets in bits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.FindIPAvailabilities(context.Background(), id).Cidr(cidr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAvailabilities`: IPAvailabilitiesList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | IP Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | **string** | Size of subnets in bits | 

### Return type

[**IPAvailabilitiesList**](IPAvailabilitiesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPReservations

> IPReservationList FindIPReservations(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve all ip reservations



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
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.FindIPReservations(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPReservations`: IPReservationList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPReservationList**](IPReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestIPReservation

> IPReservation RequestIPReservation(ctx, id).IpReservationRequest(ipReservationRequest).Execute()

Requesting IP reservations



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
    ipReservationRequest := *openapiclient.NewIPReservationRequestInput("Type_example", int32(123)) // IPReservationRequestInput | IP Reservation Request to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.RequestIPReservation(context.Background(), id).IpReservationRequest(ipReservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.RequestIPReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestIPReservation`: IPReservation
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.RequestIPReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestIPReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipReservationRequest** | [**IPReservationRequestInput**](IPReservationRequestInput.md) | IP Reservation Request to create | 

### Return type

[**IPReservation**](IPReservation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPAddress

> IPAssignment UpdateIPAddress(ctx, id).Details(details).Customdata(customdata).Execute()

Update an ip address



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
    id := TODO // string | IP Address UUID
    details := "details_example" // string | Notes for this IP Assignment
    customdata := "customdata_example" // string | Provides the custom metadata stored for this IP Assignment in json format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesApi.UpdateIPAddress(context.Background(), id).Details(details).Customdata(customdata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.UpdateIPAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIPAddress`: IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.UpdateIPAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **string** | Notes for this IP Assignment | 
 **customdata** | **string** | Provides the custom metadata stored for this IP Assignment in json format | 

### Return type

[**IPAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

