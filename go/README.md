# How to run without Docker Compose
This app expects redis to be listening at redis:6379. With that done, you can run this app as a standalone API listening on port 8090:
```
go mod download
go run server.go
```
