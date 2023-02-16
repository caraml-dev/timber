import json
import typing

import requests


class DatasetServiceClientError(Exception):
    pass


class NotFound(DatasetServiceClientError):
    pass


class ServerError(DatasetServiceClientError):
    pass


class DatasetServiceClient:
    def __init__(
        self,
        dataset_service_url="http://localhost:8080/v1",
    ):
        self._dataset_service_url = dataset_service_url.rstrip("/")

    def create_observation_service(
        self,
        project_id: int,
        create_observation_service_request: typing.Dict[str, typing.Any],
    ):
        return requests.post(
            f"{self._dataset_service_url}/projects/{project_id}/observation_services",
            data=json.dumps(create_observation_service_request),
            headers={"Content-Type": "application/json"},
        )

    def update_observation_service(
        self,
        project_id: int,
        id: int,
        update_observation_service_request: typing.Dict[str, typing.Any],
    ):
        return requests.put(
            f"{self._dataset_service_url}/projects/{project_id}/observation_services/{id}",
            data=json.dumps(update_observation_service_request),
            headers={"Content-Type": "application/json"},
        )

    def create_log_writer(
        self,
        project_id: int,
        create_log_writer_request: typing.Dict[str, typing.Any],
    ):
        return requests.post(
            f"{self._dataset_service_url}/projects/{project_id}/log_writers",
            data=json.dumps(create_log_writer_request),
            headers={"Content-Type": "application/json"},
        )
