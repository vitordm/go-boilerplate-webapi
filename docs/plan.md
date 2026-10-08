Claro. Eu faria um **plano incremental**, evitando tentar reorganizar tudo de uma vez. A ideia é chegar numa arquitetura coerente e, principalmente, **compilando e funcionando em cada etapa**.

## Plano — `new-architecture`

### Fase 1 — Limpeza da estrutura atual

**Objetivo:** eliminar restos da arquitetura antiga.

- [ ] Remover referências a `internal/app/...`
- [ ] Remover referências a `internal/core/...`
- [ ] Remover `internal/app/helpers/ioc`
- [ ] Remover `internal/core/ioc`
- [ ] Remover `internal/core/server`
- [ ] Remover `internal/core/cache`
- [ ] Remover `internal/core/utils`
- [ ] Remover `internal/app/startup.go`
- [ ] Remover estruturas antigas que ficaram sem uso
- [ ] Garantir que não existem imports apontando para a arquitetura antiga

**Resultado esperado:**

```text
cmd
internal/
├── api
├── application
├── composition
├── contracts
├── domain
├── infrastructure
├── shared
└── worker
pkg/
```

---

# Fase 2 — Composition Root

**Objetivo:** deixar claro onde a aplicação é montada.

Organizar:

```text
internal/composition/
└── api.go
```

E posteriormente:

```text
internal/composition/
├── api.go
└── worker.go
```

O `api.go` deve ser responsável por:

```text
Infrastructure
      ↓
Application
      ↓
API
      ↓
Echo
```

Por exemplo:

```text
PostgresRepository
       ↓
Application Handler
       ↓
API Handler
       ↓
Routes
```

### Regras

- [ ] Nenhum handler deve acessar `dig`
- [ ] Nenhum Application Handler deve acessar `dig`
- [ ] Nenhum componente deve criar suas próprias dependências
- [ ] `dig` fica restrito à composição
- [ ] `main.go` apenas inicia a aplicação

Ideal:

```go
func main() {
    app := composition.NewAPI()

    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
```

---

# Fase 3 — API

**Objetivo:** separar claramente HTTP do restante da aplicação.

Estrutura:

```text
internal/api/
├── example/
│   └── handler.go
├── middleware/
├── responses/
├── routes/
└── server/
    ├── server.go
    └── validator.go
```

### Fazer

- [ ] Corrigir `api/example/handler.go`
- [ ] Remover dependência direta do container
- [ ] Criar handlers com dependências explícitas
- [ ] Manter `responses` responsável somente por HTTP
- [ ] Mover `OutputRoutes` para `api/server/routes.go`, se ainda necessário
- [ ] `routes/routes.go` deve apenas registrar endpoints
- [ ] Remover `api/handlers/` se não houver necessidade real

Exemplo:

```go
type Handler struct {
    service example.Service
}

func NewHandler(service example.Service) *Handler {
    return &Handler{
        service: service,
    }
}
```

---

# Fase 4 — Application

**Objetivo:** transformar o Application na camada de casos de uso.

Estrutura:

```text
internal/application/
├── example/
│   ├── create/
│   │   ├── command.go
│   │   └── handler.go
│   └── get/
│       ├── query.go
│       └── handler.go
└── todo/
```

### Fazer

- [ ] Corrigir `GetExampleHandler`
- [ ] Corrigir `NewExampleService` → `NewHandler`
- [ ] Remover interfaces sem necessidade
- [ ] Criar `CreateExample`, se fizer parte do exemplo
- [ ] Criar `GetExample`
- [ ] Definir claramente Command vs Query
- [ ] Application não deve importar Echo
- [ ] Application não deve importar Postgres
- [ ] Application não deve importar `dig`

### Regra simples

```text
Command → altera estado
Query   → consulta estado
```

Não precisa criar Command/Query para absolutamente tudo.

Se algo for trivial:

```text
get/
└── handler.go
```

pode ser suficiente.

---

# Fase 5 — Contracts

**Objetivo:** estabelecer as fronteiras entre Application e Infrastructure.

Manter:

```text
internal/contracts/
├── persistence/
├── payments/
├── messaging/
└── email/
```

### Fazer

- [ ] Revisar `ExampleRepository`
- [ ] Revisar `TodoRepository`
- [ ] Definir interfaces realmente utilizadas
- [ ] Garantir que Application depende das interfaces
- [ ] Garantir que Infrastructure implementa as interfaces

Fluxo:

```text
Application
     ↓
contracts
     ↑
Infrastructure
```

### Importante

Não criar interfaces apenas porque "arquitetura limpa manda".

Se não existe uma necessidade real de abstração, não criar.

---

# Fase 6 — Infrastructure

**Objetivo:** concentrar implementações externas.

Estrutura:

```text
internal/infrastructure/
├── cache/
├── email/
├── logging/
├── messaging/
├── payments/
└── persistence/
    ├── postgres/
    └── repositories/
```

### Fazer

- [ ] Revisar repositories
- [ ] Revisar cache
- [ ] Revisar logging
- [ ] Preparar Postgres
- [ ] Manter integrações externas aqui
- [ ] Garantir que Infrastructure não vaza para Application

---

