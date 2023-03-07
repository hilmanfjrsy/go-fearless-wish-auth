# go-fearless-wish-auth

A simple authentication with Go, based on [gin](https://github.com/gin-gonic/gin) and [gorm](https://github.com/go-gorm/gorm).

## Prerequisite
- Golang
- Docker
- Makefile

## Run Application
- `make docker-up` or `docker-compose up -d`
- `make run` or `go run main.go`

## Directory Structure

```
.
├── Makefile
├── README.md
├── config
│   └── database.go
├── controllers
│   └── authentication.go
├── datatransfers
│   └── user.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── auto_migrate.go
│   └── user.go
└── utils
    ├── auth.go
    └── response.go
```

## Authors
[@hilmanfjrsy](https://github.com/hilmanfjrsy)