# openapi_client.DatasetServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dataset_service_create_log_writer**](DatasetServiceApi.md#dataset_service_create_log_writer) | **POST** /v1/projects/{projectId}/log_writers | CreateLogWriter creates a new log writer deployment as specified by the details given in the request body.
[**dataset_service_create_observation_service**](DatasetServiceApi.md#dataset_service_create_observation_service) | **POST** /v1/projects/{projectId}/observation_services | CreateObservationService creates a new observation service deployment as specified by the details given in the request body.
[**dataset_service_get_log**](DatasetServiceApi.md#dataset_service_get_log) | **GET** /v1/projects/{projectId}/logs/{id} | GetLog return details of a log.
[**dataset_service_get_log_writer**](DatasetServiceApi.md#dataset_service_get_log_writer) | **GET** /v1/projects/{projectId}/log_writers/{id} | GetLogWriter return details of the log writer deployment.
[**dataset_service_get_observation_service**](DatasetServiceApi.md#dataset_service_get_observation_service) | **GET** /v1/projects/{projectId}/observation_services/{id} | GetObservationService return details of the observation service deployment.
[**dataset_service_list_log_writers**](DatasetServiceApi.md#dataset_service_list_log_writers) | **GET** /v1/projects/{projectId}/log_writers | ListLogWriters return paginated list of log writers under a project and filtered by query string.
[**dataset_service_list_logs**](DatasetServiceApi.md#dataset_service_list_logs) | **GET** /v1/projects/{projectId}/logs | ListLogs return paginated list of logs under a project and filtered by query string.
[**dataset_service_list_observation_services**](DatasetServiceApi.md#dataset_service_list_observation_services) | **GET** /v1/projects/{projectId}/observation_services | ListObservationServices return paginated list of observation services under a project and filtered by query string.
[**dataset_service_update_log_writer**](DatasetServiceApi.md#dataset_service_update_log_writer) | **PUT** /v1/projects/{projectId}/log_writers/{id} | UpdateLogWriter updates an existing log writer deployment as specified by the details given in the request body.
[**dataset_service_update_observation_service**](DatasetServiceApi.md#dataset_service_update_observation_service) | **PUT** /v1/projects/{projectId}/observation_services/{id} | UpdateObservationService updates an existing observation service deployment as specified by the details given in the request body.


# **dataset_service_create_log_writer**
> V1CreateLogWriterResponse dataset_service_create_log_writer(project_id, body)

CreateLogWriter creates a new log writer deployment as specified by the details given in the request body.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_create_log_writer_response import V1CreateLogWriterResponse
from openapi_client.model.dataset_service_create_log_writer_request import DatasetServiceCreateLogWriterRequest
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to write log resource from.
    body = DatasetServiceCreateLogWriterRequest(
        log_writer=V1LogWriter(
            type=V1LogWriterType("LOG_WRITER_TYPE_UNSPECIFIED"),
            fluentd_config=V1FluentdConfig(
                type=V1FluentdOutputType("FLUENTD_OUTPUT_TYPE_UNSPECIFIED"),
                host="host_example",
                port=1,
                tag="tag_example",
                config=V1FluentdOutputBQConfig(
                    project="project_example",
                    dataset="dataset_example",
                    table="table_example",
                ),
            ),
        ),
    ) # DatasetServiceCreateLogWriterRequest | 

    # example passing only required values which don't have defaults set
    try:
        # CreateLogWriter creates a new log writer deployment as specified by the details given in the request body.
        api_response = api_instance.dataset_service_create_log_writer(project_id, body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_create_log_writer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to write log resource from. |
 **body** | [**DatasetServiceCreateLogWriterRequest**](DatasetServiceCreateLogWriterRequest.md)|  |

### Return type

[**V1CreateLogWriterResponse**](V1CreateLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_create_observation_service**
> V1CreateObservationServiceResponse dataset_service_create_observation_service(project_id, body)

CreateObservationService creates a new observation service deployment as specified by the details given in the request body.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.v1_create_observation_service_response import V1CreateObservationServiceResponse
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.dataset_service_create_observation_service_request import DatasetServiceCreateObservationServiceRequest
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.
    body = DatasetServiceCreateObservationServiceRequest(
        observation_service=V1ObservationServiceConfig(
            id="id_example",
            source=V1ObservationServiceDataSource(
                type=V1ObservationServiceDataSourceType("OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED"),
                kafka_config=V1KafkaConfig(
                    brokers="brokers_example",
                    topic="topic_example",
                    max_message_bytes="max_message_bytes_example",
                    compression_type="compression_type_example",
                    connection_timeout=1,
                    poll_interval=1,
                    offset_reset=V1KafkaInitialOffset("KAFKA_INITIAL_OFFSET_UNSPECIFIED"),
                ),
            ),
            sink=V1ObservationServiceDataSink(
                type=V1ObservationServiceDataSinkType("OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED"),
                kafka_config=V1KafkaConfig(
                    brokers="brokers_example",
                    topic="topic_example",
                    max_message_bytes="max_message_bytes_example",
                    compression_type="compression_type_example",
                    connection_timeout=1,
                    poll_interval=1,
                    offset_reset=V1KafkaInitialOffset("KAFKA_INITIAL_OFFSET_UNSPECIFIED"),
                ),
                fluentd_config=V1FluentdConfig(
                    type=V1FluentdOutputType("FLUENTD_OUTPUT_TYPE_UNSPECIFIED"),
                    host="host_example",
                    port=1,
                    tag="tag_example",
                    config=V1FluentdOutputBQConfig(
                        project="project_example",
                        dataset="dataset_example",
                        table="table_example",
                    ),
                ),
            ),
        ),
    ) # DatasetServiceCreateObservationServiceRequest | 

    # example passing only required values which don't have defaults set
    try:
        # CreateObservationService creates a new observation service deployment as specified by the details given in the request body.
        api_response = api_instance.dataset_service_create_observation_service(project_id, body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_create_observation_service: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |
 **body** | [**DatasetServiceCreateObservationServiceRequest**](DatasetServiceCreateObservationServiceRequest.md)|  |

### Return type

[**V1CreateObservationServiceResponse**](V1CreateObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_get_log**
> V1GetLogResponse dataset_service_get_log(project_id, id)

GetLog return details of a log.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_get_log_response import V1GetLogResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.
    id = "id_example" # str | The ID of the log resource to retrieve.

    # example passing only required values which don't have defaults set
    try:
        # GetLog return details of a log.
        api_response = api_instance.dataset_service_get_log(project_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_get_log: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |
 **id** | **str**| The ID of the log resource to retrieve. |

### Return type

[**V1GetLogResponse**](V1GetLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_get_log_writer**
> V1GetLogWriterResponse dataset_service_get_log_writer(project_id, id)

GetLogWriter return details of the log writer deployment.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_get_log_writer_response import V1GetLogWriterResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to write log resource from.
    id = "id_example" # str | The ID of the Log Writer resource to retrieve.

    # example passing only required values which don't have defaults set
    try:
        # GetLogWriter return details of the log writer deployment.
        api_response = api_instance.dataset_service_get_log_writer(project_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_get_log_writer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to write log resource from. |
 **id** | **str**| The ID of the Log Writer resource to retrieve. |

### Return type

[**V1GetLogWriterResponse**](V1GetLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_get_observation_service**
> V1GetObservationServiceResponse dataset_service_get_observation_service(project_id, id)

GetObservationService return details of the observation service deployment.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_get_observation_service_response import V1GetObservationServiceResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.
    id = "id_example" # str | The ID of the Observation Service resource to retrieve.

    # example passing only required values which don't have defaults set
    try:
        # GetObservationService return details of the observation service deployment.
        api_response = api_instance.dataset_service_get_observation_service(project_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_get_observation_service: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |
 **id** | **str**| The ID of the Observation Service resource to retrieve. |

### Return type

[**V1GetObservationServiceResponse**](V1GetObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_list_log_writers**
> V1ListLogWritersResponse dataset_service_list_log_writers(project_id)

ListLogWriters return paginated list of log writers under a project and filtered by query string.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_list_log_writers_response import V1ListLogWritersResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.

    # example passing only required values which don't have defaults set
    try:
        # ListLogWriters return paginated list of log writers under a project and filtered by query string.
        api_response = api_instance.dataset_service_list_log_writers(project_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_list_log_writers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |

### Return type

[**V1ListLogWritersResponse**](V1ListLogWritersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_list_logs**
> V1ListLogsResponse dataset_service_list_logs(project_id)

ListLogs return paginated list of logs under a project and filtered by query string.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.v1_list_logs_response import V1ListLogsResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.

    # example passing only required values which don't have defaults set
    try:
        # ListLogs return paginated list of logs under a project and filtered by query string.
        api_response = api_instance.dataset_service_list_logs(project_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_list_logs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |

### Return type

[**V1ListLogsResponse**](V1ListLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_list_observation_services**
> V1ListObservationServicesResponse dataset_service_list_observation_services(project_id)

ListObservationServices return paginated list of observation services under a project and filtered by query string.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.v1_list_observation_services_response import V1ListObservationServicesResponse
from openapi_client.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.

    # example passing only required values which don't have defaults set
    try:
        # ListObservationServices return paginated list of observation services under a project and filtered by query string.
        api_response = api_instance.dataset_service_list_observation_services(project_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_list_observation_services: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |

### Return type

[**V1ListObservationServicesResponse**](V1ListObservationServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_update_log_writer**
> V1UpdateLogWriterResponse dataset_service_update_log_writer(project_id, id, body)

UpdateLogWriter updates an existing log writer deployment as specified by the details given in the request body.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.v1_update_log_writer_response import V1UpdateLogWriterResponse
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.dataset_service_update_log_writer_request import DatasetServiceUpdateLogWriterRequest
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to write log resource from.
    id = "id_example" # str | The ID of Log Writer to update.
    body = DatasetServiceUpdateLogWriterRequest(
        log_writer=V1LogWriter(
            type=V1LogWriterType("LOG_WRITER_TYPE_UNSPECIFIED"),
            fluentd_config=V1FluentdConfig(
                type=V1FluentdOutputType("FLUENTD_OUTPUT_TYPE_UNSPECIFIED"),
                host="host_example",
                port=1,
                tag="tag_example",
                config=V1FluentdOutputBQConfig(
                    project="project_example",
                    dataset="dataset_example",
                    table="table_example",
                ),
            ),
        ),
    ) # DatasetServiceUpdateLogWriterRequest | 

    # example passing only required values which don't have defaults set
    try:
        # UpdateLogWriter updates an existing log writer deployment as specified by the details given in the request body.
        api_response = api_instance.dataset_service_update_log_writer(project_id, id, body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_update_log_writer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to write log resource from. |
 **id** | **str**| The ID of Log Writer to update. |
 **body** | [**DatasetServiceUpdateLogWriterRequest**](DatasetServiceUpdateLogWriterRequest.md)|  |

### Return type

[**V1UpdateLogWriterResponse**](V1UpdateLogWriterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dataset_service_update_observation_service**
> V1UpdateObservationServiceResponse dataset_service_update_observation_service(project_id, id, body)

UpdateObservationService updates an existing observation service deployment as specified by the details given in the request body.

### Example


```python
import time
import openapi_client
from openapi_client.api import dataset_service_api
from openapi_client.model.v1_update_observation_service_response import V1UpdateObservationServiceResponse
from openapi_client.model.rpc_status import RpcStatus
from openapi_client.model.dataset_service_update_observation_service_request import DatasetServiceUpdateObservationServiceRequest
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = dataset_service_api.DatasetServiceApi(api_client)
    project_id = "projectId_example" # str | The CaraML project ID to retrieve log resource from.
    id = "id_example" # str | The ID of Observation Service to update.
    body = DatasetServiceUpdateObservationServiceRequest(
        observation_service=V1ObservationServiceConfig(
            id="id_example",
            source=V1ObservationServiceDataSource(
                type=V1ObservationServiceDataSourceType("OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED"),
                kafka_config=V1KafkaConfig(
                    brokers="brokers_example",
                    topic="topic_example",
                    max_message_bytes="max_message_bytes_example",
                    compression_type="compression_type_example",
                    connection_timeout=1,
                    poll_interval=1,
                    offset_reset=V1KafkaInitialOffset("KAFKA_INITIAL_OFFSET_UNSPECIFIED"),
                ),
            ),
            sink=V1ObservationServiceDataSink(
                type=V1ObservationServiceDataSinkType("OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED"),
                kafka_config=V1KafkaConfig(
                    brokers="brokers_example",
                    topic="topic_example",
                    max_message_bytes="max_message_bytes_example",
                    compression_type="compression_type_example",
                    connection_timeout=1,
                    poll_interval=1,
                    offset_reset=V1KafkaInitialOffset("KAFKA_INITIAL_OFFSET_UNSPECIFIED"),
                ),
                fluentd_config=V1FluentdConfig(
                    type=V1FluentdOutputType("FLUENTD_OUTPUT_TYPE_UNSPECIFIED"),
                    host="host_example",
                    port=1,
                    tag="tag_example",
                    config=V1FluentdOutputBQConfig(
                        project="project_example",
                        dataset="dataset_example",
                        table="table_example",
                    ),
                ),
            ),
        ),
    ) # DatasetServiceUpdateObservationServiceRequest | 

    # example passing only required values which don't have defaults set
    try:
        # UpdateObservationService updates an existing observation service deployment as specified by the details given in the request body.
        api_response = api_instance.dataset_service_update_observation_service(project_id, id, body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DatasetServiceApi->dataset_service_update_observation_service: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| The CaraML project ID to retrieve log resource from. |
 **id** | **str**| The ID of Observation Service to update. |
 **body** | [**DatasetServiceUpdateObservationServiceRequest**](DatasetServiceUpdateObservationServiceRequest.md)|  |

### Return type

[**V1UpdateObservationServiceResponse**](V1UpdateObservationServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

