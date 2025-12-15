# Project Agent Information

This document outlines the technical details and conventions for this project, intended to be a guide for other agents (and humans!).

## Project Overview

- **Language:** Go 1.25
- **Framework:** Gin
- **API:** RESTful
- **Architecture:** Domain-Driven Design (DDD)
- **Database:** PostgreSQL
- **Deployment:** Docker Image

## Project Structure

The project follows a standard Go project layout:

- `cmd/`: Main applications for the project.
  - `api/`: The main REST API application.
  - `migrate/`: A command-line tool for running database migrations.
- `internal/`: Private application and library code.
  - `application/`: Application services that orchestrate business logic.
  - `config/`: Configuration loading and management.
  - `domain/`: Core domain models and repository interfaces.
  - `infrastructure/`: Infrastructure-level concerns.
    - `http/`: HTTP handlers and routing.
    - `persistence/`: Database repositories.
    - `migration/`: Database migration logic.
- `pkg/`: Library code that's ok to use by external applications (currently not used).
- `migration/`: Contains the SQL migration files.
- `mocks/`: Contains generated mocks for testing.
- `openapi/`: Contains the generated OpenAPI specification.
- `misc/`: Miscellaneous documentation and definitions.

## Configuration

The application is configured using environment variables. For local development, these variables are loaded from a `.env` file.

To set up your local configuration:
1.  Copy the `.env.dist` file to a new file named `.env`:
    ```bash
    cp .env.dist .env
    ```
2.  Modify the `.env` file as needed.

The configuration is loaded by the `internal/config` package.

## Tooling

- **Linter:** `golangci-lint` is used for linting the codebase.
- **Testing:** `mockery` is used for generating mocks for testing.
- **Migrations:** `golang-migrate/migrate` is used for database schema migrations.
- **API Documentation:** `swag` is used to generate OpenAPI documentation from code comments.

## Makefile Commands

A `Makefile` is provided to streamline common development tasks.

- `make build`: Compiles the main application binary.
- `make test`: Runs the test suite. It first runs `make mock` to ensure mocks are up-to-date.
- `make migrate`: Applies all database migrations.
- `make migrate-down`: Reverts all database migrations.
- `make up`: Starts the application and its dependencies in detached mode using `docker-compose`.
- `make down`: Stops the application and its dependencies.
- `make run`: Starts the application and its dependencies with a build, and runs migrations automatically.
- `make lint`: Runs the linter.
- `make mock`: Generates mocks for all interfaces in the project.
- `make swagger`: Generates OpenAPI documentation.