# DockerTemplate_Go

## Project Overview
This project is a simple Go web application running within a Docker container. It responds with "Hello, world!" on accessing the `/incoming` endpoint.

## Requirements
- Docker
- Docker Compose

## Installation and Running

### Building and Running with Docker Compose
To build and start the server using Docker Compose, run the following commands:

```bash
make new
```

This will build the Docker container and start the service.

### Stopping the Service
To stop and remove the container, you can use:

```bash
make clean
```

### Accessing the Application
Once the application is running, you can access it at http://localhost:9505/incoming.

Development and Debugging
To debug the application, you can use the provided debug command which will show logs:

```bash
Copy code
make debug
```

### Testing
Once up and running, you can test this with the following command to get the "Hello World"
return from the app.

```bash
curl http://localhost:9505/incoming
```
