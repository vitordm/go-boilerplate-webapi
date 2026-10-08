# Go Boilerplate Web API

A pragmatic and scalable Go boilerplate for building Web APIs with a modular architecture, dependency injection, clear separation of responsibilities, and room to grow without unnecessary complexity.

The goal of this project is not to implement a specific architecture pattern by the book. Instead, it provides a practical structure that keeps the code easy to navigate, test, maintain, and extend.

---

## Architecture

The project follows a modular, feature-oriented architecture with clear boundaries between HTTP, application logic, domain rules, infrastructure, and external contracts.

```text
cmd/
├── api/
│   └── main.go
└── worker/
    └── main.go

internal/
├── api/
├── application/
├── composition/
├── contracts/
├── domain/
├── infrastructure/
├── shared/
└── worker/

pkg/
└── contracts/
```

The main dependency direction is:

```text
                ┌──────────────┐
                │     API      │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │ Application  │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │    Domain    │
                └──────────────┘

                Application
                     │
                     ▼
                 Contracts
                     ▲
                     │
               Infrastructure
```

Infrastructure implements contracts required by the application.

---

## Project Structure

### `cmd/`

Contains the executable entry points of the application.

```text
cmd/
├── api/
│   └── main.go
└── worker/
    └── main.go
```

Each executable should have a very small `main.go`.

The executable is responsible for starting the application, while dependency composition is handled by `internal/composition`.

Example:

```go
func main() {
    app := composition.NewAPI()

    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
```

---

### `internal/api/`

Contains everything related to the HTTP API.

```text
internal/api/
├── example/
├── middleware/
├── responses/
├── routes/
└── server/
```

Responsibilities include:

- HTTP handlers
- Middleware
- Route registration
- HTTP response helpers
- Request validation
- HTTP server configuration

The API layer knows about HTTP and the web framework.

The Application and Domain layers should not depend on HTTP concepts.

---

### `internal/application/`

Contains application use cases.

The organization is feature-first:

```text
internal/application/
├── example/
│   ├── create/
│   └── get/
└── todo/
```

A feature can contain multiple operations:

```text
orders/
├── create/
├── cancel/
├── get/
└── update/
```

Each operation should contain only the files it actually needs.

For example:

```text
create/
├── command.go
└── handler.go
```

or, for a simple operation:

```text
get/
└── handler.go
```

### Command vs Query

Commands represent operations that change state.

```text
CreateOrder
CancelOrder
UpdateOrder
```

Queries retrieve information.

```text
GetOrder
ListOrders
```

This distinction is useful for organizing the application, but it should not become ceremony.

Do not create additional abstractions just to follow a pattern.

---

### `internal/domain/`

Contains business rules and domain models.

```text
internal/domain/
├── orders/
├── products/
├── customers/
└── todo/
```

The Domain should remain independent from infrastructure and transport concerns.

It should not know about:

- Echo
- HTTP
- PostgreSQL
- Redis
- AWS
- RabbitMQ
- Stripe
- Mercado Pago

For example, an `Order` should contain rules related to an order, but should not know how an order is persisted.

---

### `internal/contracts/`

Contains interfaces that define boundaries required by the application.

```text
internal/contracts/
├── persistence/
├── payments/
├── messaging/
└── email/
```

For example:

```go
type OrderRepository interface {
    Save(ctx context.Context, order *Order) error
    GetByID(ctx context.Context, id string) (*Order, error)
}
```

The application depends on the contract.

Infrastructure provides the implementation.

```text
Application
     ↓
OrderRepository
     ↑
PostgresOrderRepository
```

Contracts should describe what the application needs, rather than expose infrastructure details.

---

### `internal/infrastructure/`

Contains concrete implementations and integrations with external systems.

```text
internal/infrastructure/
├── cache/
├── email/
├── logging/
├── messaging/
├── payments/
└── persistence/
```

Examples:

