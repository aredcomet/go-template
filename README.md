# Go Template

This is a simple Go template project that can be used to start a new project.

## Features

- [ ] PostgresSQL Database
- [ ] Gin Web Framework
- [ ] PGX PostgresSQL Driver
- [ ] SQLC
- [ ] go-migrate
- [ ] Auth User tables
- [ ] JWT Authentication (similar to [djangorestframework-simplejwt](https://github.com/jazzband/djangorestframework-simplejwt))


## Getting Started
```bash
cp .env ~/Projects/envs/{{.ProjectName}}/local.env
go mod tidy
make runserver
```
