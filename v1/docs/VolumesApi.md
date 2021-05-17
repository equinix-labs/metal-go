# \VolumesApi

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneVolume**](VolumesApi.md#CloneVolume) | **Post** /storage/{id}/clone | Clone volume/snapshot
[**CreateVolume**](VolumesApi.md#CreateVolume) | **Post** /projects/{id}/storage | Create a volume
[**CreateVolumeAttachment**](VolumesApi.md#CreateVolumeAttachment) | **Post** /storage/{id}/attachments | Attach your volume
[**CreateVolumeSnapshotPolicy**](VolumesApi.md#CreateVolumeSnapshotPolicy) | **Post** /storage/{id}/snapshot-policies | Create a volume snapshot policy
[**DeleteVolume**](VolumesApi.md#DeleteVolume) | **Delete** /storage/{id} | Delete the volume
[**DeleteVolumeAttachment**](VolumesApi.md#DeleteVolumeAttachment) | **Delete** /storage/attachments/{id} | Detach volume
[**DeleteVolumeSnapshot**](VolumesApi.md#DeleteVolumeSnapshot) | **Delete** /storage/{volume_id}/snapshots/{id} | Delete volume snapshot
[**DeleteVolumeSnapshotPolicy**](VolumesApi.md#DeleteVolumeSnapshotPolicy) | **Delete** /storage/snapshot-policies/{id} | Delete the volume snapshot policy
[**FindVolumeAttachmentById**](VolumesApi.md#FindVolumeAttachmentById) | **Get** /storage/attachments/{id} | Retrieve an attachment
[**FindVolumeAttachments**](VolumesApi.md#FindVolumeAttachments) | **Get** /storage/{id}/attachments | Retrieve all volume attachment
[**FindVolumeById**](VolumesApi.md#FindVolumeById) | **Get** /storage/{id} | Retrieve a volume
[**FindVolumeCustomdata**](VolumesApi.md#FindVolumeCustomdata) | **Get** /storage/{id}/customdata | Retrieve the custom metadata of a storage volume
[**FindVolumeSnapshots**](VolumesApi.md#FindVolumeSnapshots) | **Get** /storage/{id}/snapshots | Retrieve all volume snapshot
[**FindVolumes**](VolumesApi.md#FindVolumes) | **Get** /projects/{id}/storage | Retrieve all volumes
[**RestoreVolume**](VolumesApi.md#RestoreVolume) | **Post** /storage/{id}/restore | Restore volume
[**UpdateVolume**](VolumesApi.md#UpdateVolume) | **Put** /storage/{id} | Update the volume
[**UpdateVolumeSnapshotPolicy**](VolumesApi.md#UpdateVolumeSnapshotPolicy) | **Put** /storage/snapshot-policies/{id} | Update the volume snapshot policy



## CloneVolume

> Volume CloneVolume(ctx, id).SnapshotTimestamp(snapshotTimestamp).Execute()

Clone volume/snapshot



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
    id := TODO // string | Volume UUID
    snapshotTimestamp := "snapshotTimestamp_example" // string | snapshot timestamp (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CloneVolume(context.Background(), id).SnapshotTimestamp(snapshotTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CloneVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CloneVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotTimestamp** | **string** | snapshot timestamp | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolume

> Volume CreateVolume(ctx, id).Volume(volume).Execute()

Create a volume



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
    volume := *openapiclient.NewVolumeCreateInput("Facility_example", "Plan_example", int32(123)) // VolumeCreateInput | Volume to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CreateVolume(context.Background(), id).Volume(volume).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CreateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**VolumeCreateInput**](VolumeCreateInput.md) | Volume to create | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeAttachment

> VolumeAttachment CreateVolumeAttachment(ctx, id).Attachment(attachment).Execute()

Attach your volume



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
    id := TODO // string | Volume UUID
    attachment := *openapiclient.NewVolumeAttachmentInput("DeviceId_example") // VolumeAttachmentInput | Device to attach

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CreateVolumeAttachment(context.Background(), id).Attachment(attachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolumeAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeAttachment`: VolumeAttachment
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CreateVolumeAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachment** | [**VolumeAttachmentInput**](VolumeAttachmentInput.md) | Device to attach | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeSnapshotPolicy

> SnapshotPolicy CreateVolumeSnapshotPolicy(ctx, id).SnapshotFrequency(snapshotFrequency).SnapshotCount(snapshotCount).Execute()

Create a volume snapshot policy



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
    id := TODO // string | Volume UUID
    snapshotFrequency := "snapshotFrequency_example" // string | Snapshot frequency
    snapshotCount := int32(56) // int32 | Snapshot count (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CreateVolumeSnapshotPolicy(context.Background(), id).SnapshotFrequency(snapshotFrequency).SnapshotCount(snapshotCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolumeSnapshotPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeSnapshotPolicy`: SnapshotPolicy
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CreateVolumeSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotFrequency** | **string** | Snapshot frequency | 
 **snapshotCount** | **int32** | Snapshot count | 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolume(ctx, id).Execute()

Delete the volume



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
    id := TODO // string | Volume UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DeleteVolume(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## DeleteVolumeAttachment

> DeleteVolumeAttachment(ctx, id).Execute()

Detach volume



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
    id := TODO // string | Attachment UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DeleteVolumeAttachment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolumeAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Attachment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeAttachmentRequest struct via the builder pattern


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


## DeleteVolumeSnapshot

> DeleteVolumeSnapshot(ctx, volumeId, id).Execute()

Delete volume snapshot



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
    volumeId := TODO // string | Volume UUID
    id := TODO // string | Snapshot UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DeleteVolumeSnapshot(context.Background(), volumeId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolumeSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | [**string**](.md) | Volume UUID | 
**id** | [**string**](.md) | Snapshot UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotRequest struct via the builder pattern


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


## DeleteVolumeSnapshotPolicy

> DeleteVolumeSnapshotPolicy(ctx, id).Execute()

Delete the volume snapshot policy



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
    id := TODO // string | Snapshot Policy UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DeleteVolumeSnapshotPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolumeSnapshotPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Snapshot Policy UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotPolicyRequest struct via the builder pattern


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


## FindVolumeAttachmentById

> VolumeAttachment FindVolumeAttachmentById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve an attachment



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
    id := TODO // string | Attachment UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumeAttachmentById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumeAttachmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVolumeAttachmentById`: VolumeAttachment
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.FindVolumeAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Attachment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumeAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVolumeAttachments

> VolumeAttachmentList FindVolumeAttachments(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve all volume attachment



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
    id := TODO // string | Volume UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumeAttachments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumeAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVolumeAttachments`: VolumeAttachmentList
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.FindVolumeAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumeAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**VolumeAttachmentList**](VolumeAttachmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVolumeById

> Volume FindVolumeById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a volume



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
    id := TODO // string | Volume UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumeById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVolumeById`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.FindVolumeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVolumeCustomdata

> FindVolumeCustomdata(ctx, id).Execute()

Retrieve the custom metadata of a storage volume



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
    id := TODO // string | Storage Volume UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumeCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumeCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Storage Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumeCustomdataRequest struct via the builder pattern


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


## FindVolumeSnapshots

> VolumeSnapshotList FindVolumeSnapshots(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve all volume snapshot



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
    id := TODO // string | Volume UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumeSnapshots(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumeSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVolumeSnapshots`: VolumeSnapshotList
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.FindVolumeSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumeSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**VolumeSnapshotList**](VolumeSnapshotList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVolumes

> VolumeList FindVolumes(ctx, id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all volumes



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
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.FindVolumes(context.Background(), id).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.FindVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVolumes`: VolumeList
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.FindVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**VolumeList**](VolumeList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVolume

> Volume RestoreVolume(ctx, id).RestorePoint(restorePoint).Execute()

Restore volume



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
    id := TODO // string | Volume UUID
    restorePoint := "restorePoint_example" // string | restore point

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.RestoreVolume(context.Background(), id).RestorePoint(restorePoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.RestoreVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.RestoreVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restorePoint** | **string** | restore point | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> Volume UpdateVolume(ctx, id).Volume(volume).Execute()

Update the volume



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
    id := TODO // string | Volume UUID
    volume := *openapiclient.NewVolumeUpdateInput() // VolumeUpdateInput | Volume to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.UpdateVolume(context.Background(), id).Volume(volume).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.UpdateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Volume UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**VolumeUpdateInput**](VolumeUpdateInput.md) | Volume to update | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolumeSnapshotPolicy

> SnapshotPolicy UpdateVolumeSnapshotPolicy(ctx, id).SnapshotFrequency(snapshotFrequency).SnapshotCount(snapshotCount).Execute()

Update the volume snapshot policy



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
    id := TODO // string | Snapshot Policy UUID
    snapshotFrequency := "snapshotFrequency_example" // string | Snapshot frequency
    snapshotCount := int32(56) // int32 | Snapshot count (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.UpdateVolumeSnapshotPolicy(context.Background(), id).SnapshotFrequency(snapshotFrequency).SnapshotCount(snapshotCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.UpdateVolumeSnapshotPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolumeSnapshotPolicy`: SnapshotPolicy
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.UpdateVolumeSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Snapshot Policy UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotFrequency** | **string** | Snapshot frequency | 
 **snapshotCount** | **int32** | Snapshot count | 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

