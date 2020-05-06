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

Signup:

```
POST /auth/signup
Payload
{
	"username": "user3",
	"password": "123456",
	"fullname": "User 3"
}
```

Login:

```
POST /auth/login
Payload
{
	"username": "user1",
	"password": "123456"
}
```

Current User:

```
GET /me/
Header X-Access-Token=<accessToken>
```

User List:

```
GET /users/
Header X-Access-Token=<accessToken>
```

List messages:

```
GET /messages/<toUser>
Header X-Access-Token=<accessToken>
```

Write Message:

```
Post /messages/<toUser>
Header X-Access-Token=<accessToken>
Payload
{
	"message": "hello"
}
```

## Todos

- [ ] Add unit tests
- [ ] Add websocket support
