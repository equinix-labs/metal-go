# \PortsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNativeVlan**](PortsApi.md#AssignNativeVlan) | **Post** /ports/{id}/native-vlan | Assign a native VLAN
[**AssignPort**](PortsApi.md#AssignPort) | **Post** /ports/{id}/assign | Assign a port to virtual network
[**BondPort**](PortsApi.md#BondPort) | **Post** /ports/{id}/bond | Enabling bonding
[**ConvertLayer2**](PortsApi.md#ConvertLayer2) | **Post** /ports/{id}/convert/layer-2 | Convert to Layer 2
[**ConvertLayer3**](PortsApi.md#ConvertLayer3) | **Post** /ports/{id}/convert/layer-3 | Convert to Layer 3
[**CreatePortVlanAssignmentBatch**](PortsApi.md#CreatePortVlanAssignmentBatch) | **Post** /ports/{id}/vlan-assignments/batches | Create a new Port-VLAN Assignment management batch
[**DeleteNativeVlan**](PortsApi.md#DeleteNativeVlan) | **Delete** /ports/{id}/native-vlan | Remove native VLAN
[**DisbondPort**](PortsApi.md#DisbondPort) | **Post** /ports/{id}/disbond | Disabling bonding
[**FindPortById**](PortsApi.md#FindPortById) | **Get** /ports/{id} | Retrieve a port
[**FindPortVlanAssignmentBatchByPortIdAndBatchId**](PortsApi.md#FindPortVlanAssignmentBatchByPortIdAndBatchId) | **Get** /ports/{id}/vlan-assignments/batches/{batch_id} | Retrieve a VLAN Assignment Batch&#39;s details
[**FindPortVlanAssignmentBatches**](PortsApi.md#FindPortVlanAssignmentBatches) | **Get** /ports/{id}/vlan-assignments/batches | List the VLAN Assignment Batches for a port
[**FindPortVlanAssignmentByPortIdAndAssignmentId**](PortsApi.md#FindPortVlanAssignmentByPortIdAndAssignmentId) | **Get** /ports/{id}/vlan-assignments/{assignment_id} | Show a particular Port VLAN assignment&#39;s details
[**FindPortVlanAssignments**](PortsApi.md#FindPortVlanAssignments) | **Get** /ports/{id}/vlan-assignments | List Current VLAN assignments for a port
[**UnassignPort**](PortsApi.md#UnassignPort) | **Post** /ports/{id}/unassign | Unassign a port



## AssignNativeVlan

> FindDeviceById200ResponseNetworkPortsInner AssignNativeVlan(ctx, id).Vnid(vnid).Execute()

Assign a native VLAN



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    vnid := "vnid_example" // string | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: '1001').

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.AssignNativeVlan(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.AssignNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNativeVlan`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.AssignNativeVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignNativeVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vnid** | **string** | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: &#39;1001&#39;). | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignPort

> FindDeviceById200ResponseNetworkPortsInner AssignPort(ctx, id).AssignPortRequest(assignPortRequest).Execute()

Assign a port to virtual network



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    assignPortRequest := *openapiclient.NewAssignPortRequest() // AssignPortRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.AssignPort(context.Background(), id).AssignPortRequest(assignPortRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.AssignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignPort`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.AssignPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignPortRequest** | [**AssignPortRequest**](AssignPortRequest.md) |  | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BondPort

> FindDeviceById200ResponseNetworkPortsInner BondPort(ctx, id).BulkEnable(bulkEnable).Execute()

Enabling bonding



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    bulkEnable := true // bool | enable both ports (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.BondPort(context.Background(), id).BulkEnable(bulkEnable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.BondPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BondPort`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.BondPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBondPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkEnable** | **bool** | enable both ports | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertLayer2

> FindDeviceById200ResponseNetworkPortsInner ConvertLayer2(ctx, id).AssignPortRequest(assignPortRequest).Execute()

Convert to Layer 2



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    assignPortRequest := *openapiclient.NewAssignPortRequest() // AssignPortRequest | Virtual Network ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.ConvertLayer2(context.Background(), id).AssignPortRequest(assignPortRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.ConvertLayer2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvertLayer2`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.ConvertLayer2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertLayer2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignPortRequest** | [**AssignPortRequest**](AssignPortRequest.md) | Virtual Network ID | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertLayer3

> FindDeviceById200ResponseNetworkPortsInner ConvertLayer3(ctx, id).ConvertLayer3Request(convertLayer3Request).Execute()

Convert to Layer 3



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    convertLayer3Request := *openapiclient.NewConvertLayer3Request() // ConvertLayer3Request | IPs to request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.ConvertLayer3(context.Background(), id).ConvertLayer3Request(convertLayer3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.ConvertLayer3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvertLayer3`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.ConvertLayer3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertLayer3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convertLayer3Request** | [**ConvertLayer3Request**](ConvertLayer3Request.md) | IPs to request | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortVlanAssignmentBatch

> FindPortVlanAssignmentBatches200ResponseBatchesInner CreatePortVlanAssignmentBatch(ctx, id).CreatePortVlanAssignmentBatchRequest(createPortVlanAssignmentBatchRequest).Execute()

Create a new Port-VLAN Assignment management batch



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    createPortVlanAssignmentBatchRequest := *openapiclient.NewCreatePortVlanAssignmentBatchRequest() // CreatePortVlanAssignmentBatchRequest | VLAN Assignment batch details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.CreatePortVlanAssignmentBatch(context.Background(), id).CreatePortVlanAssignmentBatchRequest(createPortVlanAssignmentBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.CreatePortVlanAssignmentBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePortVlanAssignmentBatch`: FindPortVlanAssignmentBatches200ResponseBatchesInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.CreatePortVlanAssignmentBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortVlanAssignmentBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPortVlanAssignmentBatchRequest** | [**CreatePortVlanAssignmentBatchRequest**](CreatePortVlanAssignmentBatchRequest.md) | VLAN Assignment batch details | 

### Return type

[**FindPortVlanAssignmentBatches200ResponseBatchesInner**](FindPortVlanAssignmentBatches200ResponseBatchesInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNativeVlan

> FindDeviceById200ResponseNetworkPortsInner DeleteNativeVlan(ctx, id).Execute()

Remove native VLAN



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.DeleteNativeVlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.DeleteNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNativeVlan`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.DeleteNativeVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNativeVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisbondPort

> FindDeviceById200ResponseNetworkPortsInner DisbondPort(ctx, id).BulkDisable(bulkDisable).Execute()

Disabling bonding



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    bulkDisable := true // bool | disable both ports (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.DisbondPort(context.Background(), id).BulkDisable(bulkDisable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.DisbondPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisbondPort`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.DisbondPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisbondPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkDisable** | **bool** | disable both ports | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortById

> FindDeviceById200ResponseNetworkPortsInner FindPortById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a port



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.FindPortById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortById`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentBatchByPortIdAndBatchId

> FindPortVlanAssignmentBatches200ResponseBatchesInner FindPortVlanAssignmentBatchByPortIdAndBatchId(ctx, id, batchId).Execute()

Retrieve a VLAN Assignment Batch's details



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    batchId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Batch ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId(context.Background(), id, batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentBatchByPortIdAndBatchId`: FindPortVlanAssignmentBatches200ResponseBatchesInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 
**batchId** | **string** | Batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentBatchByPortIdAndBatchIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FindPortVlanAssignmentBatches200ResponseBatchesInner**](FindPortVlanAssignmentBatches200ResponseBatchesInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentBatches

> FindPortVlanAssignmentBatches200Response FindPortVlanAssignmentBatches(ctx, id).Execute()

List the VLAN Assignment Batches for a port



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.FindPortVlanAssignmentBatches(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentBatches`: FindPortVlanAssignmentBatches200Response
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindPortVlanAssignmentBatches200Response**](FindPortVlanAssignmentBatches200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentByPortIdAndAssignmentId

> FindPortVlanAssignments200ResponseVlanAssignmentsInner FindPortVlanAssignmentByPortIdAndAssignmentId(ctx, id, assignmentId).Include(include).Exclude(exclude).Execute()

Show a particular Port VLAN assignment's details



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    assignmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Assignment ID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional) (default to ["port","virtual_network"])
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId(context.Background(), id, assignmentId).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentByPortIdAndAssignmentId`: FindPortVlanAssignments200ResponseVlanAssignmentsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 
**assignmentId** | **string** | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentByPortIdAndAssignmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | [default to [&quot;port&quot;,&quot;virtual_network&quot;]]
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindPortVlanAssignments200ResponseVlanAssignmentsInner**](FindPortVlanAssignments200ResponseVlanAssignmentsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignments

> FindPortVlanAssignments200Response FindPortVlanAssignments(ctx, id).Include(include).Exclude(exclude).Execute()

List Current VLAN assignments for a port



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional) (default to ["port","virtual_network"])
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.FindPortVlanAssignments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignments`: FindPortVlanAssignments200Response
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | [default to [&quot;port&quot;,&quot;virtual_network&quot;]]
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindPortVlanAssignments200Response**](FindPortVlanAssignments200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignPort

> FindDeviceById200ResponseNetworkPortsInner UnassignPort(ctx, id).AssignPortRequest(assignPortRequest).Execute()

Unassign a port



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
    assignPortRequest := *openapiclient.NewAssignPortRequest() // AssignPortRequest | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: '1001').

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsApi.UnassignPort(context.Background(), id).AssignPortRequest(assignPortRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.UnassignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignPort`: FindDeviceById200ResponseNetworkPortsInner
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.UnassignPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignPortRequest** | [**AssignPortRequest**](AssignPortRequest.md) | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: &#39;1001&#39;). | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

