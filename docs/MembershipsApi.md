# \MembershipsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMembership**](MembershipsApi.md#DeleteMembership) | **Delete** /memberships/{id} | Delete the membership
[**FindMembershipById**](MembershipsApi.md#FindMembershipById) | **Get** /memberships/{id} | Retrieve a membership
[**UpdateMembership**](MembershipsApi.md#UpdateMembership) | **Put** /memberships/{id} | Update the membership



## DeleteMembership

> DeleteMembership(ctx, id).Execute()

Delete the membership



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Membership UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.DeleteMembership(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.DeleteMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Membership UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMembershipRequest struct via the builder pattern


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


## FindMembershipById

> FindInvitations200ResponseInvitationsInner FindMembershipById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a membership



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Membership UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.FindMembershipById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.FindMembershipById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMembershipById`: FindInvitations200ResponseInvitationsInner
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.FindMembershipById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Membership UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMembershipByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindInvitations200ResponseInvitationsInner**](FindInvitations200ResponseInvitationsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMembership

> FindInvitations200ResponseInvitationsInner UpdateMembership(ctx, id).UpdateMembershipRequest(updateMembershipRequest).Execute()

Update the membership



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Membership UUID
    updateMembershipRequest := *openapiclient.NewUpdateMembershipRequest() // UpdateMembershipRequest | Membership to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.UpdateMembership(context.Background(), id).UpdateMembershipRequest(updateMembershipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.UpdateMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMembership`: FindInvitations200ResponseInvitationsInner
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.UpdateMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Membership UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMembershipRequest** | [**UpdateMembershipRequest**](UpdateMembershipRequest.md) | Membership to update | 

### Return type

[**FindInvitations200ResponseInvitationsInner**](FindInvitations200ResponseInvitationsInner.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

