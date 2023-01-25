# Adapter Design Pattern

The Adapter design pattern helps us to find a way to adapt a service interface to other third parties interfaces that we do not have power upon.
You can read more about it in [Refactor Guru - Adapter](https://refactoring.guru/design-patterns/adapter/).

## Example Escenario

In this example we considered that we have a Client interface that is used to send queries to several services.

These queries are sent using several third parties clients with different interfaces.

In order to use these third parties clients we need to implement adapters to match their method to send queries with our Client interface.

## Packet Structure

```bash
.
├── adapters
│   ├── alpha_client_adapter.go
│   ├── beta_client_adapter.go
│   └── client.go
├── adapter_test.go
├── README.md
└── third_party
    ├── alpha_client.go
    └── beta_client.go
```
