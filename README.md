# Golang Practice

This repository contains simple Go modules written to implement Go concepts and learn Go programming through practical example projects. Each module is designed to focus on specific aspects of Go, making it easier to explore the language fundamentals incrementally.

## Project Structure

The project is organized into several subdirectories, each representing a distinct Go module or example project. Below is the structure of the repository:

```
├── src
│   ├── hello-world
│   │   └── main.go
│   └── print-directory
│       ├── README.md
│       ├── main.go
│       └── main_test.go
```

## Running Modules

Navigate to the desired module and run the Go program. For example, to run the hello-world module:
```sh
cd src/hello-world
go run main.go
```

To run the print-directory module:
```sh
cd src/print-directory
go run main.go /path/to/directory
```

## Running Tests
Some modules include test cases to verify functionality. To run the tests, navigate to the module directory and execute:
```sh
go test -v
```
