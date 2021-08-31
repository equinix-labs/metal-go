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

> Port AssignNativeVlan(ctx, id).Vnid(vnid).Execute()

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
    id := TODO // string | Port UUID
    vnid := "vnid_example" // string | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: '1001').

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.AssignNativeVlan(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.AssignNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNativeVlan`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.AssignNativeVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignNativeVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vnid** | **string** | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: &#39;1001&#39;). | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignPort

> Port AssignPort(ctx, id).Vnid(vnid).Execute()

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
    id := TODO // string | Port UUID
    vnid := *openapiclient.NewPortAssignInput() // PortAssignInput | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: '1001').

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.AssignPort(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.AssignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.AssignPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vnid** | [**PortAssignInput**](PortAssignInput.md) | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: &#39;1001&#39;). | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BondPort

> Port BondPort(ctx, id).BulkEnable(bulkEnable).Execute()

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
    id := TODO // string | Port UUID
    bulkEnable := true // bool | enable both ports (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.BondPort(context.Background(), id).BulkEnable(bulkEnable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.BondPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BondPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.BondPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBondPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkEnable** | **bool** | enable both ports | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertLayer2

> Port ConvertLayer2(ctx, id).Vnid(vnid).Execute()

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
    id := TODO // string | Port UUID
    vnid := *openapiclient.NewPortAssignInput() // PortAssignInput | Virtual Network ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.ConvertLayer2(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.ConvertLayer2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvertLayer2`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.ConvertLayer2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertLayer2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vnid** | [**PortAssignInput**](PortAssignInput.md) | Virtual Network ID | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertLayer3

> Port ConvertLayer3(ctx, id).RequestIps(requestIps).Execute()

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
    id := TODO // string | Port UUID
    requestIps := *openapiclient.NewPortConvertLayer3Input() // PortConvertLayer3Input | IPs to request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.ConvertLayer3(context.Background(), id).RequestIps(requestIps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.ConvertLayer3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvertLayer3`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.ConvertLayer3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertLayer3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestIps** | [**PortConvertLayer3Input**](PortConvertLayer3Input.md) | IPs to request | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortVlanAssignmentBatch

> PortVlanAssignmentBatch CreatePortVlanAssignmentBatch(ctx, id).VlanAssignments(vlanAssignments).Execute()

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
    id := TODO // string | Port UUID
    vlanAssignments := *openapiclient.NewPortVlanAssignmentBatchCreateInput() // PortVlanAssignmentBatchCreateInput | VLAN Assignment batch details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.CreatePortVlanAssignmentBatch(context.Background(), id).VlanAssignments(vlanAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.CreatePortVlanAssignmentBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePortVlanAssignmentBatch`: PortVlanAssignmentBatch
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.CreatePortVlanAssignmentBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortVlanAssignmentBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vlanAssignments** | [**PortVlanAssignmentBatchCreateInput**](PortVlanAssignmentBatchCreateInput.md) | VLAN Assignment batch details | 

### Return type

[**PortVlanAssignmentBatch**](PortVlanAssignmentBatch.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNativeVlan

> Port DeleteNativeVlan(ctx, id).Execute()

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
    id := TODO // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.DeleteNativeVlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.DeleteNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNativeVlan`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.DeleteNativeVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNativeVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisbondPort

> Port DisbondPort(ctx, id).BulkDisable(bulkDisable).Execute()

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
    id := TODO // string | Port UUID
    bulkDisable := true // bool | disable both ports (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.DisbondPort(context.Background(), id).BulkDisable(bulkDisable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.DisbondPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisbondPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.DisbondPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisbondPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkDisable** | **bool** | disable both ports | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortById

> Port FindPortById(ctx, id).Include(include).Exclude(exclude).Execute()

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
    id := TODO // string | Port UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.FindPortById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortById`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentBatchByPortIdAndBatchId

> PortVlanAssignmentBatch FindPortVlanAssignmentBatchByPortIdAndBatchId(ctx, id, batchId).Execute()

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
    id := TODO // string | Port UUID
    batchId := TODO // string | Batch ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId(context.Background(), id, batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentBatchByPortIdAndBatchId`: PortVlanAssignmentBatch
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 
**batchId** | [**string**](.md) | Batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentBatchByPortIdAndBatchIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PortVlanAssignmentBatch**](PortVlanAssignmentBatch.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentBatches

> PortVlanAssignmentBatchList FindPortVlanAssignmentBatches(ctx, id).Execute()

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
    id := TODO // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.FindPortVlanAssignmentBatches(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentBatches`: PortVlanAssignmentBatchList
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortVlanAssignmentBatchList**](PortVlanAssignmentBatchList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignmentByPortIdAndAssignmentId

> PortVlanAssignment FindPortVlanAssignmentByPortIdAndAssignmentId(ctx, id, assignmentId).Include(include).Exclude(exclude).Execute()

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
    id := TODO // string | Port UUID
    assignmentId := TODO // string | Assignment ID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional) (default to ["port","virtual_network"])
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId(context.Background(), id, assignmentId).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignmentByPortIdAndAssignmentId`: PortVlanAssignment
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 
**assignmentId** | [**string**](.md) | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentByPortIdAndAssignmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | [default to [&quot;port&quot;,&quot;virtual_network&quot;]]
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**PortVlanAssignment**](PortVlanAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortVlanAssignments

> PortVlanAssignmentList FindPortVlanAssignments(ctx, id).Include(include).Exclude(exclude).Execute()

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
    id := TODO // string | Port UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional) (default to ["port","virtual_network"])
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.FindPortVlanAssignments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortVlanAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPortVlanAssignments`: PortVlanAssignmentList
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.FindPortVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortVlanAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | [default to [&quot;port&quot;,&quot;virtual_network&quot;]]
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**PortVlanAssignmentList**](PortVlanAssignmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignPort

> Port UnassignPort(ctx, id).Vnid(vnid).Execute()

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
    id := TODO // string | Port UUID
    vnid := *openapiclient.NewPortAssignInput() // PortAssignInput | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: '1001').

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortsApi.UnassignPort(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.UnassignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `PortsApi.UnassignPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vnid** | [**PortAssignInput**](PortAssignInput.md) | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself (ex: &#39;1001&#39;). | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

