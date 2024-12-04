# gRPC Multiplayer Mode Cluster  

A gRPC-based multiplayer service implemented in Go, using MongoDB and Redis as the primary database and cache, respectively. The project is containerized with Docker Compose to enable seamless deployment in a clustered environment.  


## Features  
- **gRPC Communication**: Enables efficient, scalable, and low-latency communication between services.  
- **MongoDB Integration**: Secure user and game data storage with environment-configured collections.  
- **Redis Cache**: Provides rapid data access for game states and user session management.  
- **Dockerized Deployment**: Fully containerized using Docker Compose with a bridge network for inter-service communication.  
- **Environment Configuration**: Fully configurable via `.env` variables for flexibility.

## Prerequisites
Ensure you have the following installed:

- Docker: To run the containers.
- Docker Compose: For service orchestration.
- Redis: To set up local caching.
- MongoDB: To set up the local database.
- Golang 1.20 or higher: To run and develop the application.

## Docker Hub
The prebuilt Docker image for this project is available on Docker Hub: https://hub.docker.com/r/rohanyh/multiplayer-mode-usage-web-service

## Installation

### Run Locally,
1. Clone the repository:
```bash
  git clone https://github.com/rohanyh101/multiplayer-mode-usage-web-service.git multiplayer-mode-usage-web-service
  cd multiplayer-mode-usage-web-service
```

2. Set up environment variables:
- Copy the `.env.example` file to  `.env` and update values as required.

```bash
  cp .env.example .env
```

3. Insert dummy data:
 - Insert the dummy data into MongoDB container service via running `script/main.go` go script

```bash
  make script
```

4. Generate gRPC code:
- Ensure the `protoc` compiler is installed. Then, run:

```bash
  make gen
```

4. Run the service locally (cmd/main.go):

```bash
  make run
```

### Run Via Docker(DockerHub Image)

1. First Pull the Docker Image:
- Pull the Docker Image from DockerHub Repository, [here](https://hub.docker.com/r/rohanyh/multiplayer-mode-usage-web-service)

```bash
  docker pull rohanyh/multiplayer-mode-usage-web-service
```

2. Start the cluster:

```bash
  docker run rohanyh/multiplayer-mode-usage-web-service
```

3. Insert dummy data:
 - Insert the dummy data into MongoDB container service via running `script/main.go` go script

```bash
  make script
```

4. Access the services:
- MongoDB: `mongodb://mongodb:27017`
- Redis: `redis://redis:6379`
- gRPC Service: Accessible via port `50051`.

## Test Commands
- All the gRPC Test Commands are in the `grpc.txt` file of the repo just copy and paste see the responses accordingly, [here](grpc.txt)

### Docker Compose Services
- **multiplayer**: Runs the gRPC service. Built using a multi-stage Dockerfile.
- **mongo**: MongoDB service with root authentication.
- **redis**: Redis service with password protection enabled.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request for any improvements or bug fixes.