```text
persistence/
├── postgres/
└── repositories/

payments/
├── stripe/
└── mercadopago/

messaging/
├── rabbitmq/
└── sqs/
```

Infrastructure is responsible for implementation details such as:

- PostgreSQL
- Redis
- Message brokers
- Payment providers
- Email providers
- External APIs
- Logging infrastructure

The application should not directly instantiate these implementations.

---

### `internal/composition/`

This is the composition root of the application.

It is responsible for assembling the application and connecting interfaces with their implementations.

For example:

```text
Postgres Repository
        ↓
Application Handler
        ↓
API Handler
        ↓
Routes
```

The dependency injection mechanism is an implementation detail of this layer.

The rest of the application should not need to know whether dependencies are created using:

- `dig`
- manual dependency injection
- another DI library

This makes it possible to change the DI mechanism without changing business logic.

Typical structure:

```text
internal/composition/
├── api.go
└── worker.go
```

The API and Worker can share the same Application, Domain, Contracts, and Infrastructure while having different entry points and composition.

---

### `internal/worker/`

Contains background processing and message consumers.

```text
internal/worker/
└── consumers/
```

The Worker should reuse application use cases instead of duplicating business logic.

Example:

```text
Message Broker
      ↓
Worker Consumer
      ↓
Application Handler
      ↓
Domain
      ↓
Infrastructure
```

The Worker does not depend on the HTTP API.

---

### `internal/shared/`

Contains genuinely generic code shared by multiple parts of the application.

```text
internal/shared/
├── constants/
├── errors/
└── utils/
```

Examples include generic helpers such as:

```go
func DefaultInt64(value *int64, defaultValue int64) int64 {
    if value == nil {
        return defaultValue
    }

    return *value
}
```

`shared` should not become a replacement for a generic `core`, `common`, or `helpers` folder.

Business-specific logic should remain close to the feature that owns it.

Infrastructure-specific code belongs in `infrastructure`.

---

### `pkg/`

Contains packages that are intentionally reusable outside of the `internal` application.

Currently:

```text
pkg/
└── contracts/
```

For example:

```text
pkg/contracts/
├── orders/
├── products/
└── customers/
```

This can be useful for:

- Public API contracts
- OpenAPI generation
- SDK generation
- Shared contracts between repositories
- Types intentionally consumed by external packages

Code should only be placed under `pkg` when external reuse is intentional.

`pkg` should not be used as a generic dumping ground for models, helpers, or utilities.

---

## Dependency Rules

The architecture follows a few simple rules.

### API

The API knows about:

- HTTP
- Echo
- Request/Response
- Middleware
- Routes

The API calls Application use cases.

---

### Application

The Application knows about:

- Use cases
- Commands
- Queries
- Domain
- Contracts

The Application should not know about:

- HTTP
- Echo
- PostgreSQL
- Redis
- AWS
- Stripe
- RabbitMQ

---

### Domain

The Domain knows about business rules.

It should not know about:

- HTTP
- Databases
- Infrastructure
- Message brokers
- External APIs

---

### Infrastructure

Infrastructure knows how to communicate with external systems.

It implements the contracts required by the Application.

---

### Composition

Composition is where concrete implementations are assembled.

This is the only layer that should need to know which concrete implementation is being used.

---

## Dependency Injection

Dependencies should be injected rather than discovered by application components.

Prefer:

```go
type Handler struct {
    repository Repository
}

func NewHandler(repository Repository) *Handler {
    return &Handler{
        repository: repository,
    }
}
```

Instead of:

```go
func (h *Handler) Execute() {
    repository := container.Resolve(...)
}
```

Application components should not access the DI container directly.

The container belongs to the composition layer.

---

## Interfaces

Interfaces should be introduced when they provide a meaningful boundary.

Do not create interfaces automatically for every struct.

Avoid unnecessary abstractions such as:

```go
type UserService interface {
    ...
}

type UserManager interface {
    ...
}

type UserProcessor interface {
    ...
}
```

