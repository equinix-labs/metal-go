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

> map[string]interface{} CreateConnectionPortVirtualCircuit(ctx, connectionId, portId).Body(body).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection
    portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection port
    body := map[string]interface{}{ ... } // map[string]interface{} | Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.CreateConnectionPortVirtualCircuit(context.Background(), connectionId, portId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnectionPortVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectionPortVirtualCircuit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnectionPortVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the connection | 
**portId** | **string** | UUID of the connection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionPortVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Virtual Circuit details | 

### Return type

**map[string]interface{}**

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInterconnection

> GetInterconnection200Response CreateOrganizationInterconnection(ctx, organizationId).Body(body).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization
    body := *openapiclient.NewCreateOrganizationInterconnectionRequest("Metro_example", "Name_example", "Redundancy_example", "Type_example") // CreateOrganizationInterconnectionRequest | Connection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.CreateOrganizationInterconnection(context.Background(), organizationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateOrganizationInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInterconnection`: GetInterconnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateOrganizationInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrganizationInterconnectionRequest**](CreateOrganizationInterconnectionRequest.md) | Connection details | 

### Return type

[**GetInterconnection200Response**](GetInterconnection200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectInterconnection

> GetInterconnection200Response CreateProjectInterconnection(ctx, projectId).Body(body).Execute()

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the project
    body := *openapiclient.NewCreateOrganizationInterconnectionRequest("Metro_example", "Name_example", "Redundancy_example", "Type_example") // CreateOrganizationInterconnectionRequest | Connection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.CreateProjectInterconnection(context.Background(), projectId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateProjectInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectInterconnection`: GetInterconnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateProjectInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrganizationInterconnectionRequest**](CreateOrganizationInterconnectionRequest.md) | Connection details | 

### Return type

[**GetInterconnection200Response**](GetInterconnection200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterconnection

> GetInterconnection200Response DeleteInterconnection(ctx, connectionId).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.DeleteInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInterconnection`: GetInterconnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInterconnection200Response**](GetInterconnection200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualCircuit

> map[string]interface{} DeleteVirtualCircuit(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.DeleteVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualCircuit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionPort

> GetInterconnection200ResponsePortsInner GetConnectionPort(ctx, connectionId, id).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetConnectionPort(context.Background(), connectionId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnectionPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionPort`: GetInterconnection200ResponsePortsInner
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnectionPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the connection | 
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetInterconnection200ResponsePortsInner**](GetInterconnection200ResponsePortsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnection

> GetInterconnection200Response GetInterconnection(ctx, connectionId).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterconnection`: GetInterconnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInterconnection200Response**](GetInterconnection200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualCircuit

> map[string]interface{} GetVirtualCircuit(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualCircuit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPortVirtualCircuits

> GetInterconnection200ResponsePortsInnerVirtualCircuits ListConnectionPortVirtualCircuits(ctx, connectionId, portId).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection
    portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection port

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.ListConnectionPortVirtualCircuits(context.Background(), connectionId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ListConnectionPortVirtualCircuits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPortVirtualCircuits`: GetInterconnection200ResponsePortsInnerVirtualCircuits
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ListConnectionPortVirtualCircuits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the connection | 
**portId** | **string** | UUID of the connection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionPortVirtualCircuitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetInterconnection200ResponsePortsInnerVirtualCircuits**](GetInterconnection200ResponsePortsInnerVirtualCircuits.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPorts

> ListConnectionPorts200Response ListConnectionPorts(ctx, connectionId).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the connection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.ListConnectionPorts(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ListConnectionPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPorts`: ListConnectionPorts200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ListConnectionPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListConnectionPorts200Response**](ListConnectionPorts200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationListInterconnections

> OrganizationListInterconnections200Response OrganizationListInterconnections(ctx, organizationId).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.OrganizationListInterconnections(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.OrganizationListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationListInterconnections`: OrganizationListInterconnections200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.OrganizationListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationListInterconnections200Response**](OrganizationListInterconnections200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListInterconnections

> OrganizationListInterconnections200Response ProjectListInterconnections(ctx, projectId).Execute()

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.ProjectListInterconnections(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ProjectListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectListInterconnections`: OrganizationListInterconnections200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ProjectListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationListInterconnections200Response**](OrganizationListInterconnections200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterconnection

> GetInterconnection200Response UpdateInterconnection(ctx, connectionId).Body(body).Execute()

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
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID
    body := *openapiclient.NewUpdateInterconnectionRequest() // UpdateInterconnectionRequest | Updated connection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.UpdateInterconnection(context.Background(), connectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterconnection`: GetInterconnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateInterconnectionRequest**](UpdateInterconnectionRequest.md) | Updated connection details | 

### Return type

[**GetInterconnection200Response**](GetInterconnection200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualCircuit

> map[string]interface{} UpdateVirtualCircuit(ctx, id).Body(body).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID
    body := map[string]interface{}{ ... } // map[string]interface{} | Updated Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.UpdateVirtualCircuit(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualCircuit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Updated Virtual Circuit details | 

### Return type

**map[string]interface{}**

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

