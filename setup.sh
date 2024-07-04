mkdir internal
mkdir internal/domains
touch internal/domains/.gitkeep
echo "this folder to define domain" > internal/domains/.gitkeep
mkdir internal/infrastructure
mkdir internal/infrastructure/database
mkdir internal/infrastructure/database/mysql
touch internal/infrastructure/database/mysql/.gitkeep
touch internal/infrastructure/database/mysql/mysql.go
echo "setup for external library to setup on application" > internal/infrastructure/database/mysql/.gitkeep
echo "package mysqldb" > internal/infrastructure/database/mysql/mysql.go
mkdir internal/infrastructure/logger
mkdir internal/infrastructure/http
touch internal/infrastructure/http/server.go
echo "package http" > internal/infrastructure/http/server.go
mkdir internal/interfaces
mkdir internal/interfaces/http
mkdir internal/interfaces/http/v1
touch internal/interfaces/http/v1/routes.go
mkdir internal/interfaces/grpc
mkdir internal/interfaces/cli
echo "package v1" > internal/interfaces/http/v1/routes.go
touch internal/interfaces/.gitkeep
echo "this is for interface implementation like to http or grpc interface" > internal/interfaces/.gitkeep
mkdir internal/usecase
touch internal/usecase/.gitkeep
echo "this is for implementation usecase form interface on domain" > internal/usecase/.gitkeep
mkdir internal/repository
touch internal/repository/.gitkeep
echo "this is for implementation repository form interface on domain" > internal/repository/.gitkeep
mkdir internal/app
touch internal/app/app.go
echo "package app" > internal/app/app.go
mkdir pkg
mkdir pkg/utils
mkdir pkg/httpserver
touch pkg/httpserver/options.go
touch pkg/httpserver/server.go
echo "package httpserver" > pkg/httpserver/server.go
mkdir pkg/httpclient
touch pkg/httpclient/options.go
touch pkg/httpclient/client.go
echo "package httpserver" > pkg/httpclient/client.go
mkdir cmd
touch cmd/main.go
echo "package main" > cmd/main.go
mkdir migrations
touch migrations/.gitkeep
echo "this folder to create schema migration db " > migrations/.gitkeep
mkdir config
touch config/config.go
touch config/example-config.yml
touch config/config.yml
echo "package config" > config/config.go
touch .gitignore
mkdir tools
mkdir tools/mocks
touch .env-example
touch .env
touch Makefile
touch README.md
echo "# create readme doc for this product" > README.md

echo ".idea\nfile\n.env\nconfig/config.yml\nlogs\n" > .gitignore
git init
go get github.com/ilyakaznacheev/cleanenv