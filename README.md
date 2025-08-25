# Chain HealthCheck

This repository provides a health check service for blockchain nodes, designed to report the status of chain nodes.

## Features
- Health check endpoint for node status
- Modular service structure
- Docker support for easy deployment
- Configurable via environment variables

## Project Structure
```
docker-compose.yml
Dockerfile
go.mod
go.sum
app/
  app.go
config/
  config.go
handlers/
  health.go
models/
  status.go
response/
  response.go
routes/
  routes.go
server/
  main.go
services/
  status.go
```

## Getting Started

### Prerequisites
- Go 1.20+
- Docker (optional, for containerized deployment)

### Running Locally
```bash
git clone https://github.com/Qubetics/chain-healthCheck.git
cd chain-healthCheck
go run server/main.go
```

### Using Makefile
You can use the provided Makefile for easier commands:

To run the app locally:
```bash
make run
```

To build and run with Docker:
```bash
make docker-up
```

## Build

To build the application binary, run:

```sh
make build
```

This will create a binary named `chain-healthcheck` in the project root.

### Using Docker
```bash
 docker-compose up --build -d
```

## API Endpoints
- `/health` : Returns the health status of the node
