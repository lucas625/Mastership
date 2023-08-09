# Running

## Table of Contents

- [Running](#running)
  - [Table of Contents](#table-of-contents)
  - [Running the project](#running-the-project)
    - [Run for Development](#run-for-development)
      - [Installing](#installing)
        - [Setup Analyzer](#setup-analyzer)
        - [Setup Calculator](#setup-calculator)
        - [Setup Experimenter](#setup-experimenter)
        - [Setup Frontend](#setup-frontend)
      - [Building for Development](#building-for-development)
        - [Build Analyzer](#build-analyzer)
        - [Security Checker](#security-checker)
        - [Calculator](#calculator)
        - [Experimenter](#experimenter)
        - [Frontend](#frontend)
    - [Run with Minikube](#run-with-minikube)
    - [Run for Production](#run-for-production)

## Running the project

First of all, it is required to have the environment vars setup, so build them:

```bash
# Go to the scripts folder
cd scripts

# Builds the env vars
./build_env_vars.sh

# Go back to root folder
cd ..
```

### Run for Development

#### Installing

##### Setup Analyzer

```bash
# Go to analyzer folder
cd analyzer

# Setup environment
python3 -m venv venv
source venv/bin/activate

# Install requirements
pip install -r requirements.txt

# Leave environment
deactivate

# Return to root folder 
cd ..
```

##### Setup Calculator

```bash
# Go to calculator folder
cd calculator

# Install requirements
go mod download
go get ./...

# Go back to root folder
cd ..
```

##### Setup Experimenter

```bash
# Go to experimenter folder
cd experimenter

# Install requirements
go mod download
go get ./...

# Go back to root folder
cd ..
```

##### Setup Frontend

```bash
# Go to frontend folder
cd frontend

# Install requirements
yarn install

# Go back to root folder
cd ..
```

#### Building for Development

You will need 4 terminal tabs to do it. On each of these tabs you will run of the services in no particular order.
The project will be available on [Port 8080](http://127.0.0.1:/8080), but each service has its own port that you may check to try and do some direct requests.

##### Build Analyzer

```bash
# Go to analyzer folder
cd analyzer

# Start environment
source venv/bin/activate

# Run service
python manage.py runserver
```

##### Security Checker

```bash
# Go to security_checker folder
cd security_checker

# Start environment
source venv/bin/activate

# Run service
python manage.py runserver 8003
```

##### Calculator

```bash
# Go to calculator folder
cd calculator

# Run service
go run cmd/main.go
```

##### Experimenter

```bash
# Go to experimenter folder
cd experimenter

# Run service
go run cmd/main.go
```

##### Frontend

```bash
# Go to frontend folder
cd frontend

# Run service
yarn serve --port 8080
```

### Run with Minikube

```bash
# Go to the scripts folder
cd scripts

# Generate the env_vars.sh in case you didn't yet
./build_env_vars.sh
# Make sure you check the variables generated on the env_vars.sh file.

# Run with minikube
./run_minikube.sh

# Stop minikube when you no longer need it
./stop_minikube.sh
```

### Run for Production

Work in progress.
