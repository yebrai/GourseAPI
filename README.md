# GourseAPI - Go HTTP API with Hexagonal Architecture

GourseAPI is a Go-based application implementing Hexagonal Architecture, following best design practices for separation of concerns. It leverages an Event and Command Bus to manage use cases in a decoupled and scalable manner.

## Key Features

- **Hexagonal Architecture**: Utilizes the ports and adapters pattern to maintain a clean and independent domain.
- **Event & Command Bus**: Implements internal messaging to efficiently handle events and commands.
- **MySQL Persistence**: Decoupled repository layer using MySQL as the data store.
- **Gin Framework**: High-performance HTTP request handling with Gin.
- **Docker-ready**: Seamless setup and deployment with Docker integration.
- **Middleware & Validations**: Custom middleware and domain validations to ensure data integrity.
- **Testing**: Includes unit and integration tests to ensure code quality.

## Requirements

- Go v1.15+
- MySQL (see details below).

## Project Structure

The project is structured as a single Go module containing multiple applications. Each folder represents an independently executable application. Some of the core functionalities include:

1. **HTTP Endpoints**: RESTful routes implemented with Gin.
2. **Dependency Injection**: Decoupled repositories for business logic separation.
3. **Domain Validations**: Enforces business rules before data persistence.
4. **Domain Event Publishing**: Enables system extensibility via event-driven architecture.
5. **Middleware & Error Handling**: Implements middleware for logging, authentication, and error recovery.
6. **Docker Deployment**: Optimized containerized deployment.

## Starting the Application

1- Clone the repository:

```sh
git clone https://github.com/yebrai/GourseAPI.git
```

2- Give up docker:

```sh
docker-compose up --build
```

## Testing

To run all tests:

```sh
go test ./...
```

You can debug from logs:

```sh
docker-compose logs -f mooc-api
```

## Example API Requests

### GET Request Example

```sh
curl -X GET http://localhost:8080/courses
```

### POST Request Example

```sh
curl -X POST http://localhost:8080/courses   -H "Content-Type: application/json"   -d '{
    "ID": "8a1c5cdc-ba57-445a-994d-aa412d23724f",
    "Name": "Demo Course",
    "Duration": "10 months"
  }'
```

### GET Request Example

```sh
curl -X GET http://localhost:8080/health
```

## Acknowledgments

This project was developed following the teachings of the [CodelyTV](https://pro.codely.com/home/learn/) course and the valuable insights shared by [Friends of Go](https://friendsofgo.tech/). We are grateful to both communities for their contributions to the Go ecosystem and for fostering knowledge-sharing within the developer community.
