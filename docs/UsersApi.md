# \UsersApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users | Create a user
[**FindCurrentUser**](UsersApi.md#FindCurrentUser) | **Get** /user | Retrieve the current user
[**FindInvitations**](UsersApi.md#FindInvitations) | **Get** /invitations | Retrieve current user invitations
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /users/{id} | Retrieve a user
[**FindUserCustomdata**](UsersApi.md#FindUserCustomdata) | **Get** /users/{id}/customdata | Retrieve the custom metadata of a user
[**FindUsers**](UsersApi.md#FindUsers) | **Get** /users | Retrieve all users
[**UpdateCurrentUser**](UsersApi.md#UpdateCurrentUser) | **Put** /user | Update the current user



## CreateUser

> FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf CreateUser(ctx).Body(body).Execute()

Create a user



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
    body := *openapiclient.NewCreateUserRequest([]openapiclient.CreateUserRequestEmailsInner{*openapiclient.NewCreateUserRequestEmailsInner("Address_example")}, "FirstName_example", "LastName_example") // CreateUserRequest | User to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUserRequest**](CreateUserRequest.md) | User to create | 

### Return type

[**FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf**](FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCurrentUser

> FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf FindCurrentUser(ctx).Include(include).Exclude(exclude).Execute()

Retrieve the current user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindCurrentUser(context.Background()).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCurrentUser`: FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf**](FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInvitations

> FindInvitations200Response FindInvitations(ctx).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve current user invitations



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
    resp, r, err := apiClient.UsersApi.FindInvitations(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInvitations`: FindInvitations200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindInvitationsRequest struct via the builder pattern


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


## FindUserById

> FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf FindUserById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a user



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUserById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUserById`: FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf**](FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserCustomdata

> FindUserCustomdata(ctx, id).Execute()

Retrieve the custom metadata of a user



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUserCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUserCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserCustomdataRequest struct via the builder pattern


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


## FindUsers

> FindUsers200Response FindUsers(ctx).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all users



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
    resp, r, err := apiClient.UsersApi.FindUsers(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUsers`: FindUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindUsers200Response**](FindUsers200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUser

> FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf UpdateCurrentUser(ctx).Body(body).Execute()

Update the current user



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
    body := *openapiclient.NewUpdateCurrentUserRequest() // UpdateCurrentUserRequest | User to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UpdateCurrentUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCurrentUser`: FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateCurrentUserRequest**](UpdateCurrentUserRequest.md) | User to update | 

### Return type

[**FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf**](FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

