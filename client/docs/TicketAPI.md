# \TicketAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketArchive**](TicketAPI.md#TicketArchive) | **Post** /api/v1/ticket/archive | Archive ticket
[**TicketClose**](TicketAPI.md#TicketClose) | **Post** /api/v1/ticket/close | Close ticket
[**TicketCreate**](TicketAPI.md#TicketCreate) | **Post** /api/v1/ticket/create | Create ticket
[**TicketDelete**](TicketAPI.md#TicketDelete) | **Delete** /api/v1/ticket/delete/{ticketId} | 
[**TicketDeleteMessage**](TicketAPI.md#TicketDeleteMessage) | **Delete** /api/v1/ticket/delete/message/{messageId} | 
[**TicketEdit**](TicketAPI.md#TicketEdit) | **Post** /api/v1/ticket/edit | Edit ticket
[**TicketEditMessage**](TicketAPI.md#TicketEditMessage) | **Post** /api/v1/ticket/edit/message | Edit ticket message
[**TicketList**](TicketAPI.md#TicketList) | **Get** /api/v1/ticket/list | Retrieve list of tickets
[**TicketMessages**](TicketAPI.md#TicketMessages) | **Get** /api/v1/ticket/{ticketId}/messages | 
[**TicketOpen**](TicketAPI.md#TicketOpen) | **Post** /api/v1/ticket/open | Open ticket
[**TicketReply**](TicketAPI.md#TicketReply) | **Post** /api/v1/ticket/reply | Reply message
[**TicketSetPriority**](TicketAPI.md#TicketSetPriority) | **Post** /api/v1/ticket/set-priority | Set priority
[**TicketTransfer**](TicketAPI.md#TicketTransfer) | **Post** /api/v1/ticket/transfer | Transfer ticket
[**TicketTransferList**](TicketAPI.md#TicketTransferList) | **Get** /api/v1/ticket/transfer/list | Retrieve organization managers



## TicketArchive

> TicketArchive(ctx).ArchiveTicketCommand(archiveTicketCommand).Execute()

Archive ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    archiveTicketCommand := *openapiclient.NewArchiveTicketCommand() // ArchiveTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketArchive(context.Background()).ArchiveTicketCommand(archiveTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveTicketCommand** | [**ArchiveTicketCommand**](ArchiveTicketCommand.md) |  | 

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


## TicketClose

> TicketClose(ctx).CloseTicketCommand(closeTicketCommand).Execute()

Close ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    closeTicketCommand := *openapiclient.NewCloseTicketCommand() // CloseTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketClose(context.Background()).CloseTicketCommand(closeTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **closeTicketCommand** | [**CloseTicketCommand**](CloseTicketCommand.md) |  | 

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


## TicketCreate

> TicketCreate(ctx).CreateTicketCommand(createTicketCommand).Execute()

Create ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    createTicketCommand := *openapiclient.NewCreateTicketCommand() // CreateTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketCreate(context.Background()).CreateTicketCommand(createTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTicketCommand** | [**CreateTicketCommand**](CreateTicketCommand.md) |  | 

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


## TicketDelete

> TicketDelete(ctx, ticketId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    ticketId := "ticketId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketDelete(context.Background(), ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketDeleteMessage

> TicketDeleteMessage(ctx, messageId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    messageId := "messageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketDeleteMessage(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketDeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketEdit

> TicketEdit(ctx).EditTicketCommand(editTicketCommand).Execute()

Edit ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    editTicketCommand := *openapiclient.NewEditTicketCommand() // EditTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketEdit(context.Background()).EditTicketCommand(editTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editTicketCommand** | [**EditTicketCommand**](EditTicketCommand.md) |  | 

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


## TicketEditMessage

> TicketEditMessage(ctx).EditArticleCommand(editArticleCommand).Execute()

Edit ticket message

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    editArticleCommand := *openapiclient.NewEditArticleCommand() // EditArticleCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketEditMessage(context.Background()).EditArticleCommand(editArticleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketEditMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketEditMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editArticleCommand** | [**EditArticleCommand**](EditArticleCommand.md) |  | 

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


## TicketList

> AllTicketsList TicketList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()

Retrieve list of tickets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    ticketId := "ticketId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketAPI.TicketList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketList`: AllTicketsList
    fmt.Fprintf(os.Stdout, "Response from `TicketAPI.TicketList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **ticketId** | **string** |  | 

### Return type

[**AllTicketsList**](AllTicketsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketMessages

> ArticleList TicketMessages(ctx, ticketId).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    ticketId := "ticketId_example" // string | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketAPI.TicketMessages(context.Background(), ticketId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketMessages`: ArticleList
    fmt.Fprintf(os.Stdout, "Response from `TicketAPI.TicketMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**ArticleList**](ArticleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketOpen

> TicketOpen(ctx).OpenTicketCommand(openTicketCommand).Execute()

Open ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    openTicketCommand := *openapiclient.NewOpenTicketCommand() // OpenTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketOpen(context.Background()).OpenTicketCommand(openTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openTicketCommand** | [**OpenTicketCommand**](OpenTicketCommand.md) |  | 

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


## TicketReply

> TicketReply(ctx).ReplyTicketCommand(replyTicketCommand).Execute()

Reply message

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    replyTicketCommand := *openapiclient.NewReplyTicketCommand() // ReplyTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketReply(context.Background()).ReplyTicketCommand(replyTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replyTicketCommand** | [**ReplyTicketCommand**](ReplyTicketCommand.md) |  | 

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


## TicketSetPriority

> TicketSetPriority(ctx).SetTicketPriorityCommand(setTicketPriorityCommand).Execute()

Set priority

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    setTicketPriorityCommand := *openapiclient.NewSetTicketPriorityCommand() // SetTicketPriorityCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketSetPriority(context.Background()).SetTicketPriorityCommand(setTicketPriorityCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketSetPriority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketSetPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setTicketPriorityCommand** | [**SetTicketPriorityCommand**](SetTicketPriorityCommand.md) |  | 

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


## TicketTransfer

> TicketTransfer(ctx).TransferTicketCommand(transferTicketCommand).Execute()

Transfer ticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    transferTicketCommand := *openapiclient.NewTransferTicketCommand() // TransferTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketAPI.TicketTransfer(context.Background()).TransferTicketCommand(transferTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferTicketCommand** | [**TransferTicketCommand**](TransferTicketCommand.md) |  | 

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


## TicketTransferList

> []TransferList TicketTransferList(ctx).OrganizationId(organizationId).Execute()

Retrieve organization managers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketAPI.TicketTransferList(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketAPI.TicketTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketTransferList`: []TransferList
    fmt.Fprintf(os.Stdout, "Response from `TicketAPI.TicketTransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]TransferList**](TransferList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

