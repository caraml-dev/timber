import os
import socket
import subprocess
import time
from os import PathLike
from os.path import isfile
from pathlib import Path
from typing import List, Optional, Tuple, Union

import pytest
import yaml
from _pytest.fixtures import FixtureRequest
from dotenv import load_dotenv

DATASET_SERVICE_NAME = "dataset-service"

PROJECT_ROOT_DIR = Path(__file__).parents[3]
TEST_DIR = Path.joinpath(PROJECT_ROOT_DIR, "tests", "e2e")


def _service_dir(service_name) -> Path:
    return Path.joinpath(PROJECT_ROOT_DIR, service_name)


def _default_bin_path(service_name) -> Path:
    return Path.joinpath(
        _service_dir(service_name),
        f"cmd/{DATASET_SERVICE_NAME}/bin",
        DATASET_SERVICE_NAME,
    )


def _wait_port_open(host, port, max_wait=60):
    print(f"Waiting for port {port}")
    start = time.time()

    while True:
        try:
            socket.create_connection((host, port), timeout=1)
        except OSError:
            if time.time() - start > max_wait:
                raise

            time.sleep(1)
        else:
            return


def _start_binary(
    binary: Union[PathLike, str],
    options: List[str] = None,
    working_dir: Optional[PathLike] = None,
):
    if not isfile(binary):
        raise ValueError(f"The binary file '{binary}' doesn't exist")

    cmd = [binary]
    if options:
        cmd.extend(options)

    return subprocess.Popen(
        cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, cwd=working_dir
    )


def _service_config(env: str, service_name: str) -> PathLike:
    return Path.joinpath(TEST_DIR, "config", f"{service_name}.{env}.yaml")


def _env_config(env: str) -> PathLike:
    return Path.joinpath(TEST_DIR, "config", f"{env}.env")


def _start_dataset_service(
    bin_path: PathLike, env: str, mlp_service_url
) -> Tuple[str, subprocess.Popen]:
    load_dotenv(_env_config(env))
    os.environ["MLPCONFIG::URL"] = mlp_service_url

    config_path = _service_config(env, DATASET_SERVICE_NAME)
    process = _start_binary(
        bin_path,
        ["serve", "--config", config_path],
        working_dir=_service_dir(DATASET_SERVICE_NAME),
    )

    with open(config_path, "r") as config_file:
        config = yaml.safe_load(config_file)

    try:
        port = config.get("Port", 8080)
        _wait_port_open("localhost", port, 15)
    except OSError:
        outs, errs = process.communicate(timeout=5)
        print(outs)
        print(errs)
        raise ValueError("unable to run dataset service binary")

    dataset_service_url = f"http://localhost:{port}"

    return dataset_service_url, process


@pytest.fixture(scope="session")
def dataset_service(pytestconfig, request: FixtureRequest):
    dataset_service_url = pytestconfig.getoption("dataset_service_url")
    process = None
    if dataset_service_url == "":
        bin_path = pytestconfig.getoption("dataset_service_bin")
        if bin_path == "":
            bin_path = _default_bin_path(DATASET_SERVICE_NAME)
        env = pytestconfig.getoption("env")
        dataset_service_url, process = _start_dataset_service(
            bin_path, env, mlp_service_url=request.getfixturevalue("mlp_service")
        )
    yield dataset_service_url
    if process:
        process.terminate()
