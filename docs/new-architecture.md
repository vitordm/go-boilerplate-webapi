# New Architecture

# The Schema

web-app/
│
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
│
├── internal/
│   │
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── orders.go
│   │   │   ├── products.go
│   │   │   └── customers.go
│   │   ├── middleware/
│   │   ├── routes/
│   │   └── server/
│   │
│   ├── application/
│   │   ├── orders/
│   │   │   ├── create/
│   │   │   │   ├── command.go
│   │   │   │   ├── handler.go
│   │   │   │   └── response.go
│   │   │   ├── cancel/
│   │   │   │   ├── command.go
│   │   │   │   └── handler.go
│   │   │   └── get/
│   │   │       ├── query.go
│   │   │       ├── handler.go
│   │   │       └── response.go
│   │   │
│   │   ├── products/
│   │   └── customers/
│   │
│   ├── domain/
│   │   ├── orders/
│   │   │   ├── order.go
│   │   │   ├── order_item.go
│   │   │   ├── status.go
│   │   │   └── rules.go
│   │   ├── products/
│   │   └── customers/
│   │
│   ├── contracts/
│   │   ├── persistence/
│   │   │   ├── order_repository.go
│   │   │   ├── product_repository.go
│   │   │   └── customer_repository.go
│   │   ├── payments/
│   │   │   └── payment_gateway.go
│   │   ├── email/
│   │   │   └── email_sender.go
│   │   └── messaging/
│   │       └── message_publisher.go
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── order_repository.go
│   │   │       ├── product_repository.go
│   │   │       └── customer_repository.go
│   │   │
│   │   ├── payments/
│   │   │   ├── stripe/
│   │   │   └── mercadopago/
│   │   ├── email/
│   │   ├── messaging/
│   │   ├── cache/
│   │   └── logging/
│   │
│   ├── worker/
│   │   └── consumers/
│   │       ├── order_created.go
│   │       └── payment_approved.go
│   │
│   ├── composition/
│   │   ├── api.go
│   │   └── worker.go
│   │
│   └── shared/
│       ├── constants/
│       │   ├── errors.go
│       │   ├── headers.go
│       │   └── status.go
│       │
│       ├── utils/
│       │   ├── pointers.go
│       │   ├── strings.go
│       │   ├── numbers.go
│       │   ├── slices.go
│       │   └── time.go
│       │
│       └── errors/
│           └── errors.go
│
├── pkg/
│   └── contracts/
│       ├── orders/
│       ├── products/
│       └── customers/
│
├── migrations/
├── configs/
├── deployments/
├── tests/
│   ├── integration/
│   └── e2e/
│
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml


## Flow
cmd/api/main.go
       │
       ▼
internal/composition/api.go
       │
       ├── Infrastructure
       │       ├── PostgreSQL
       │       ├── Payment
       │       └── Messaging
       │
       ├── Application
       │       └── CreateOrderHandler
       │
       └── API
               ├── Routes
               ├── Middleware
               └── Handlers