# Development Guide: Contributing to Observation Service

Thank you for your interest in contributing to Observation Service. This document provides some suggestions and guidelines on how you can get involved.

## Become a contributor

You can contribute to Observation Service in several ways:

- Contribute to feature development for the Observation Service codebase
- Report bugs
- Create articles and documentation for users and contributors
- Help others answer questions about Observation Service

### Report bugs

Report a bug by creating an issue. Provide as much information as possible
on how to reproduce the bug.

Before submitting the bug report, please make sure there are no existing issues
with a similar bug report. You can search the existing issues for similar issues.

### Suggest features

If you have an idea to improve Observation Service, submit a feature request. It will be good
to describe the use cases and how it will benefit Observation Service users in your feature
request.

## Making a pull request

You can submit pull requests to fix bugs, add new features or improve our documentation.

Here are some considerations you should keep in mind when making changes:

- While making changes
  - Make your changes in a [forked repo](#forking-the-repo) (instead of making a branch on the main Observation Service repo)
  - [Rebase from master](#incorporating-upstream-changes-from-master) instead of using `git pull` on your PR branch
  - Install [pre-commit hooks](#pre-commit-hooks) to ensure all the default linters / formatters are run when you push.
- When making the PR
  - Make a pull request from the forked repo you made
  - Ensure you leave a release note for any user facing changes in the PR. There is a field automatically generated in the PR request. You can write `NONE` in that field if there are no user facing changes.
  - Please run tests locally before submitting a PR:
    - For Go, the [unit tests](#go-tests).

### Forking the repo

Fork the Observation Service Github repo and clone your fork locally. Then make changes to a local branch to the fork.

See [Creating a pull request from a fork](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork)

### Pre-commit Hooks

Setup [`pre-commit`](https://pre-commit.com/) to automatically lint and format the codebase on commit:

1. Ensure that you have Python (3.7 and above) with `pip`, installed.
2. Install `pre-commit` with `pip` &amp; install pre-push hooks

    ```sh
    # Clear existing hooks    
    git config --unset-all core.hooksPath
    rm -rf .git/hooks
    # Install hooks
    make setup
    ```

3. On push, the pre-commit hook will run. This runs `make format` and `make lint`.

## Observation Service using Go

Observation service is written using Go, and the following describes how to setup your development environment.

### Environment Setup

- Install Golang, [`protoc` with the Golang &amp; grpc plugins](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

### Code Style & Linting

We are using [golangci-lint](https://github.com/golangci/golangci-lint), and we can run the following commands for formatting.

```sh
# Formatting for linting issues
make format

# Checking for linting issues
make lint
```

### Go tests

For **Unit** tests, we follow the convention of keeping it beside the main source file.
