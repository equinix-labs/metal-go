# \OtpsApi

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindEnsureOtp**](OtpsApi.md#FindEnsureOtp) | **Post** /user/otp/verify/{otp} | Verify user by providing an OTP
[**FindRecoveryCodes**](OtpsApi.md#FindRecoveryCodes) | **Get** /user/otp/recovery-codes | Retrieve my recovery codes
[**ReceiveCodes**](OtpsApi.md#ReceiveCodes) | **Post** /user/otp/sms/receive | Receive an OTP per sms
[**RegenerateCodes**](OtpsApi.md#RegenerateCodes) | **Post** /user/otp/recovery-codes | Generate new recovery codes



## FindEnsureOtp

> FindEnsureOtp(ctx, otp).Execute()

Verify user by providing an OTP



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
    otp := "otp_example" // string | OTP

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OtpsApi.FindEnsureOtp(context.Background(), otp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpsApi.FindEnsureOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otp** | **string** | OTP | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEnsureOtpRequest struct via the builder pattern


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


## FindRecoveryCodes

> RecoveryCodeList FindRecoveryCodes(ctx).Execute()

Retrieve my recovery codes



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OtpsApi.FindRecoveryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpsApi.FindRecoveryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindRecoveryCodes`: RecoveryCodeList
    fmt.Fprintf(os.Stdout, "Response from `OtpsApi.FindRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindRecoveryCodesRequest struct via the builder pattern


### Return type

[**RecoveryCodeList**](RecoveryCodeList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceiveCodes

> ReceiveCodes(ctx).Execute()

Receive an OTP per sms



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OtpsApi.ReceiveCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpsApi.ReceiveCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveCodesRequest struct via the builder pattern


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


## RegenerateCodes

> RecoveryCodeList RegenerateCodes(ctx).Execute()

Generate new recovery codes



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OtpsApi.RegenerateCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpsApi.RegenerateCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegenerateCodes`: RecoveryCodeList
    fmt.Fprintf(os.Stdout, "Response from `OtpsApi.RegenerateCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateCodesRequest struct via the builder pattern


### Return type

[**RecoveryCodeList**](RecoveryCodeList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

