# Messenger Challenge

I developed a simple Messenger API with golang to make a practice.

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

## API Documentation

Postman Documentation: https://documenter.getpostman.com/view/191558/SzmcZe7s?version=latest

## Todos

- [ ] Add unit tests
- [ ] Add websocket support
- [ ] Add Github Actions
