# Gin REST Starter

Minimum [Gin](https://github.com/gin-gonic/gin) starter template for REST API.

This template uses 
- [GORM](https://github.com/go-gorm/gorm) for the ORM.
- [swag](https://github.com/swaggo/swag) to auto generate Swagger Documentation

By default it uses Postgres connection. If you are using different Database, you can easily change the dsn, check [GORM: Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

## Getting started

This template uses [air](https://github.com/cosmtrek/air) with hot reload feature in development environment. Read the installation guide [here](https://github.com/cosmtrek/air#installation).

1. Copy `.env.example` to `.env` and modify the env value as needed
2. `go mod download` to install all required dependencies
3. `air .air.toml`: start a dev server with hot reload (auto updates swagger)
4. `go test -v ./...`: run all tests on all packages

**Note:** set `GO_ENV` to `production` to run Gin app in release mode

To access swagger documentation, open http://localhost:3000/swagger/index.html on your browser. Read more about swag's comment annotation [here](https://github.com/swaggo/swag#declarative-comments-format)

#### Route example:
- `GET http://localhost:3000/api/compA`: `"Hello World"`
- `GET http://localhost:3000/api/compA/3/5`: 
```json
{
  "result": 8
}
```

## Directory Structure

```
.
├── compA
│   ├── models.go
│   ├── router.go
│   ├── service.go
│   └── service_test.go
├── compB
│   └── ...
├── main.go
├── go.mod
└── ...
```

This template separate your code by your app's components, for example for a social media app, you will have something like this:

```
.
├── posts
│   └── ...
├── users
│   └── ...
├── main.go
├── go.mod
└── ...
```

Each component will usually contains 4 files:
- `models.go`: contains the component's [GORM Models](https://gorm.io/docs/models.html) (notice that it's `models.go`, not `model.go` it's because one component can have more than one model, for example for `posts` component, you can have `Post` model and `Comment` model, though you can also separate `comments` into it's own component folder)
- `router.go`: contains api routes and it's controller
- `service.go`: contains the business logic of the component (avoid putting business logic directly in router to make testing easier)
- `service_test.go`: contains tests for your `service.go`
