# FileStorage API

Cloud storage for files written in golang

## Used frameworks
- Gin
- swag (auto swagger)
- go-jwt (authentication)
- sqlc (auto sql queries generate)

## Routes
Documentation available in `/swagger/index.html`
- GET `/files` Get files list
- GET `/files/:filename` Get file body
- POST `/files` Store file
- POST `/auth/login` Log in account
- POST `/auth/register` Create an account

## Run
For development: `go run main.go`
FOr production: `go build && ./filestorage`
