# Donation Campaign API

This is the API for the Donation Campaign application.

## Overview

- **Language:** Go 1.25
- **Framework:** Gin
- **API:** RESTful
- **Architecture:** Domain-Driven Design (DDD)
- **Database:** PostgreSQL
- **Deployment:** Docker Image

## Getting Started

### Prerequisites

- Docker
- Go 1.25

### Configuration

Before running the application, you need to set up your environment variables.

1.  Copy the `.env.dist` file to a new file named `.env`:
    ```bash
    cp .env.dist .env
    ```

2.  Open the `.env` file and modify the environment variables as needed for your local setup. The default values are configured to work with the provided `docker-compose.yml`.

### Running the application

1.  **Set up your environment:**
    Make sure you have a `.env` file as described in the "Configuration" section.

2.  **Start the application:**
    ```bash
    make run
    ```
    This will start the application and the database using Docker Compose, and will also run the migrations.

3.  **Check the health of the application:**
    Open your browser and navigate to `http://localhost:8080/health`.

### Makefile Commands

A `Makefile` is provided to streamline common development tasks.

- `make build`: Compiles the application.
- `make test`: Runs the test suite (generates mocks first).
- `make migrate`: Applies database migrations.
- `make migrate-down`: Reverts database migrations.
- `make up`: Starts the application in detached mode.
- `make down`: Stops the application.
- `make run`: Starts the application with a build, and runs migrations.
- `make lint`: Runs the linter.
- `make mock`: Generates mocks for all interfaces.
- `make swagger`: Generates OpenAPI documentation.

## Migrations

The `./migration` directory contains database migrations. These are now run automatically when you start the application with `make run`. You can also run them manually:

To apply migrations, run:
`make migrate`

To revert migrations, run:
`make migrate-down`