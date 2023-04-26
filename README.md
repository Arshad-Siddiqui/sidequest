# Sidequest

A simple CRUD API for managing your side projects.

## Run in developer mode

You will need to have codegangsta/gin installed.

```shell
go get github.com/codegangsta/gin
```

Then run the following command to start the server.

```shell
gin -i --appPort 8080 --port 3000 --all run main.go
```

## When running tests locally

Make sure to migrate the database before running the tests. Especially when you modify any schemas.

```shell

go run migrate/migrate.go

```
