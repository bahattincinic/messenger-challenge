# Messenger Challenge

I developed a simple Messenger API with golang to make a practice.

## Requirements
- Golang
- The Clean Architecture (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Folder Structure

```
.
├── README.md
├── api                   # HTTP API is under this folder
│   ├── handlers          # API controller/handlers
│   ├── middlewares       # API middlewares
│   └── routes.go         # All Api Routes
├── config                # Project Configurations
│   ├── config.go
│   ├── development.yaml
│   └── production.yaml
├── domain                # All business logic related things are under this folder.
│   ├── models            # Database and data models
│   ├── repositories      # Database query layer
│   └── usecases          # business usecases
├── go.mod
├── go.sum
├── main.go
└── messenger.db
```

## Installation

```
$ git clone git@github.com:bahattincinic/messenger-challenge.git
$ cd messenger-challenge
$ go mod download
```

## Running Unit Tests

```
$ go test ./... -coverprofile cp.out

?   	github.com/bahattincinic/messenger-challenge	[no test files]
?   	github.com/bahattincinic/messenger-challenge/api	[no test files]
?   	github.com/bahattincinic/messenger-challenge/api/handlers	[no test files]
ok  	github.com/bahattincinic/messenger-challenge/api/middlewares	0.284s	coverage: 17.6% of statements
ok  	github.com/bahattincinic/messenger-challenge/config	0.299s	coverage: 90.0% of statements
?   	github.com/bahattincinic/messenger-challenge/domain/models	[no test files]
ok  	github.com/bahattincinic/messenger-challenge/domain/repositories	0.373s	coverage: 63.0% of statements
ok  	github.com/bahattincinic/messenger-challenge/domain/usecases	0.564s	coverage: 73.1% of statements
```

## API Documentation

Postman Documentation: https://documenter.getpostman.com/view/191558/SzmcZe7s?version=latest
