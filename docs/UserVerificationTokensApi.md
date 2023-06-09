# \UserVerificationTokensApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumeVerificationRequest**](UserVerificationTokensApi.md#ConsumeVerificationRequest) | **Put** /verify-email | Verify a user using an email verification token
[**CreateValidationRequest**](UserVerificationTokensApi.md#CreateValidationRequest) | **Post** /verify-email | Create an email verification request



## ConsumeVerificationRequest

> ConsumeVerificationRequest(ctx).VerifyEmail(verifyEmail).Execute()

Verify a user using an email verification token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func main() {
    verifyEmail := *openapiclient.NewVerifyEmail("UserToken_example") // VerifyEmail | Email to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserVerificationTokensApi.ConsumeVerificationRequest(context.Background()).VerifyEmail(verifyEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserVerificationTokensApi.ConsumeVerificationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumeVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyEmail** | [**VerifyEmail**](VerifyEmail.md) | Email to create | 

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


## CreateValidationRequest

> CreateValidationRequest(ctx).Login(login).Execute()

Create an email verification request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func main() {
    login := "login_example" // string | Email for verification request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserVerificationTokensApi.CreateValidationRequest(context.Background()).Login(login).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserVerificationTokensApi.CreateValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValidationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **string** | Email for verification request | 

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