when there is no real reason for them to exist.

A concrete type is often enough.

Interfaces are particularly useful at boundaries such as:

```text
Application
     ↓
Repository Contract
     ↑
Infrastructure
```

---

## Feature-First Organization

The project prefers feature-oriented organization over global folders based only on technical type.

Prefer:

```text
application/
└── orders/
    ├── create/
    ├── cancel/
    └── get/
```

over:

```text
application/
├── handlers/
├── commands/
├── queries/
├── services/
└── responses/
```

The first structure makes it easier to answer:

> "Where is the code responsible for creating an order?"

without searching through several unrelated folders.

---

## Avoiding Overengineering

This project intentionally avoids architecture for architecture's sake.

The following should not be added automatically:

- Generic repositories
- Generic services
- Generic managers
- Generic factories
- Generic builders
- Multiple DTO mappings with no benefit
- Interfaces with a single meaningless implementation
- Empty abstraction layers
- Global `Common` or `Core` folders
- Utilities that contain business logic

A useful rule is:

> An abstraction should solve a real problem.

If removing an abstraction makes the code simpler without losing an important boundary, the abstraction probably does not need to exist.

---

## HTTP Request Flow

A typical API operation should look like:

```text
HTTP Request
     ↓
API Handler
     ↓
Command / Query
     ↓
Application Handler
     ↓
Domain
     ↓
Contract
     ↓
Infrastructure
     ↓
Application Response
     ↓
API Response
     ↓
HTTP Response
```

For example:

```text
POST /orders
      ↓
CreateOrder HTTP Handler
      ↓
CreateOrderCommand
      ↓
CreateOrderHandler
      ↓
Order Domain
      ↓
OrderRepository
      ↓
PostgresOrderRepository
      ↓
CreateOrderResponse
      ↓
HTTP 201
```

The API layer is responsible for HTTP concerns such as status codes.

The Application layer should not return HTTP status codes.

---

## Error Handling

Errors should be translated at the appropriate boundary.

For example:

```text
Domain Error
     ↓
Application
     ↓
API Error Handler
     ↓
HTTP Status Code
```

The Domain should not return HTTP-specific errors such as:

```go
http.StatusBadRequest
```

Instead, it should expose meaningful domain/application errors.

The API layer decides how those errors are represented over HTTP.

---

## Testing

Tests should focus on behavior rather than implementation details.

Recommended priorities:

```text
Domain
   ↓
Application
   ↓
Infrastructure
   ↓
API
   ↓
Integration
   ↓
E2E
```

Not every component needs every type of test.

Tests should provide useful confidence rather than exist only to increase coverage percentages.

---

## Running the Application

### API

```bash
go run ./cmd/api
```

### Worker

```bash
go run ./cmd/worker
```

### Tests

```bash
go test ./...
```

### Build

```bash
go build ./...
```

---

## Development Principles

This project follows a few practical principles:

1. **Prefer simple code over clever code.**
2. **Keep business rules close to the domain that owns them.**
3. **Use feature-first organization.**
4. **Keep infrastructure behind meaningful boundaries.**
5. **Inject dependencies instead of discovering them.**
6. **Keep HTTP concerns inside the API layer.**
7. **Keep the Domain independent.**
8. **Use interfaces where they provide real value.**
9. **Avoid unnecessary abstractions.**
10. **Make the code easy for another developer to navigate.**

The architecture should evolve with the application.

It should not become a constraint that prevents simple solutions.

---

## Project Status

This repository is intended to be a starting point for Go Web API projects.

The structure is intentionally opinionated, but individual projects may simplify or expand it based on their actual requirements.

The most important goal is not to have the most sophisticated architecture.

The goal is to have an architecture where a developer can quickly answer:

- Where does this request enter the system?
- Where is the use case implemented?
- Where are the business rules?
- Where is the database implementation?
- Where are external integrations configured?
- Where are dependencies assembled?

If those questions can be answered quickly, the architecture is doing its job.