# Facade Design Pattern

A facade provides a simplified interface for a libray, framework or a set of classes. It allows to clean the code of the main logic and let the facade class to handle the instantiation and orchestration of the set of classes.

You can read more about it in [Refactor Guru - Facade](https://refactoring.guru/design-patterns/facade/).

## Example Escenario

We have a service that has to instantiate several clients to send queries to different APIs. After this all the gathered information should be put together into a list and returned by the service.

Instead of having all the logic in the services we will create a facade class to handle the instantiation of the clients, perform the queries and gather the data.

This will allow the Service to simple call a function of the facade to get the data and continue its processing.

## Packet Structure

```bash
.
├── clients
│   ├── alpha_client.go
│   ├── beta_client.go
│   ├── client.go
│   └── client_test.go
├── facades
│   ├── facades_test.go
│   └── serviceFacade.go
├── README.md
└── service
    ├── service.go
    └── service_test.go
```
