# \DatasetServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatasetServiceCreateLogWriter**](DatasetServiceApi.md#DatasetServiceCreateLogWriter) | **Post** /v1/projects/{projectId}/log_writers | CreateLogWriter creates a new log writer deployment as specified by the details given in the request body.
[**DatasetServiceCreateObservationService**](DatasetServiceApi.md#DatasetServiceCreateObservationService) | **Post** /v1/projects/{projectId}/observation_services | CreateObservationService creates a new observation service deployment as specified by the details given in the request body.
[**DatasetServiceGetLog**](DatasetServiceApi.md#DatasetServiceGetLog) | **Get** /v1/projects/{projectId}/logs/{id} | GetLog return details of a log.
[**DatasetServiceGetLogWriter**](DatasetServiceApi.md#DatasetServiceGetLogWriter) | **Get** /v1/projects/{projectId}/log_writers/{id} | GetLogWriter return details of the log writer deployment.
[**DatasetServiceGetObservationService**](DatasetServiceApi.md#DatasetServiceGetObservationService) | **Get** /v1/projects/{projectId}/observation_services/{id} | GetObservationService return details of the observation service deployment.
[**DatasetServiceListLogWriters**](DatasetServiceApi.md#DatasetServiceListLogWriters) | **Get** /v1/projects/{projectId}/log_writers | ListLogWriters return paginated list of log writers under a project and filtered by query string.
[**DatasetServiceListLogs**](DatasetServiceApi.md#DatasetServiceListLogs) | **Get** /v1/projects/{projectId}/logs | ListLogs return paginated list of logs under a project and filtered by query string.
[**DatasetServiceListObservationServices**](DatasetServiceApi.md#DatasetServiceListObservationServices) | **Get** /v1/projects/{projectId}/observation_services | ListObservationServices return paginated list of observation services under a project and filtered by query string.
[**DatasetServiceUpdateLogWriter**](DatasetServiceApi.md#DatasetServiceUpdateLogWriter) | **Put** /v1/projects/{projectId}/log_writers/{id} | UpdateLogWriter updates an existing log writer deployment as specified by the details given in the request body.
[**DatasetServiceUpdateObservationService**](DatasetServiceApi.md#DatasetServiceUpdateObservationService) | **Put** /v1/projects/{projectId}/observation_services/{id} | UpdateObservationService updates an existing observation service deployment as specified by the details given in the request body.



## DatasetServiceCreateLogWriter

> V1CreateLogWriterResponse DatasetServiceCreateLogWriter(ctx, projectId).Body(body).Execute()

CreateLogWriter creates a new log writer deployment as specified by the details given in the request body.

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
    projectId := "projectId_example" // string | The CaraML project ID to write log resource from.
    body := *openapiclient.NewDatasetServiceCreateLogWriterRequest() // DatasetServiceCreateLogWriterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceCreateLogWriter(context.Background(), projectId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceCreateLogWriter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceCreateLogWriter`: V1CreateLogWriterResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceCreateLogWriter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to write log resource from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceCreateLogWriterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DatasetServiceCreateLogWriterRequest**](DatasetServiceCreateLogWriterRequest.md) |  | 

### Return type

[**V1CreateLogWriterResponse**](V1CreateLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceCreateObservationService

> V1CreateObservationServiceResponse DatasetServiceCreateObservationService(ctx, projectId).Body(body).Execute()

CreateObservationService creates a new observation service deployment as specified by the details given in the request body.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.
    body := *openapiclient.NewDatasetServiceCreateObservationServiceRequest() // DatasetServiceCreateObservationServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceCreateObservationService(context.Background(), projectId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceCreateObservationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceCreateObservationService`: V1CreateObservationServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceCreateObservationService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceCreateObservationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DatasetServiceCreateObservationServiceRequest**](DatasetServiceCreateObservationServiceRequest.md) |  | 

### Return type

[**V1CreateObservationServiceResponse**](V1CreateObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceGetLog

> V1GetLogResponse DatasetServiceGetLog(ctx, projectId, id).Execute()

GetLog return details of a log.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.
    id := "id_example" // string | The ID of the log resource to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceGetLog(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceGetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceGetLog`: V1GetLogResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceGetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 
**id** | **string** | The ID of the log resource to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1GetLogResponse**](V1GetLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceGetLogWriter

> V1GetLogWriterResponse DatasetServiceGetLogWriter(ctx, projectId, id).Execute()

GetLogWriter return details of the log writer deployment.

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
    projectId := "projectId_example" // string | The CaraML project ID to write log resource from.
    id := "id_example" // string | The ID of the Log Writer resource to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceGetLogWriter(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceGetLogWriter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceGetLogWriter`: V1GetLogWriterResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceGetLogWriter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to write log resource from. | 
**id** | **string** | The ID of the Log Writer resource to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceGetLogWriterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1GetLogWriterResponse**](V1GetLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceGetObservationService

> V1GetObservationServiceResponse DatasetServiceGetObservationService(ctx, projectId, id).Execute()

GetObservationService return details of the observation service deployment.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.
    id := "id_example" // string | The ID of the Observation Service resource to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceGetObservationService(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceGetObservationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceGetObservationService`: V1GetObservationServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceGetObservationService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 
**id** | **string** | The ID of the Observation Service resource to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceGetObservationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1GetObservationServiceResponse**](V1GetObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceListLogWriters

> V1ListLogWritersResponse DatasetServiceListLogWriters(ctx, projectId).Execute()

ListLogWriters return paginated list of log writers under a project and filtered by query string.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceListLogWriters(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceListLogWriters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceListLogWriters`: V1ListLogWritersResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceListLogWriters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceListLogWritersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListLogWritersResponse**](V1ListLogWritersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceListLogs

> V1ListLogsResponse DatasetServiceListLogs(ctx, projectId).Execute()

ListLogs return paginated list of logs under a project and filtered by query string.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceListLogs(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceListLogs`: V1ListLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceListLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListLogsResponse**](V1ListLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceListObservationServices

> V1ListObservationServicesResponse DatasetServiceListObservationServices(ctx, projectId).Execute()

ListObservationServices return paginated list of observation services under a project and filtered by query string.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceListObservationServices(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceListObservationServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceListObservationServices`: V1ListObservationServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceListObservationServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceListObservationServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListObservationServicesResponse**](V1ListObservationServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceUpdateLogWriter

> V1UpdateLogWriterResponse DatasetServiceUpdateLogWriter(ctx, projectId, id).Body(body).Execute()

UpdateLogWriter updates an existing log writer deployment as specified by the details given in the request body.

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
    projectId := "projectId_example" // string | The CaraML project ID to write log resource from.
    id := "id_example" // string | The ID of Log Writer to update.
    body := *openapiclient.NewDatasetServiceUpdateLogWriterRequest() // DatasetServiceUpdateLogWriterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceUpdateLogWriter(context.Background(), projectId, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceUpdateLogWriter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceUpdateLogWriter`: V1UpdateLogWriterResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceUpdateLogWriter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to write log resource from. | 
**id** | **string** | The ID of Log Writer to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceUpdateLogWriterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DatasetServiceUpdateLogWriterRequest**](DatasetServiceUpdateLogWriterRequest.md) |  | 

### Return type

[**V1UpdateLogWriterResponse**](V1UpdateLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasetServiceUpdateObservationService

> V1UpdateObservationServiceResponse DatasetServiceUpdateObservationService(ctx, projectId, id).Body(body).Execute()

UpdateObservationService updates an existing observation service deployment as specified by the details given in the request body.

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
    projectId := "projectId_example" // string | The CaraML project ID to retrieve log resource from.
    id := "id_example" // string | The ID of Observation Service to update.
    body := *openapiclient.NewDatasetServiceUpdateObservationServiceRequest() // DatasetServiceUpdateObservationServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetServiceApi.DatasetServiceUpdateObservationService(context.Background(), projectId, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetServiceApi.DatasetServiceUpdateObservationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasetServiceUpdateObservationService`: V1UpdateObservationServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetServiceApi.DatasetServiceUpdateObservationService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The CaraML project ID to retrieve log resource from. | 
**id** | **string** | The ID of Observation Service to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasetServiceUpdateObservationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DatasetServiceUpdateObservationServiceRequest**](DatasetServiceUpdateObservationServiceRequest.md) |  | 

### Return type

[**V1UpdateObservationServiceResponse**](V1UpdateObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

