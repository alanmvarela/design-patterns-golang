# Decorator Design Pattern

Decorator pattern allows to add new behaviors to a classes implementing an interface by generating a wraper for it. The wraper will implement the same interface as the class and add some additional logic before or after calling the same methods of the wraped classes. We can define which class implementing the interface are we using by changing it when instatiating the wraper(decorator).
You can read more about it in [Refactor Guru - Decorator](https://refactoring.guru/design-patterns/decorator/).

## Example Escenario

We have a Service that needs to send queries to different APIs, depending on the scenario, get the data and perform some modifications on it before returning to the client. We need to be able to easily add new APIs calls to this service without messing around too much with the code.
Enter decorator pattern: we can create a Client interface to send the Queries to specific APIs and a Service decorator that can be instantiated using this clients. Therefore, if we want to make the Service able to send a query to a different API we can instantiate it using a new function for it.

## Packet Structure

```bash
.
├── clients
│   ├── alpha_client.go
│   ├── beta_client.go
│   ├── client.go
│   └── client_test.go
├── README.md
└── service
    ├── service.go
    └── service_test.go
```
