# \GoogleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GooglecloudBillingAccountList**](GoogleAPI.md#GooglecloudBillingAccountList) | **Post** /api/v1/googlecloud/billing-accounts | 
[**GooglecloudCreate**](GoogleAPI.md#GooglecloudCreate) | **Post** /api/v1/googlecloud/create | 
[**GooglecloudList**](GoogleAPI.md#GooglecloudList) | **Get** /api/v1/googlecloud/list | Retrieve list of google cloud credentials
[**GooglecloudRegionList**](GoogleAPI.md#GooglecloudRegionList) | **Post** /api/v1/googlecloud/regions | 
[**GooglecloudZoneList**](GoogleAPI.md#GooglecloudZoneList) | **Post** /api/v1/googlecloud/zones/{region} | 



## GooglecloudBillingAccountList

> []CommonStringBasedDropdownDto GooglecloudBillingAccountList(ctx).Config(config).Execute()



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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleAPI.GooglecloudBillingAccountList(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudBillingAccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudBillingAccountList`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudBillingAccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudBillingAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudCreate

> ApiResponse GooglecloudCreate(ctx).ImportProject(importProject).AzCount(azCount).Name(name).FolderId(folderId).BillingAccountId(billingAccountId).Region(region).OrganizationId(organizationId).Config(config).Execute()



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
    importProject := true // bool | 
    azCount := int32(56) // int32 | 
    name := "name_example" // string |  (optional)
    folderId := "folderId_example" // string |  (optional)
    billingAccountId := "billingAccountId_example" // string |  (optional)
    region := "region_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleAPI.GooglecloudCreate(context.Background()).ImportProject(importProject).AzCount(azCount).Name(name).FolderId(folderId).BillingAccountId(billingAccountId).Region(region).OrganizationId(organizationId).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importProject** | **bool** |  | 
 **azCount** | **int32** |  | 
 **name** | **string** |  | 
 **folderId** | **string** |  | 
 **billingAccountId** | **string** |  | 
 **region** | **string** |  | 
 **organizationId** | **int32** |  | 
 **config** | ***os.File** |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudList

> GoogleCredentialList GooglecloudList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of google cloud credentials

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
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleAPI.GooglecloudList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudList`: GoogleCredentialList
    fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**GoogleCredentialList**](GoogleCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudRegionList

> []string GooglecloudRegionList(ctx).File(file).Execute()



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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleAPI.GooglecloudRegionList(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudRegionList`: []string
    fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudZoneList

> AzResult GooglecloudZoneList(ctx, region).CloudId(cloudId).Config(config).Execute()



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
    region := "region_example" // string | 
    cloudId := int32(56) // int32 |  (optional)
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleAPI.GooglecloudZoneList(context.Background(), region).CloudId(cloudId).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudZoneList`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudZoneList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudId** | **int32** |  | 
 **config** | ***os.File** |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

