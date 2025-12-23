# Clean Architecture - Sistema de Pedidos

Sistema de gerenciamento de pedidos desenvolvido seguindo os princípios da Clean Architecture em Go.

## Arquitetura

O projeto segue a Clean Architecture com as seguintes camadas:

```
├── cmd/ordersystem/       # Ponto de entrada da aplicação
├── configs/               # Configurações da aplicação
├── internal/
│   ├── entity/            # Entidades de domínio
│   ├── usecase/           # Casos de uso da aplicação
│   └── infra/
│       ├── database/      # Repositórios de banco de dados
│       ├── web/           # Handlers REST
│       ├── grpc/          # Serviço gRPC
│       └── graph/         # Resolvers GraphQL
├── api/                   # Arquivos de API (api.http)
└── sql/migrations/        # Migrações do banco de dados
```

## Serviços e Portas

| Serviço | Porta | Descrição |
|---------|-------|-----------|
| REST API | 8000 | API REST para criar e listar pedidos |
| gRPC | 50051 | Serviço gRPC para criar e listar pedidos |
| GraphQL | 8080 | API GraphQL com playground |

## Pré-requisitos

- Docker
- Docker Compose
- Go 1.24+ (para desenvolvimento local)

## Como Executar

### 1. Subir a aplicação com Docker Compose

```bash
docker compose up -d
```

Este comando irá:
- Criar e iniciar o container MySQL
- Executar as migrações automaticamente
- Criar e iniciar o container da aplicação

### 2. Verificar se os serviços estão rodando

```bash
docker compose ps
```

## Endpoints

### REST API (Porta 8000)

#### Criar Pedido
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{"id": "pedido-1", "price": 100.50, "tax": 10.05}'
```

#### Listar Pedidos
```bash
curl http://localhost:8000/order
```

### gRPC (Porta 50051)

O serviço gRPC pode ser testado usando grpcurl ou Evans.

#### Usando grpcurl

```bash
# Listar serviços
grpcurl -plaintext localhost:50051 list

# Criar Pedido
grpcurl -plaintext -d '{"id": "pedido-grpc-1", "price": 100.50, "tax": 10.05}' \
  localhost:50051 pb.OrderService/CreateOrder

# Listar Pedidos
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

### GraphQL (Porta 8080)

Acesse o GraphQL Playground em: http://localhost:8080

#### Mutation - Criar Pedido
```graphql
mutation {
  createOrder(input: {id: "pedido-graphql-1", price: 100.50, tax: 10.05}) {
    id
    price
    tax
    finalPrice
  }
}
```

#### Query - Listar Pedidos
```graphql
query {
  orders {
    id
    price
    tax
    finalPrice
  }
}
```

## Arquivo api.http

O arquivo `api/api.http` contém exemplos de requisições para testar a API REST e GraphQL.

## Desenvolvimento Local

### Instalar dependências

```bash
go mod tidy
```

### Subir apenas o banco de dados

```bash
docker compose up mysql -d
```

### Executar a aplicação localmente

```bash
go run cmd/ordersystem/main.go
```

## Estrutura de Dados

### Pedido (Order)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| id | string | Identificador único do pedido |
| price | float64 | Preço do pedido |
| tax | float64 | Taxa do pedido |
| final_price | float64 | Preço final (price + tax) |

## Tecnologias Utilizadas

- **Go** - Linguagem de programação
- **MySQL** - Banco de dados relacional
- **Chi** - Router HTTP
- **gRPC** - Framework para comunicação RPC
- **gqlgen** - Biblioteca GraphQL para Go
- **Viper** - Gerenciamento de configurações
- **Docker** - Containerização
