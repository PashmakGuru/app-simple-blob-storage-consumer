# Simple Blob Storage Consumer

A simple API-based app which identifies itself through Azure Entra ID and consumes Azure Blob Storage.

## Getting started

This project requires Go and Swag to be installed.

Copy the content of `.env.example` to `.env` and reflect the development configuration. Then run it using:

```bash
go get .
go run main.go
```

When your contribution is done, run the following command to generate the docs:
```bash
swag init
```

## TODO
- [ ] Move `swag init` to CI
- [ ] Move commands to a `Makefile`
