# \ConnectionsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionPortVirtualCircuit**](ConnectionsApi.md#CreateConnectionPortVirtualCircuit) | **Post** /connections/{connection_id}/ports/{port_id}/virtual-circuits | Create a new Virtual Circuit
[**CreateOrganizationInterconnection**](ConnectionsApi.md#CreateOrganizationInterconnection) | **Post** /organizations/{organization_id}/connections | Request a new connection for the organization
[**CreateProjectInterconnection**](ConnectionsApi.md#CreateProjectInterconnection) | **Post** /projects/{project_id}/connections | Request a new connection for the project&#39;s organization
[**DeleteInterconnection**](ConnectionsApi.md#DeleteInterconnection) | **Delete** /connections/{connection_id} | Delete connection
[**DeleteVirtualCircuit**](ConnectionsApi.md#DeleteVirtualCircuit) | **Delete** /virtual-circuits/{id} | Delete a virtual circuit
[**GetConnectionPort**](ConnectionsApi.md#GetConnectionPort) | **Get** /connections/{connection_id}/ports/{id} | Get a connection port
[**GetInterconnection**](ConnectionsApi.md#GetInterconnection) | **Get** /connections/{connection_id} | Get connection
[**GetVirtualCircuit**](ConnectionsApi.md#GetVirtualCircuit) | **Get** /virtual-circuits/{id} | Get a virtual circuit
[**ListConnectionPortVirtualCircuits**](ConnectionsApi.md#ListConnectionPortVirtualCircuits) | **Get** /connections/{connection_id}/ports/{port_id}/virtual-circuits | List a connection port&#39;s virtual circuits
[**ListConnectionPorts**](ConnectionsApi.md#ListConnectionPorts) | **Get** /connections/{connection_id}/ports | List a connection&#39;s ports
[**OrganizationListInterconnections**](ConnectionsApi.md#OrganizationListInterconnections) | **Get** /organizations/{organization_id}/connections | List organization connections
[**ProjectListInterconnections**](ConnectionsApi.md#ProjectListInterconnections) | **Get** /projects/{project_id}/connections | List project connections
[**UpdateInterconnection**](ConnectionsApi.md#UpdateInterconnection) | **Put** /connections/{connection_id} | Update connection
[**UpdateVirtualCircuit**](ConnectionsApi.md#UpdateVirtualCircuit) | **Put** /virtual-circuits/{id} | Update a virtual circuit



## CreateConnectionPortVirtualCircuit

> VirtualCircuitList CreateConnectionPortVirtualCircuit(ctx, connectionId, portId).VirtualCircuit(virtualCircuit).Execute()

Create a new Virtual Circuit



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
    connectionId := TODO // string | UUID of the connection
    portId := TODO // string | UUID of the connection port
    virtualCircuit := *openapiclient.NewVirtualCircuitCreateInput() // VirtualCircuitCreateInput | Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.CreateConnectionPortVirtualCircuit(context.Background(), connectionId, portId).VirtualCircuit(virtualCircuit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnectionPortVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectionPortVirtualCircuit`: VirtualCircuitList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnectionPortVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | UUID of the connection | 
**portId** | [**string**](.md) | UUID of the connection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionPortVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualCircuit** | [**VirtualCircuitCreateInput**](VirtualCircuitCreateInput.md) | Virtual Circuit details | 

### Return type

[**VirtualCircuitList**](VirtualCircuitList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInterconnection

> Interconnection CreateOrganizationInterconnection(ctx, organizationId).Connection(connection).Execute()

Request a new connection for the organization



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
    organizationId := TODO // string | UUID of the organization
    connection := *openapiclient.NewInterconnectionCreateInput("Name_example", "Facility_example", "Type_example", "Redundancy_example") // InterconnectionCreateInput | Connection details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.CreateOrganizationInterconnection(context.Background(), organizationId).Connection(connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateOrganizationInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateOrganizationInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**InterconnectionCreateInput**](InterconnectionCreateInput.md) | Connection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectInterconnection

> Interconnection CreateProjectInterconnection(ctx, projectId).Connection(connection).Execute()

Request a new connection for the project's organization



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
    projectId := TODO // string | UUID of the project
    connection := *openapiclient.NewInterconnectionCreateInput("Name_example", "Facility_example", "Type_example", "Redundancy_example") // InterconnectionCreateInput | Connection details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.CreateProjectInterconnection(context.Background(), projectId).Connection(connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateProjectInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateProjectInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**InterconnectionCreateInput**](InterconnectionCreateInput.md) | Connection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterconnection

> Interconnection DeleteInterconnection(ctx, connectionId).Execute()

Delete connection



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
    connectionId := TODO // string | Connection UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.DeleteInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualCircuit

> VirtualCircuit DeleteVirtualCircuit(ctx, id).Execute()

Delete a virtual circuit



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
    id := TODO // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.DeleteVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualCircuit`: VirtualCircuit
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualCircuit**](VirtualCircuit.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionPort

> InterconnectionPort GetConnectionPort(ctx, connectionId, id).Execute()

Get a connection port



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
    connectionId := TODO // string | UUID of the connection
    id := TODO // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.GetConnectionPort(context.Background(), connectionId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnectionPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionPort`: InterconnectionPort
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnectionPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | UUID of the connection | 
**id** | [**string**](.md) | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InterconnectionPort**](InterconnectionPort.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnection

> Interconnection GetInterconnection(ctx, connectionId).Execute()

Get connection



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
    connectionId := TODO // string | Connection UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.GetInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualCircuit

> VirtualCircuit GetVirtualCircuit(ctx, id).Execute()

Get a virtual circuit



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
    id := TODO // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.GetVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualCircuit`: VirtualCircuit
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualCircuit**](VirtualCircuit.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPortVirtualCircuits

> VirtualCircuitList ListConnectionPortVirtualCircuits(ctx, connectionId, portId).Execute()

List a connection port's virtual circuits



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
    connectionId := TODO // string | UUID of the connection
    portId := TODO // string | UUID of the connection port

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.ListConnectionPortVirtualCircuits(context.Background(), connectionId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ListConnectionPortVirtualCircuits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPortVirtualCircuits`: VirtualCircuitList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ListConnectionPortVirtualCircuits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | UUID of the connection | 
**portId** | [**string**](.md) | UUID of the connection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionPortVirtualCircuitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VirtualCircuitList**](VirtualCircuitList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPorts

> InterconnectionPortList ListConnectionPorts(ctx, connectionId).Execute()

List a connection's ports



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
    connectionId := TODO // string | UUID of the connection

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.ListConnectionPorts(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ListConnectionPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPorts`: InterconnectionPortList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ListConnectionPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | UUID of the connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionPortList**](InterconnectionPortList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationListInterconnections

> InterconnectionList OrganizationListInterconnections(ctx, organizationId).Execute()

List organization connections



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
    organizationId := TODO // string | UUID of the organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.OrganizationListInterconnections(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.OrganizationListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationListInterconnections`: InterconnectionList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.OrganizationListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionList**](InterconnectionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListInterconnections

> InterconnectionList ProjectListInterconnections(ctx, projectId).Execute()

List project connections



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
    projectId := TODO // string | UUID of the project

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.ProjectListInterconnections(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ProjectListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectListInterconnections`: InterconnectionList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ProjectListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionList**](InterconnectionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterconnection

> Interconnection UpdateInterconnection(ctx, connectionId).Connection(connection).Execute()

Update connection



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
    connectionId := TODO // string | Connection UUID
    connection := *openapiclient.NewInterconnectionUpdateInput() // InterconnectionUpdateInput | Updated connection details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.UpdateInterconnection(context.Background(), connectionId).Connection(connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**](.md) | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**InterconnectionUpdateInput**](InterconnectionUpdateInput.md) | Updated connection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualCircuit

> VirtualCircuit UpdateVirtualCircuit(ctx, id).VirtualCircuit(virtualCircuit).Execute()

Update a virtual circuit



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
    id := TODO // string | Virtual Circuit UUID
    virtualCircuit := *openapiclient.NewVirtualCircuitUpdateInput() // VirtualCircuitUpdateInput | Updated Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsApi.UpdateVirtualCircuit(context.Background(), id).VirtualCircuit(virtualCircuit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualCircuit`: VirtualCircuit
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualCircuit** | [**VirtualCircuitUpdateInput**](VirtualCircuitUpdateInput.md) | Updated Virtual Circuit details | 

### Return type

[**VirtualCircuit**](VirtualCircuit.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

