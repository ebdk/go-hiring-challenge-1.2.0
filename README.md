# Go Hiring Challenge

This repository contains a Go application for managing products and their prices, including functionalities for CRUD operations and seeding the database with initial data.

## Project Structure

1. **cmd/**: Contains the main application and seed command entry points.

   - `server/main.go`: The main application entry point, serves the REST API.
   - `seed/main.go`: Command to seed the database with initial product data.

2. **app/**: Contains the application logic.
3. **sql/**: Contains a very simple database migration scripts setup.
4. **models/**: Contains the data models and repositories used in the application.
5. `.env`: Environment variables file for configuration.

## Setup Code Repository

1. Create a github/bitbucket/gitlab repository and push all this code as-is.
2. Create a new branch, and provide a pull-request against the main branch with your changes. Instructions to follow.

## Application Setup

- Ensure you have Go installed on your machine.
- Ensure you have Docker installed on your machine.
- Important makefile targets:
  - `make tidy`: will install all dependencies.
  - `make docker-up`: will start the required infrastructure services via docker containers.
  - `make seed`: ⚠️ Will destroy and re-create the database tables.
  - `make test`: Will run the tests.
  - `make run`: Will start the application.
  - `make docker-down`: Will stop the docker containers.

Follow up for the assignemnt here: [ASSIGNMENT.md](ASSIGNMENT.md)

## API Usage

- Catalog list
  - `GET /catalog?offset={int}&limit={int}&category={code}&price_lt={float}`
  - Defaults: `offset=0`, `limit=10` (clamped 1..100)
  - Example: `curl "http://localhost:8484/catalog?offset=0&limit=5&category=shoes&price_lt=20"`

- Product details
  - `GET /catalog/{code}`
  - Variants without price inherit the product price.
  - Example: `curl "http://localhost:8484/catalog/PROD001"`

- Categories
  - `GET /categories`
    - Example: `curl "http://localhost:8484/categories"`
  - `POST /categories`
    - Body: `{ "code": "bags", "name": "Bags" }`
    - 409 Conflict on duplicate code.
    - Example: `curl -X POST -H 'Content-Type: application/json' -d '{"code":"bags","name":"Bags"}' http://localhost:8484/categories`

## Architecture

This project follows a hexagonal architecture to keep boundaries clear and the code maintainable:

- Domain (core):
  - `domain/` contains entity types used across the app: `Product`, `Category`, `Variant`, and domain errors.
  - Handlers and repositories speak in terms of these types, not DB or HTTP.

- Ports (interfaces):
  - `app/catalog/port.go`, `app/categories/port.go` define service interfaces returning domain types.
  - Handlers depend on ports; repositories implement them.

- Adapters – Persistence:
  - `models/` contains GORM models (DB entities) and repositories.
  - Repos map DB models ⇄ domain entities (see `models/mappers.go`).
  - Duplicate key handling maps Postgres unique violations (SQLSTATE `23505`) to a domain error, keeping DB specifics inside the adapter.

- Adapters – HTTP:
  - `app/*/handler.go` contain handlers. They parse/validate inputs, call ports, and translate errors to HTTP codes.
  - DTOs and mappers in `app/*/dto.go` and `app/*/mapper.go` isolate HTTP payloads from domain types.
  - Common JSON helpers in `app/api/response.go` standardize success and error responses.

- Operational polish:
  - Request logging middleware and sane HTTP timeouts configured in `cmd/server/main.go`.
  - Pagination limits are validated and clamped; responses include `total` for list endpoints.
