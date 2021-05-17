# \BGPApi

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpSession**](BGPApi.md#CreateBgpSession) | **Post** /devices/{id}/bgp/sessions | Create a BGP session
[**DeleteBgpSession**](BGPApi.md#DeleteBgpSession) | **Delete** /bgp/sessions/{id} | Delete the BGP session
[**FindBgpConfigByProject**](BGPApi.md#FindBgpConfigByProject) | **Get** /projects/{id}/bgp-config | Retrieve a bgp config
[**FindBgpSessionById**](BGPApi.md#FindBgpSessionById) | **Get** /bgp/sessions/{id} | Retrieve a BGP session
[**FindBgpSessions**](BGPApi.md#FindBgpSessions) | **Get** /devices/{id}/bgp/sessions | Retrieve all BGP sessions
[**FindProjectBgpSessions**](BGPApi.md#FindProjectBgpSessions) | **Get** /projects/{id}/bgp/sessions | Retrieve all BGP sessions for project
[**GetBgpNeighborData**](BGPApi.md#GetBgpNeighborData) | **Get** /devices/{id}/bgp/neighbors | Retrieve BGP neighbor data for this device
[**RequestBgpConfig**](BGPApi.md#RequestBgpConfig) | **Post** /projects/{id}/bgp-configs | Requesting bgp config
[**UpdateBgpSession**](BGPApi.md#UpdateBgpSession) | **Put** /bgp/sessions/{id} | Update the BGP session



## CreateBgpSession

> BgpSession CreateBgpSession(ctx, id).BgpSession(bgpSession).Execute()

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
    id := TODO // string | Device UUID
    bgpSession := *openapiclient.NewBGPSessionInput() // BGPSessionInput | BGP session to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.CreateBgpSession(context.Background(), id).BgpSession(bgpSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.CreateBgpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpSession`: BgpSession
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.CreateBgpSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpSession** | [**BGPSessionInput**](BGPSessionInput.md) | BGP session to create | 

### Return type

[**BgpSession**](BgpSession.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBgpSession

> DeleteBgpSession(ctx, id).Execute()

Delete the BGP session



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
    id := TODO // string | BGP session UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.DeleteBgpSession(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.DeleteBgpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | BGP session UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBgpSessionRequest struct via the builder pattern


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


## FindBgpConfigByProject

> BgpConfig FindBgpConfigByProject(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a bgp config



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
    resp, r, err := api_client.BGPApi.FindBgpConfigByProject(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.FindBgpConfigByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBgpConfigByProject`: BgpConfig
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.FindBgpConfigByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBgpConfigByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**BgpConfig**](BgpConfig.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBgpSessionById

> BgpSession FindBgpSessionById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a BGP session



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
    id := TODO // string | BGP session UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.FindBgpSessionById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.FindBgpSessionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBgpSessionById`: BgpSession
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.FindBgpSessionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | BGP session UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBgpSessionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**BgpSession**](BgpSession.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBgpSessions

> BgpSessionList FindBgpSessions(ctx, id).Execute()

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
    id := TODO // string | Device UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.FindBgpSessions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.FindBgpSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBgpSessions`: BgpSessionList
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.FindBgpSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBgpSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpSessionList**](BgpSessionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectBgpSessions

> BgpSessionList FindProjectBgpSessions(ctx, id).Execute()

Retrieve all BGP sessions for project



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.FindProjectBgpSessions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.FindProjectBgpSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectBgpSessions`: BgpSessionList
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.FindProjectBgpSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectBgpSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpSessionList**](BgpSessionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpNeighborData

> BgpSessionNeighbors GetBgpNeighborData(ctx, id).Execute()

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
    id := TODO // string | Device UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.GetBgpNeighborData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.GetBgpNeighborData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBgpNeighborData`: BgpSessionNeighbors
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.GetBgpNeighborData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpNeighborDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpSessionNeighbors**](BgpSessionNeighbors.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestBgpConfig

> RequestBgpConfig(ctx, id).BgpConfigRequest(bgpConfigRequest).Execute()

Requesting bgp config



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
    bgpConfigRequest := *openapiclient.NewBgpConfigRequestInput("DeploymentType_example", int32(123)) // BgpConfigRequestInput | BGP config Request to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.RequestBgpConfig(context.Background(), id).BgpConfigRequest(bgpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.RequestBgpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestBgpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpConfigRequest** | [**BgpConfigRequestInput**](BgpConfigRequestInput.md) | BGP config Request to create | 

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


## UpdateBgpSession

> UpdateBgpSession(ctx, id).DefaultRoute(defaultRoute).Execute()

Update the BGP session



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
    id := TODO // string | BGP session UUID
    defaultRoute := true // bool | Default route

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BGPApi.UpdateBgpSession(context.Background(), id).DefaultRoute(defaultRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.UpdateBgpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | BGP session UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultRoute** | **bool** | Default route | 

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

