#!/bin/bash

echo "This project implements Domain-Driven Design (DDD) principles in Go. It follows a layered architecture, separating the application into distinct layers for domain logic, infrastructure, and interfaces. The domain layer defines the core business rules and entities, while the infrastructure layer handles data access and external interactions. The application layer orchestrates the use cases, and the interface layer defines how the application interacts with the outside world. This structure enables a clean separation of concerns, making the system more maintainable, scalable, and testable." > README.md

if [ ! -d .git ]; then
    git init
fi

echo "/*/**/README.md" >> .gitignore

if [ ! -f go.mod ]; then
    projectName=ddd-implement
    if [ ! -z "$1" ]; then 
        projectName="$1"
    fi
    go mod init $projectName
fi

# tak perlu unit test
mkdir -p cmd
echo "package main" > cmd/main.go

mkdir -p config
echo "package config" > config/config.go
echo "" > config/config.yml.example

mkdir -p internal/app
echo "package app" > internal/app/app.go
echo "This directory contains the application logic of the system. It defines the structure and behavior of the application." > internal/app/README.md
# -------------------------

mkdir -p internal/domain
echo "This directory contains the domain logic of the system. It defines the core business rules and entities of the application." > internal/domain/README.md

mkdir -p internal/infrastructure/delivery/http
echo "package router" > internal/infrastructure/delivery/http/router.go
echo "This directory contains the HTTP delivery layer of the system. It handles incoming HTTP requests and sends responses." > internal/infrastructure/delivery/http/README.md

mkdir -p internal/infrastructure/repository
echo "This directory contains the repository layer of the system. It handles the data access and manipulation." > internal/infrastructure/repository/README.md

mkdir -p internal/infrastructure/services
echo "This directory contains the service layer of the system. It encapsulates the business logic and provides a layer of abstraction between the application and the infrastructure." > internal/infrastructure/services/README.md

mkdir -p internal/infrastructure/usecase
echo "This directory contains the use case layer of the system. It defines the actions that can be performed on the domain entities." > internal/infrastructure/usecase/README.md

mkdir -p pkg/database
echo "This directory contains the database layer of the system. It handles the interaction with the database and provides data storage and retrieval capabilities." > pkg/database/README.md

mkdir -p pkg/httpclient
echo "package httpclient" > pkg/httpclient/client.go
echo "package httpclient" > pkg/httpclient/options.go
echo "This directory contains the HTTP client layer of the system. It handles outgoing HTTP requests and receives responses." > pkg/httpclient/README.md

mkdir -p pkg/httpserver
echo "package httpserver" > pkg/httpserver/server.go
echo "package httpserver" > pkg/httpserver/options.go
echo "This directory contains the HTTP client layer of the system. It handles outgoing HTTP requests and receives responses." > pkg/httpclient/README.md

mkdir -p pkg/utils
echo "package utils" > pkg/utils/utils.go
echo "This directory contains utility functions and tools that can be used throughout the system." > pkg/utils/README.md

mkdir -p proto
echo "This directory contains the protocol buffer files of the system. It defines the structure and behavior of the data interchange format." > proto/README.md
