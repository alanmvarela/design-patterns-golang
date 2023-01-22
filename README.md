[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
![Coverage](https://img.shields.io/badge/Coverage-77.8%25-brightgreen)
[![Last Commit](https://img.shields.io/github/last-commit/alanmvarela/golang-design-patterns?style=flat-square)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/alanmvarela/golang-design-patterns)](https://goreportcard.com/report/github.com/alanmvarela/golang-design-patterns)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alanmvarela/golang-design-patterns?style=flat-square)]()

# Golang Design Patterns

A recopilation of design patterns implemented in Golang. This repository was created with the purpose of understanding the approach taken by Golang language to implement different Design Patterns.

All of the examples included in this project are based on the awsome explanations provided by [Refactor Guru](https://refactoring.guru/design-patterns/).

## Tech Stack

**Server:** Golang

## Features

- Examples and Implementation of [Adapter Design Pattern](https://github.com/alanmvarela/golang-design-patterns/tree/master/pkg/adapter).
- Examples and Implementation of [Bridge Design Pattern](https://github.com/alanmvarela/golang-design-patterns/tree/master/pkg/bridge).
- Examples and Implementation of [Decorator Design Pattern](https://github.com/alanmvarela/golang-design-patterns/tree/master/pkg/decorator).
- Examples and Implementation of [Facade Design Pattern](https://github.com/alanmvarela/golang-design-patterns/tree/master/pkg/facade).
- Tests with 100% coverage using golang tests package.

## Run Locally

Clone the project

```bash
  git clone https://github.com/alanmvarela/golang-design-patterns.git
```

Go to the project directory

```bash
  cd golang-design-patterns
```

Run the examples

```bash
  go run cmd/adapter/adapter.go
  go run cmd/bridge/bridge.go
  go run cmd/decorator/decorator.go
  go run cmd/facade/facade.go
```

## Running Tests

To run tests, run the following command

```bash
  go test ./... -cover
```

## Project Structure

```bash
.
├── cmd
│   ├── adapter
│   │   └── adapter.go
│   └── bridge
│       └── bridge.go
├── go.mod
├── pkg
│   ├── adapter
│   │   ├── adapters
│   │   │   ├── alpha_client_adapter.go
│   │   │   ├── betha_client_adapter.go
│   │   │   └── client.go
│   │   ├── adapter_test.go
│   │   ├── README.md
│   │   └── third_party
│   │       ├── alpha_client.go
│   │       └── betha_client.go
│   ├── bridge
│   │   ├── clients
│   │   │   ├── alpha_client.go
│   │   │   ├── betha_client.go
│   │   │   ├── client.go
│   │   │   └── client_test.go
│   │   ├── credentials_manager
│   │   │   ├── credential_manager_one.go
│   │   │   ├── credential_manager_two.go
│   │   │   ├── credentials_manager.go
│   │   │   └── credentials_manager_test.go
│   │   ├── logger
│   │   │   ├── logger.go
│   │   │   ├── logger_test.go
│   │   │   ├── log_one.go
│   │   │   └── log_two.go
│   │   └── README.md
│   ├── decorator
│   │   └── README.md
│   └── facade
│       └── README.md
└── README.md
```

## Lessons Learned

I learnt to identify when to implement each design pattern by applying practical examples.

Also I was able to provide implementations of each pattern while using golang best practices and most accepted approaches related to: error handling, code documentation, project structure and testing.

## Roadmap

- Add e.g for one or more implementations of proxy design pattern.
- Add execution of functions with routines and channels for error handling.
- Fix coverage test badge workflow.

## Authors

- [@alanmvarela](https://www.github.com/alanmvarela)
