# \VLANsApi

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNativeVlan**](VLANsApi.md#AssignNativeVlan) | **Post** /ports/{id}/native-vlan | Assign a native VLAN
[**AssignPort**](VLANsApi.md#AssignPort) | **Post** /ports/{id}/assign | Assign a port to virtual network
[**CreateVirtualNetwork**](VLANsApi.md#CreateVirtualNetwork) | **Post** /projects/{id}/virtual-networks | Create a virtual network
[**DeleteNativeVlan**](VLANsApi.md#DeleteNativeVlan) | **Delete** /ports/{id}/native-vlan | Remove native VLAN
[**DeleteVirtualNetwork**](VLANsApi.md#DeleteVirtualNetwork) | **Delete** /virtual-networks/{id} | Delete a virtual network
[**FindVirtualNetworks**](VLANsApi.md#FindVirtualNetworks) | **Get** /projects/{id}/virtual-networks | Retrieve all virtual networks
[**GetVirtualNetwork**](VLANsApi.md#GetVirtualNetwork) | **Get** /virtual-networks/{id} | Get a virtual network
[**UnassignPort**](VLANsApi.md#UnassignPort) | **Post** /ports/{id}/unassign | Unassign a port



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
    resp, r, err := api_client.VLANsApi.AssignNativeVlan(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.AssignNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNativeVlan`: Port
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.AssignNativeVlan`: %v\n", resp)
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
    resp, r, err := api_client.VLANsApi.AssignPort(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.AssignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.AssignPort`: %v\n", resp)
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


## CreateVirtualNetwork

> VirtualNetwork CreateVirtualNetwork(ctx, id).VirtualNetwork(virtualNetwork).Execute()

Create a virtual network



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
    virtualNetwork := *openapiclient.NewVirtualNetworkCreateInput("ProjectId_example") // VirtualNetworkCreateInput | Virtual Network to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VLANsApi.CreateVirtualNetwork(context.Background(), id).VirtualNetwork(virtualNetwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualNetwork`: VirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.CreateVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualNetwork** | [**VirtualNetworkCreateInput**](VirtualNetworkCreateInput.md) | Virtual Network to create | 

### Return type

[**VirtualNetwork**](VirtualNetwork.md)

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
    resp, r, err := api_client.VLANsApi.DeleteNativeVlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.DeleteNativeVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNativeVlan`: Port
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.DeleteNativeVlan`: %v\n", resp)
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


## DeleteVirtualNetwork

> VirtualNetwork DeleteVirtualNetwork(ctx, id).Execute()

Delete a virtual network



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
    id := TODO // string | Virtual Network UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VLANsApi.DeleteVirtualNetwork(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.DeleteVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualNetwork`: VirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.DeleteVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualNetwork**](VirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVirtualNetworks

> VirtualNetworkList FindVirtualNetworks(ctx, id).Include(include).Exclude(exclude).Facility(facility).Metro(metro).Execute()

Retrieve all virtual networks



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
    facility := "facility_example" // string | Filter by Facility ID (uuid) or Facility Code (optional)
    metro := "metro_example" // string | Filter by Metro ID (uuid) or Metro Code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VLANsApi.FindVirtualNetworks(context.Background(), id).Include(include).Exclude(exclude).Facility(facility).Metro(metro).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.FindVirtualNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVirtualNetworks`: VirtualNetworkList
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.FindVirtualNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVirtualNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **facility** | **string** | Filter by Facility ID (uuid) or Facility Code | 
 **metro** | **string** | Filter by Metro ID (uuid) or Metro Code | 

### Return type

[**VirtualNetworkList**](VirtualNetworkList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualNetwork

> VirtualNetwork GetVirtualNetwork(ctx, id).Execute()

Get a virtual network



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
    id := TODO // string | Virtual Network UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VLANsApi.GetVirtualNetwork(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.GetVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualNetwork`: VirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.GetVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualNetwork**](VirtualNetwork.md)

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
    resp, r, err := api_client.VLANsApi.UnassignPort(context.Background(), id).Vnid(vnid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.UnassignPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignPort`: Port
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.UnassignPort`: %v\n", resp)
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

