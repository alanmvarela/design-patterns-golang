# Bridge Design Pattern

The Bridge design pattern helps us to separate the implementation of one or several related classes in interfaces and implementations.
You can read more about it in [Refactor Guru - Bridge](https://refactoring.guru/design-patterns/bridge/).

## Example Escenario

In this example we considered that we have a Client interface that is used to send queries to a service. Several clients implement this interface sending queries to different APIs each. Each client Credential Manager and logger handler can be different from each other.

Having this in mind we should be able to define the CredentialManager and Logger for each Client at the moment of instantiation.

In order to do this, the bridge pattern is used to detach the implementation of the Logger and CredentialManager from the Client interface.

## Packet Structure

```bash
.
├── clients
│   ├── alpha_client.go
│   ├── betha_client.go
│   ├── client.go
│   └── client_test.go
├── credentials_manager
│   ├── credential_manager_one.go
│   ├── credential_manager_two.go
│   ├── credentials_manager.go
│   └── credentials_manager_test.go
├── logger
│   ├── logger.go
│   ├── logger_test.go
│   ├── log_one.go
│   └── log_two.go
└── README.md
```
