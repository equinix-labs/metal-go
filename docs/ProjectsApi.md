# \ProjectsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /projects | Create a project
[**CreateProjectInvitation**](ProjectsApi.md#CreateProjectInvitation) | **Post** /projects/{project_id}/invitations | Create an invitation for a project
[**CreateTransferRequest**](ProjectsApi.md#CreateTransferRequest) | **Post** /projects/{id}/transfers | Create a transfer request
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /projects/{id} | Delete the project
[**FindIPReservationCustomdata**](ProjectsApi.md#FindIPReservationCustomdata) | **Get** /projects/{project_id}/ips/{id}/customdata | Retrieve the custom metadata of an IP Reservation
[**FindProjectById**](ProjectsApi.md#FindProjectById) | **Get** /projects/{id} | Retrieve a project
[**FindProjectCustomdata**](ProjectsApi.md#FindProjectCustomdata) | **Get** /projects/{id}/customdata | Retrieve the custom metadata of a project
[**FindProjectInvitations**](ProjectsApi.md#FindProjectInvitations) | **Get** /projects/{project_id}/invitations | Retrieve project invitations
[**FindProjectMemberships**](ProjectsApi.md#FindProjectMemberships) | **Get** /projects/{project_id}/memberships | Retrieve project memberships
[**FindProjects**](ProjectsApi.md#FindProjects) | **Get** /projects | Retrieve all projects
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /projects/{id} | Update the project



## CreateProject

> MoveHardwareReservation201ResponseProject CreateProject(ctx).Body(body).Execute()

Create a project



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
    body := *openapiclient.NewCreateProjectRequest("Name_example") // CreateProjectRequest | Project to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProject(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: MoveHardwareReservation201ResponseProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateProjectRequest**](CreateProjectRequest.md) | Project to create | 

### Return type

[**MoveHardwareReservation201ResponseProject**](MoveHardwareReservation201ResponseProject.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectInvitation

> FindInvitationById200Response CreateProjectInvitation(ctx, projectId).Body(body).Execute()

Create an invitation for a project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    body := *openapiclient.NewCreateOrganizationInvitationRequest("Invitee_example") // CreateOrganizationInvitationRequest | Invitation to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectInvitation(context.Background(), projectId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectInvitation`: FindInvitationById200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrganizationInvitationRequest**](CreateOrganizationInvitationRequest.md) | Invitation to create | 

### Return type

[**FindInvitationById200Response**](FindInvitationById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransferRequest

> FindOrganizationTransfers200ResponseTransfersInner CreateTransferRequest(ctx, id).Body(body).Execute()

Create a transfer request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the project to be transferred
    body := *openapiclient.NewCreateTransferRequestRequest() // CreateTransferRequestRequest | Transfer Request to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateTransferRequest(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateTransferRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransferRequest`: FindOrganizationTransfers200ResponseTransfersInner
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateTransferRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the project to be transferred | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransferRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateTransferRequestRequest**](CreateTransferRequestRequest.md) | Transfer Request to create | 

### Return type

[**FindOrganizationTransfers200ResponseTransfersInner**](FindOrganizationTransfers200ResponseTransfersInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, id).Execute()

Delete the project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.DeleteProject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## FindIPReservationCustomdata

> FindIPReservationCustomdata(ctx, projectId, id).Execute()

Retrieve the custom metadata of an IP Reservation



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Ip Reservation UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.FindIPReservationCustomdata(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindIPReservationCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 
**id** | **string** | Ip Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPReservationCustomdataRequest struct via the builder pattern


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


## FindProjectById

> MoveHardwareReservation201ResponseProject FindProjectById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.FindProjectById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindProjectById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectById`: MoveHardwareReservation201ResponseProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.FindProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**MoveHardwareReservation201ResponseProject**](MoveHardwareReservation201ResponseProject.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectCustomdata

> FindProjectCustomdata(ctx, id).Execute()

Retrieve the custom metadata of a project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.FindProjectCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindProjectCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectCustomdataRequest struct via the builder pattern


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


## FindProjectInvitations

> FindInvitations200Response FindProjectInvitations(ctx, projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve project invitations



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.FindProjectInvitations(context.Background(), projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindProjectInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectInvitations`: FindInvitations200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.FindProjectInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindInvitations200Response**](FindInvitations200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectMemberships

> FindProjectMemberships200Response FindProjectMemberships(ctx, projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve project memberships



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.FindProjectMemberships(context.Background(), projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindProjectMemberships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectMemberships`: FindProjectMemberships200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.FindProjectMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindProjectMemberships200Response**](FindProjectMemberships200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjects

> FindOrganizationProjects200Response FindProjects(ctx).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all projects



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
    resp, r, err := apiClient.ProjectsApi.FindProjects(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.FindProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjects`: FindOrganizationProjects200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.FindProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindOrganizationProjects200Response**](FindOrganizationProjects200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> MoveHardwareReservation201ResponseProject UpdateProject(ctx, id).Body(body).Execute()

Update the project



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
    body := *openapiclient.NewUpdateProjectRequest() // UpdateProjectRequest | Project to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.UpdateProject(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: MoveHardwareReservation201ResponseProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateProjectRequest**](UpdateProjectRequest.md) | Project to update | 

### Return type

[**MoveHardwareReservation201ResponseProject**](MoveHardwareReservation201ResponseProject.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

