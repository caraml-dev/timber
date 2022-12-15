# V1ObservationServiceDataSourceType

- OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER: No-Op represents no need to fetch logs from any data source, this should be selected if Observation Service should be deployed for just the eager API  - OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA: Observation Service will poll logs from a Kafka source

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**value** | **str** | - OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER: No-Op represents no need to fetch logs from any data source, this should be selected if Observation Service should be deployed for just the eager API  - OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA: Observation Service will poll logs from a Kafka source | defaults to "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED",  must be one of ["OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED", "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER", "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA", ]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


