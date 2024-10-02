# Building a Dockerized gRPC-Backed Image Search Engine

## Project Overview
This project implements a gRPC-backed image search engine, allowing users to search for images based on keywords using two different programming languages: **Python** and **GoLang**. Both implementations are containerized using Docker, enabling easy deployment and management.

## Technologies Used
- **Programming Languages**: Python and GoLang
- **Python Version**: 3.10.11
- **GoLang Version**: go1.21.2
- **Frameworks**: gRPC, Flask (for Python)
- **Containerization**: Docker

## Components
The project consists of two main components: 
1. **gRPC Image Search Server and Client using Python**
2. **gRPC Image Search Server and Client using GoLang**

### 1. gRPC Image Search Server and Client using Python

#### Server Implementation
- **File**: `server.py`
- The server exposes the `ImageSearchService` with a single method `SearchImage`. It responds with a random image associated with the provided keyword.
- **Start the Server**: 
  ```bash
  python server.py
  ```
- **Port**: The server listens on port **50051**.

#### Client Implementation
- **File**: `app.py`
- A simple web application that allows users to enter a keyword and search for images via the gRPC server.
- **Required Packages**: Install the necessary Python packages:
  ```bash
  pip install grpcio grpcio-tools flask
  ```
- **Start the Client Application**: 
  ```bash
  python app.py
  ```
- **Access the Client**: Open a web browser and navigate to `http://localhost:4000`. Enter a keyword and click the "Search" button to retrieve the image.

#### Dockerization
The server and client are dockerized and accessed through a network called **mynetwork**.
- **Docker Commands**:
  ```bash
  docker-compose build
  docker ps                  # Check if containers are running
  docker network create mynetwork
  docker-compose up
  ```

#### gRPC Protocol
The service definitions are provided in the `image_search.proto` file. Compile the proto file to generate the client and server code using:
```bash
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. image_search.proto
```

---

### 2. gRPC Image Search Server and Client using GoLang

#### Requirements
- **Golang**
- **Protocol Buffers** (`protoc`)
- **gRPC Go Packages**:
  - `github.com/golang/protobuf/protoc-gen-go`
  - `google.golang.org/grpc`

#### Server Implementation
- **File**: `my_golang_project/server/server.go`
- The server implements image search functionality and listens on port **8080**.
- **Start the Server**: 
  ```bash
  go run server.go
  ```

#### Client Implementation
- **File**: `my_golang_project/client/client.go`
- Provides a web interface for users to enter keywords and view corresponding images.
- **Run the Client**: 
  ```bash
  go run client.go
  ```
- **Access the Client**: Open a web browser and go to `http://localhost:8081`.

#### Dockerization
Both server and client are defined in the `docker-compose.yaml` file and are containerized.
- **Docker Commands**:
  ```bash
  docker-compose build
  docker ps                  # Check if containers are running
  docker network create mynetwork
  docker-compose up
  ```

#### gRPC Service Definition
The gRPC service is defined in the `pb/image_search.proto` file, detailing the service methods and message types used for image searches.

---

### References
1. [Golang gRPC Example](https://github.com/DekivadiyaKishan/golang-grpc)
2. [Golang gRPC Tutorial: Building High-Performance Web Services](https://bacancytechnology.com)
3. [Protocol Buffers](https://github.com/protocolbuffers/protobuf)
4. [Golang gRPC Overview](https://www.golinuxcloud.com/golang-grpc/)
5. [Using gRPC with Docker](https://medium.com/@matzhouse/go-grpc-and-docker-c1fb2ec8e9f0)
6. [gRPC with Python](https://speedscale.com/blog/using-grpc-with-python/)
7. [Raising Server Errors to the Client with gRPC](https://stackoverflow.com/questions/40998199/raising-a-server-error-to-the-client-with-grpc)
8. [gRPC Python Docker Example](https://github.com/flavienbwk/gRPC-Python-Docker-Example)

--- 
