# \OperatingSystemVersionsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindOperatingSystemVersion**](OperatingSystemVersionsApi.md#FindOperatingSystemVersion) | **Get** /operating-system-versions | Retrieve all operating system versions



## FindOperatingSystemVersion

> []OperatingSystem FindOperatingSystemVersion(ctx).Execute()

Retrieve all operating system versions



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
    resp, r, err := api_client.OperatingSystemVersionsApi.FindOperatingSystemVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatingSystemVersionsApi.FindOperatingSystemVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOperatingSystemVersion`: []OperatingSystem
    fmt.Fprintf(os.Stdout, "Response from `OperatingSystemVersionsApi.FindOperatingSystemVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindOperatingSystemVersionRequest struct via the builder pattern


### Return type

[**[]OperatingSystem**](OperatingSystem.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

