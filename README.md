# Mastership

This is the Mastership project for Lucas Aurelio.
The application consists on an interface and a few internal services, used to test the response time of consecutive requests so performance can be analyzed.

## Table of Contents

- [Mastership](#mastership)
  - [Table of Contents](#table-of-contents)
  - [Personal](#personal)
  - [Requirements](#requirements)
  - [Project Structure](#project-structure)
  - [How To Run](#how-to-run)

## Personal

- Lucas Aurelio Gomes Costa - **laurelio.costa@gmail.com** (Owner)
- Carlos André Girmarães Ferraz - **cagf@cin.ufpe.br** (Advisor)

## Requirements

- Ubuntu
- Python 3.9
- Go 1.9
- Yarn
- Docker
- Minikube (for runnig with minikube)
- Kubernetes

## Project Structure

- **mastership**
  - **analyzer**: A django rest service that performs some analysis on the results of the experiments.
  - **calculator**: A golang service with common calculator operations.
  - documentation: Additional documentation about the project.
  - **experimenter**: A golang service that performs multiple calls to the calculator so as to get the RTTs.
  - **frontend**: A vue service that serves the UI.
  - kubernetes: The kubernetes configuration files.
  - **reverse_proxy**: A nginx service for being the reverse proxy of the application when using docker/kubernetes.
  - scripts: Scripts for the application, such as generating env vars ou building the project.
    - *build_env_vars.sh*: Builds the environment variables. Remember to update the **env_vars.sh** file that was generated with the necessary information.
    - *build_images.sh*: Builds the docker image.
    - *create_kubernetes_elements.sh*: Used for creating the kubernetes elements in any environment (minikube/cloud).
    - *remove_kubernetes_elements.sh*: Used for removing the kubernetes elements in any environment (minikube/cloud).
    - *run_minikube.sh*: Used for running minikube with the project. It is required to have the **env_vars.sh** file with the values setup.
    - *stop_minikube.sh*: Used for stopping minikube and cleaning the project. It is required to have the **env_vars.sh** file with the values setup.

## How To Run

Please check the [general documentations](documentation) folder to have more details on how to install dependencies of the project and [security cases](scripts/security_cases/security_cases.md) to options to run.
