import pytest

from e2e.dataset_service_client import DatasetServiceClient


def pytest_addoption(parser):
    parser.addoption(
        "--dataset-service-url", action="store", help="dataset service url", default=""
    )
    parser.addoption(
        "--dataset-service-bin",
        action="store",
        help="path to dataset service binary, if url is not provided",
        default="",
    )
    parser.addoption("--env", action="store", help="path to env", default="local")


from e2e.fixtures.mockups.mlp_service import *  # noqa
from e2e.fixtures.services import *  # noqa


@pytest.fixture
def dataset_service_client(dataset_service):
    return DatasetServiceClient(f'{dataset_service.rstrip("/")}/v1')
