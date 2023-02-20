import time

import pytest
import yaml
from kubernetes import client, config

from e2e.dataset_service_client import DatasetServiceClient

TEST_PROJECT_ID = 1
TEST_PROJECT_NAME = "test-project"
TIMEOUT_SECONDS = 300


@pytest.fixture()
def k8s_client() -> client.AppsV1Api:
    config_file = "/tmp/kubeconfig-timber-dev.yaml"
    with open(config_file, "r") as stream:
        try:
            yaml.safe_load(stream)
        except yaml.YAMLError:
            raise
    config.load_kube_config(config_file)
    api_instance = client.AppsV1Api()
    return api_instance


@pytest.mark.order(1)
def test_simple_observation_service_creation(
    dataset_service_client: DatasetServiceClient,
    k8s_client: client.AppsV1Api,
):
    # Create Observation Service
    service_name = "my-observation"
    req_body = {
        "observation_service": {
            "project_id": 1,
            "name": service_name,
            "source": {
                "type": "OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA",
                "kafka": {"brokers": "kafka.mlp.svc.cluster.local", "topic": "sample"},
            },
        }
    }
    resp = dataset_service_client.create_observation_service(TEST_PROJECT_ID, req_body)

    assert resp.status_code == 200
    assert resp.content
    body = resp.json()

    # TODO: Improve the assertion once DB is implemented
    assert body["observation_service"]
    assert body["observation_service"]["status"] == "STATUS_DEPLOYED"

    wait_statefulset_ready(k8s_client, TEST_PROJECT_NAME, "os-my-observation-fluentd")
    wait_deployment_ready(
        k8s_client, TEST_PROJECT_NAME, "os-my-observation-observation-svc"
    )


@pytest.mark.order(2)
def test_simple_observation_service_updation(
    dataset_service_client: DatasetServiceClient,
    k8s_client: client.CoreV1Api,
):
    # Upgrade Observation Service
    service_name = "my-observation"
    req_body = {
        "observation_service": {
            "project_id": 1,
            "name": service_name,
            "source": {
                "type": "OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA",
                "kafka": {"brokers": "kafka.mlp.svc.cluster.local", "topic": "hello"},
            },
        }
    }
    resp = dataset_service_client.update_observation_service(
        1, TEST_PROJECT_ID, req_body
    )

    assert resp.status_code == 200
    assert resp.content
    body = resp.json()

    # TODO: Improve the assertion once DB is implemented
    assert body["observation_service"]
    assert body["observation_service"]["status"] == "STATUS_DEPLOYED"

    wait_statefulset_ready(k8s_client, TEST_PROJECT_NAME, "os-my-observation-fluentd")
    wait_deployment_ready(
        k8s_client, TEST_PROJECT_NAME, "os-my-observation-observation-svc"
    )


def test_simple_router_log_writer_creation(
    dataset_service_client: DatasetServiceClient,
    k8s_client: client.AppsV1Api,
):
    # Create router log wrtier
    req_body = {
        "log_writer": {
            "project_id": 1,
            "name": "my-router-log",
            "source": {
                "type": "LOG_WRITER_SOURCE_TYPE_ROUTER_LOG",
                "router_log_source": {
                    "router_id": 1,
                    "router_name": "my-router",
                    "kafka": {
                        "brokers": "kafka.mlp.svc.cluster.local",
                        "topic": "my-router-log",
                    },
                },
            },
        }
    }
    resp = dataset_service_client.create_log_writer(TEST_PROJECT_ID, req_body)

    assert resp.status_code == 200
    assert resp.content
    body = resp.json()

    # TODO: Improve the assertion once DB is implemented
    assert body["log_writer"]
    assert body["log_writer"]["status"] == "STATUS_DEPLOYED"

    wait_statefulset_ready(k8s_client, TEST_PROJECT_NAME, "rl-my-router-log-fluentd")


def test_simple_prediction_log_writer_creation(
    dataset_service_client: DatasetServiceClient,
    k8s_client: client.AppsV1Api,
):
    # Create prediction log writer
    req_body = {
        "log_writer": {
            "project_id": 1,
            "name": "my-model-log",
            "source": {
                "type": "LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG",
                "prediction_log_source": {
                    "model_id": 1,
                    "model_name": "my-model",
                    "kafka": {
                        "brokers": "kafka.mlp.svc.cluster.local",
                        "topic": "my-prediction-log",
                    },
                },
            },
        }
    }
    resp = dataset_service_client.create_log_writer(TEST_PROJECT_ID, req_body)

    assert resp.status_code == 200
    assert resp.content
    body = resp.json()

    # TODO: Improve the assertion once DB is implemented
    assert body["log_writer"]
    assert body["log_writer"]["status"] == "STATUS_DEPLOYED"

    wait_statefulset_ready(k8s_client, TEST_PROJECT_NAME, "pl-my-model-log-fluentd")


def wait_deployment_ready(
    k8s_client: client.AppsV1Api,
    namespace: str,
    deployment_name: str,
    timeout: int = 300,
):
    start = time.time()
    while time.time() - start < timeout:
        time.sleep(2)
        response = k8s_client.read_namespaced_deployment(deployment_name, namespace)
        s = response.status
        if (
            s.updated_replicas == response.spec.replicas
            and s.replicas == response.spec.replicas
            and s.available_replicas == response.spec.replicas
            and s.observed_generation >= response.metadata.generation
        ):
            return True
        else:
            print(f"Waiting for deployment {deployment_name} to be ready")

    raise RuntimeError(f"Waiting timeout for deployment {deployment_name}")


def wait_statefulset_ready(
    k8s_client: client.AppsV1Api,
    namespace: str,
    statefulset_name: str,
    timeout: int = 300,
):
    start = time.time()
    while time.time() - start < timeout:
        time.sleep(2)
        response = k8s_client.read_namespaced_stateful_set(statefulset_name, namespace)
        s = response.status
        if (
            s.updated_replicas == response.spec.replicas
            and s.replicas == response.spec.replicas
            and s.available_replicas == response.spec.replicas
            and s.observed_generation >= response.metadata.generation
        ):
            return True
        else:
            print(f"Waiting for statefulset {statefulset_name} to be ready")

    raise RuntimeError(f"Waiting timeout for statefuset {statefulset_name}")
