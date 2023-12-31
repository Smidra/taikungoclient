# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesAwsCommonImages**](ImagesAPI.md#ImagesAwsCommonImages) | **Get** /api/v1/images/aws/common/{cloudId} | Commonly used aws images
[**ImagesAwsImagesList**](ImagesAPI.md#ImagesAwsImagesList) | **Post** /api/v1/images/aws | Retrieve aws images
[**ImagesAwsPersonalImages**](ImagesAPI.md#ImagesAwsPersonalImages) | **Get** /api/v1/images/aws/personal/{cloudId} | Aws personal images
[**ImagesAzureCommonImages**](ImagesAPI.md#ImagesAzureCommonImages) | **Get** /api/v1/images/azure/common/{cloudId} | Commonly used azure images
[**ImagesAzureImages**](ImagesAPI.md#ImagesAzureImages) | **Get** /api/v1/images/azure/{cloudId}/{publisherName}/{offer}/{sku} | 
[**ImagesAzurePersonalImages**](ImagesAPI.md#ImagesAzurePersonalImages) | **Get** /api/v1/images/azure/personal/{cloudId} | Azure personal images
[**ImagesBindImagesToProject**](ImagesAPI.md#ImagesBindImagesToProject) | **Post** /api/v1/images/bind | Bind images to project
[**ImagesCommonGoogleImages**](ImagesAPI.md#ImagesCommonGoogleImages) | **Get** /api/v1/images/google/common/{cloudId} | Commonly used google images
[**ImagesGoogleImages**](ImagesAPI.md#ImagesGoogleImages) | **Get** /api/v1/images/google/{cloudId}/{type} | 
[**ImagesImageDetails**](ImagesAPI.md#ImagesImageDetails) | **Post** /api/v1/images/details | Get image details
[**ImagesOpenstackImages**](ImagesAPI.md#ImagesOpenstackImages) | **Get** /api/v1/images/openstack/{cloudId} | Retrieve openstack images
[**ImagesProxmoxImages**](ImagesAPI.md#ImagesProxmoxImages) | **Get** /api/v1/images/proxmox/{cloudId} | Retrieve proxmox images
[**ImagesSelectedImagesForProject**](ImagesAPI.md#ImagesSelectedImagesForProject) | **Get** /api/v1/images/projects/list | Retrieve selected images for projects
[**ImagesTanzuImages**](ImagesAPI.md#ImagesTanzuImages) | **Get** /api/v1/images/tanzu/{cloudId} | Retrieve tanzu images
[**ImagesUnbindImagesFromProject**](ImagesAPI.md#ImagesUnbindImagesFromProject) | **Post** /api/v1/images/unbind | Unbind images from project



## ImagesAwsCommonImages

> []AwsOwnerDetails ImagesAwsCommonImages(ctx, cloudId).Execute()

Commonly used aws images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAwsCommonImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAwsCommonImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsCommonImages`: []AwsOwnerDetails
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAwsCommonImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsCommonImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AwsOwnerDetails**](AwsOwnerDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAwsImagesList

> AwsImagesPostList ImagesAwsImagesList(ctx).AwsImagesPostListCommand(awsImagesPostListCommand).Execute()

Retrieve aws images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    awsImagesPostListCommand := *openapiclient.NewAwsImagesPostListCommand() // AwsImagesPostListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAwsImagesList(context.Background()).AwsImagesPostListCommand(awsImagesPostListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAwsImagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsImagesList`: AwsImagesPostList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAwsImagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsImagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsImagesPostListCommand** | [**AwsImagesPostListCommand**](AwsImagesPostListCommand.md) |  | 

### Return type

[**AwsImagesPostList**](AwsImagesPostList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAwsPersonalImages

> []CommonStringBasedDropdownDto ImagesAwsPersonalImages(ctx, cloudId).Search(search).Execute()

Aws personal images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAwsPersonalImages(context.Background(), cloudId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAwsPersonalImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsPersonalImages`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAwsPersonalImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsPersonalImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAzureCommonImages

> []AzurePublisherDetails ImagesAzureCommonImages(ctx, cloudId).Execute()

Commonly used azure images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAzureCommonImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAzureCommonImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAzureCommonImages`: []AzurePublisherDetails
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAzureCommonImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAzureCommonImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AzurePublisherDetails**](AzurePublisherDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAzureImages

> AzureImageList ImagesAzureImages(ctx, cloudId, publisherName, offer, sku).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Latest(latest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    publisherName := "publisherName_example" // string | 
    offer := "offer_example" // string | 
    sku := "sku_example" // string | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    latest := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAzureImages(context.Background(), cloudId, publisherName, offer, sku).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Latest(latest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAzureImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAzureImages`: AzureImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAzureImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisherName** | **string** |  | 
**offer** | **string** |  | 
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAzureImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **latest** | **bool** |  | [default to false]

### Return type

[**AzureImageList**](AzureImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAzurePersonalImages

> []CommonStringBasedDropdownDto ImagesAzurePersonalImages(ctx, cloudId).Execute()

Azure personal images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesAzurePersonalImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesAzurePersonalImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAzurePersonalImages`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesAzurePersonalImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAzurePersonalImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesBindImagesToProject

> ImagesBindImagesToProject(ctx).BindImageToProjectCommand(bindImageToProjectCommand).Execute()

Bind images to project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    bindImageToProjectCommand := *openapiclient.NewBindImageToProjectCommand() // BindImageToProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesAPI.ImagesBindImagesToProject(context.Background()).BindImageToProjectCommand(bindImageToProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesBindImagesToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesBindImagesToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindImageToProjectCommand** | [**BindImageToProjectCommand**](BindImageToProjectCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesCommonGoogleImages

> []GoogleOwnerDetails ImagesCommonGoogleImages(ctx, cloudId).Execute()

Commonly used google images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesCommonGoogleImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesCommonGoogleImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesCommonGoogleImages`: []GoogleOwnerDetails
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesCommonGoogleImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesCommonGoogleImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GoogleOwnerDetails**](GoogleOwnerDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesGoogleImages

> GoogleImageList ImagesGoogleImages(ctx, cloudId, type_).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    type_ := "type__example" // string | 
    latest := true // bool | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesGoogleImages(context.Background(), cloudId, type_).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesGoogleImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesGoogleImages`: GoogleImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesGoogleImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesGoogleImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **latest** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**GoogleImageList**](GoogleImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesImageDetails

> string ImagesImageDetails(ctx).ImageByIdCommand(imageByIdCommand).Execute()

Get image details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    imageByIdCommand := *openapiclient.NewImageByIdCommand() // ImageByIdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesImageDetails(context.Background()).ImageByIdCommand(imageByIdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesImageDetails`: string
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesImageDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageByIdCommand** | [**ImageByIdCommand**](ImageByIdCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesOpenstackImages

> OpenstackImageList ImagesOpenstackImages(ctx, cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Personal(personal).Execute()

Retrieve openstack images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    personal := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesOpenstackImages(context.Background(), cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Personal(personal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesOpenstackImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesOpenstackImages`: OpenstackImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesOpenstackImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesOpenstackImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **personal** | **bool** |  | [default to false]

### Return type

[**OpenstackImageList**](OpenstackImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesProxmoxImages

> ProxmoxImageList ImagesProxmoxImages(ctx, cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve proxmox images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesProxmoxImages(context.Background(), cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesProxmoxImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesProxmoxImages`: ProxmoxImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesProxmoxImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesProxmoxImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ProxmoxImageList**](ProxmoxImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesSelectedImagesForProject

> BoundImagesForProjectsList ImagesSelectedImagesForProject(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).Execute()

Retrieve selected images for projects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesSelectedImagesForProject(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesSelectedImagesForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesSelectedImagesForProject`: BoundImagesForProjectsList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesSelectedImagesForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesSelectedImagesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**BoundImagesForProjectsList**](BoundImagesForProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTanzuImages

> TanzuImageList ImagesTanzuImages(ctx, cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve tanzu images

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.ImagesTanzuImages(context.Background(), cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesTanzuImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTanzuImages`: TanzuImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesTanzuImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTanzuImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**TanzuImageList**](TanzuImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesUnbindImagesFromProject

> ImagesUnbindImagesFromProject(ctx).DeleteImageFromProjectCommand(deleteImageFromProjectCommand).Execute()

Unbind images from project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    deleteImageFromProjectCommand := *openapiclient.NewDeleteImageFromProjectCommand() // DeleteImageFromProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesAPI.ImagesUnbindImagesFromProject(context.Background()).DeleteImageFromProjectCommand(deleteImageFromProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesUnbindImagesFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesUnbindImagesFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteImageFromProjectCommand** | [**DeleteImageFromProjectCommand**](DeleteImageFromProjectCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

