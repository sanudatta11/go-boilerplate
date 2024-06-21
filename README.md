
# Go Boilerplate

## Overview

This project is a Go application with a structured layout and proper use of mocking and interfaces to facilitate testing and modularity. The application includes various components such as controllers, models, services, and more, making it easy to extend and maintain.

## Project Structure

```
- constants/   # Contains constant values used throughout the application
- controllers/ # Handles incoming HTTP requests and routes them to the appropriate service
- db/          # Database layer, handles database connections and queries
- mocks/       # Contains mock implementations for testing
- models/      # Defines the data models and structures
- server/      # Sets up and starts the HTTP server with Router
- services/    # Business logic and services layer
- .gitignore   # Specifies files and directories to be ignored by git
- Makefile     # Contains commands to build, test, and run the application
- go.mod       # Go module file
- go.sum       # Go module dependencies file
- main.go      # Entry point of the application
- server.go    # Server setup and configuration
- README.md    # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

   ```sh
   git clone [<repository-url>](https://github.com/sanudatta11/go-boilerplate)
   cd go-boilerplate
   ```

2. Install dependencies:

   ```sh
   go mod download
   go mod tidy
   ```

### Running the Application

1. You can run the application directly using Go:

   ```sh
   go run main.go
   ```

### Testing

The application includes a suite of tests using mock implementations to ensure reliability and correctness. To run the tests, use the following command:

```sh
make test
```

## Usage

### Configuration

<TODO>
  
### API Endpoints

The API endpoints are defined in the `controllers` package. Detailed documentation for each endpoint can be added here, including request and response formats, example requests, and any required authentication or authorization.
