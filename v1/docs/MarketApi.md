# \MarketApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindMetroSpotMarketPrices**](MarketApi.md#FindMetroSpotMarketPrices) | **Get** /market/spot/prices/metros | Get current spot market prices for metros
[**FindSpotMarketPrices**](MarketApi.md#FindSpotMarketPrices) | **Get** /market/spot/prices | Get current spot market prices
[**FindSpotMarketPricesHistory**](MarketApi.md#FindSpotMarketPricesHistory) | **Get** /market/spot/prices/history | Get spot market prices for a given period of time



## FindMetroSpotMarketPrices

> SpotMarketPricesPerMetroList FindMetroSpotMarketPrices(ctx).Metro(metro).Plan(plan).Execute()

Get current spot market prices for metros



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
    metro := "metro_example" // string | Metro to filter spot market prices (optional)
    plan := "plan_example" // string | Plan to filter spot market prices (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FindMetroSpotMarketPrices(context.Background()).Metro(metro).Plan(plan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FindMetroSpotMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetroSpotMarketPrices`: SpotMarketPricesPerMetroList
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FindMetroSpotMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindMetroSpotMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metro** | **string** | Metro to filter spot market prices | 
 **plan** | **string** | Plan to filter spot market prices | 

### Return type

[**SpotMarketPricesPerMetroList**](SpotMarketPricesPerMetroList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketPrices

> SpotMarketPricesList FindSpotMarketPrices(ctx).Facility(facility).Plan(plan).Execute()

Get current spot market prices



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
    facility := "facility_example" // string | Facility to check spot market prices (optional)
    plan := "plan_example" // string | Plan to check spot market prices (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FindSpotMarketPrices(context.Background()).Facility(facility).Plan(plan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FindSpotMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketPrices`: SpotMarketPricesList
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FindSpotMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string** | Facility to check spot market prices | 
 **plan** | **string** | Plan to check spot market prices | 

### Return type

[**SpotMarketPricesList**](SpotMarketPricesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketPricesHistory

> SpotPricesHistoryReport FindSpotMarketPricesHistory(ctx).Facility(facility).Plan(plan).From(from).Until(until).Metro(metro).Execute()

Get spot market prices for a given period of time



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
    facility := "facility_example" // string | Facility to check spot market prices
    plan := "plan_example" // string | Plan to check spot market prices
    from := "from_example" // string | Timestamp from range
    until := "until_example" // string | Timestamp to range
    metro := "metro_example" // string | Metro to check spot market price history (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FindSpotMarketPricesHistory(context.Background()).Facility(facility).Plan(plan).From(from).Until(until).Metro(metro).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FindSpotMarketPricesHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketPricesHistory`: SpotPricesHistoryReport
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FindSpotMarketPricesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketPricesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string** | Facility to check spot market prices | 
 **plan** | **string** | Plan to check spot market prices | 
 **from** | **string** | Timestamp from range | 
 **until** | **string** | Timestamp to range | 
 **metro** | **string** | Metro to check spot market price history | 

### Return type

[**SpotPricesHistoryReport**](SpotPricesHistoryReport.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

