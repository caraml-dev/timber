# V1ObservationServiceDataSinkType

- OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP: No-Op represents no need to flush logs to any data sink  - OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT: Observation Service will publish logs to standard output  - OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA: Observation Service will flush logs to a Kafka sink  - OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD: Observation Service will flush logs to Fluentd

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**value** | **str** | - OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP: No-Op represents no need to flush logs to any data sink  - OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT: Observation Service will publish logs to standard output  - OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA: Observation Service will flush logs to a Kafka sink  - OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD: Observation Service will flush logs to Fluentd | defaults to "OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED",  must be one of ["OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED", "OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP", "OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT", "OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA", "OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD", ]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


