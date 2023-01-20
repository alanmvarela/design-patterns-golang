## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Last Commit](https://img.shields.io/github/last-commit/alanmvarela/golang-design-patterns?style=flat-square)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/alanmvarela/golang-design-patterns)](https://goreportcard.com/report/github.com/alanmvarela/golang-design-patterns)

# Golang Design Patterns

A recopilation of design patterns implemented in Golang. This repository was created with the purpose of understanding the approach taken by Golang language to implement different Design Patterns.

All of the examples included in this project are based on the awsome explanations provided by [Refactor Guru](https://refactoring.guru/design-patterns/).

## Tech Stack

**Server:** Golang

## Features

- TBD

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
│   └── adapter
│       └── adapter.go
├── go.mod
├── pkg
│   ├── adapter
│   │   ├── alpha_client_adapter.go
│   │   ├── betha_client_adapter.go
│   │   ├── client_adapter_test.go
│   │   ├── clients
│   │   │   ├── alpha_client.go
│   │   │   └── betha_client.go
│   │   └── dns_client.go
│   ├── bridge
│   ├── decorator
│   └── facade
└── README.md
```

## Lessons Learned

What did you learn while building this project? What challenges did you face and how did you overcome them?

## Roadmap

- Add e.g for one or more implementations of proxy design pattern.
- Add execution of functions with routines and channels for error handling.
- Add coverage test badge.

## Authors

- [@alanmvarela](https://www.github.com/alanmvarela)