# Fase 7 — Remover duplicações

Essa fase é pequena, mas importante.

Atualmente temos coisas que conceitualmente estão duplicadas.

### Cache

Remover:

```text
internal/shared/cache.go
```

e usar:

```text
internal/infrastructure/cache/
```

### Logging

Remover:

```text
internal/shared/log.go
```

e usar:

```text
internal/infrastructure/logging/
```

### Resultado

`shared` fica realmente compartilhado:

```text
internal/shared/
├── constants/
├── errors/
└── utils/
```

---

# Fase 8 — Shared

**Objetivo:** impedir que `shared` vire um novo `core`.

Manter somente coisas como:

```text
shared/
├── constants/
├── errors/
└── utils/
```

### Fazer

- [ ] Revisar cada arquivo de `shared/utils`
- [ ] Confirmar que não conhece domínio
- [ ] Confirmar que não conhece HTTP
- [ ] Confirmar que não conhece banco
- [ ] Confirmar que não conhece AWS
- [ ] Confirmar que não conhece Echo
- [ ] Remover helpers específicos de negócio

Exemplo bom:

```go
func DefaultInt64(value *int64, defaultValue int64) int64
```

Exemplo ruim:

```go
func CalculateOrderTotal(...)
```

O segundo pertence ao domínio de Order.

---

# Fase 9 — Domain

**Objetivo:** garantir que o domínio seja realmente independente.

Estrutura:

```text
internal/domain/
├── todo/
│   └── todo.go
└── orders/
    ├── order.go
    ├── order_item.go
    └── status.go
```

### Fazer

- [ ] Remover dependências de infraestrutura
- [ ] Remover dependências HTTP
- [ ] Colocar regras de negócio aqui
- [ ] Colocar estados/status relacionados ao domínio aqui
- [ ] Evitar colocar lógica de negócio em `utils`

Regra:

```text
Domain
  ↓
não sabe como é salvo
não sabe como é transportado
não sabe qual banco existe
não sabe qual API externa existe
```

---

# Fase 10 — `pkg/contracts`

**Objetivo:** definir o que é realmente público/reutilizável.

Manter:

```text
pkg/contracts/
├── example/
└── todo/
```

### Fazer

- [ ] Identificar contratos que podem ser consumidos externamente
- [ ] Usar para OpenAPI/SDK quando fizer sentido
- [ ] Não transformar `pkg` em depósito
- [ ] Não colocar utilities genéricas aqui automaticamente

Regra:

> Se outro projeto importar esse pacote, isso é esperado?

Se a resposta for "não", provavelmente pertence ao `internal`.

---

# Fase 11 — Worker

Depois que API estiver funcionando:

```text
cmd/worker/
└── main.go

internal/worker/
└── consumers/
```

Criar:

```text
internal/composition/
├── api.go
└── worker.go
```

O Worker deve reutilizar:

```text
Application
Domain
Contracts
Infrastructure
```

mas **não deve depender da API HTTP**.

Fluxo:

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

---

# Fase 12 — Testes

Depois da arquitetura estar funcionando.

```text
tests/
├── integration/
└── e2e/
```

E testes próximos do código quando fizer sentido:

```text
application/example/get/
├── handler.go
└── handler_test.go
```

Prioridade:

1. [ ] Domain
2. [ ] Application
3. [ ] Infrastructure
4. [ ] API
5. [ ] Integration
6. [ ] E2E

Não precisa criar teste só para aumentar cobertura.

---

# Fase 13 — Documentação

Atualizar:

```text
docs/new-architecture.md
```

Eu colocaria uma seção curta explicando:

### Dependências

```text
API
 ↓
Application
 ↓
Domain

Application
 ↓
Contracts
 ↑
Infrastructure
```

### Composition

```text
cmd
 ↓
composition
 ↓
API / Worker
```

### Regras

- `internal` = uso interno
- `pkg` = reutilização externa intencional
- `shared` = somente código realmente genérico
- `application` = casos de uso
- `domain` = regras de negócio
- `infrastructure` = implementações externas
- `composition` = montagem da aplicação

---

# Ordem que eu realmente seguiria

Para não transformar isso numa refatoração interminável:

```text
1. Limpar arquitetura antiga
        ↓
2. Composition
        ↓
3. API
        ↓
4. Application
        ↓
5. Contracts
        ↓
6. Infrastructure
        ↓
7. Remover duplicações
        ↓
8. Domain
        ↓
9. pkg/contracts
        ↓
10. Worker
        ↓
11. Testes
        ↓
12. Documentação
```

### Critério de "terminado"

Eu consideraria essa arquitetura pronta quando conseguirmos demonstrar um fluxo completo:

```text
HTTP POST
   ↓
API Handler
   ↓
Application Command
   ↓
Application Handler
   ↓
Domain
   ↓
Repository Contract
   ↓
Postgres Repository
   ↓
Response
   ↓
HTTP 201
```

**sem `dig` aparecendo no meio do fluxo**, sem `Echo` entrando no Application/Domain e sem `Infrastructure` sendo conhecida diretamente pelo Application.

Esse fluxo funcionando é mais importante do que ter 40 pastas perfeitamente nomeadas.