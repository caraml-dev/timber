import json
import typing

import requests


class DatasetServiceClientError(Exception):
    pass


class NotFound(DatasetServiceClientError):
    pass


class ServerError(DatasetServiceClientError):
    pass


TEST_PROJECT_ID = 999
TEST_PROJECT_NAME = "test-project"
TIMEOUT_SECONDS = 300


class DatasetServiceClient:
    def __init__(
        self,
        dataset_service_url="http://localhost:8080/v1",
    ):
        self._dataset_service_url = dataset_service_url.rstrip("/")

    def create_observation_service(
        self, project_id: int, observation_service_config: typing.Dict[str, typing.Any]
    ):
        resp = requests.post(
            f"{self._dataset_service_url}/projects/{project_id}/observation_services",
            data=json.dumps(observation_service_config),
            headers={"Content-Type": "application/json"},
        )
        try:
            assert resp.status_code == 200, resp.content
        except AssertionError:
            if resp.status_code != 400:
                raise

        # Example written form - {'observationService': {'id': 'cde080d5-11ac-4027-906f-95d1f5d490f9'}}
        return resp.json()["observationService"]

    def update_observation_service(
        self,
        project_id: int,
        id: int,
        observation_service_config: typing.Dict[str, typing.Any],
    ):
        resp = requests.put(
            f"{self._dataset_service_url}/projects/{project_id}/observation_services/{id}",
            data=json.dumps(observation_service_config),
            headers={"Content-Type": "application/json"},
        )
        try:
            assert resp.status_code == 200, resp.content
        except AssertionError:
            if resp.status_code != 400:
                raise

        # Example written form - {'observationService': {'id': 'cde080d5-11ac-4027-906f-95d1f5d490f9'}}
        return resp.json()["observationService"]
