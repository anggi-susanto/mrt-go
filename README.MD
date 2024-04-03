# MRT Waste Water API

## Local development setup

Dependencies:
- Go 1.21
- MongoDB 10.6 or latest
- Swaggo/Swagg (for make api doc)

Create `.env` file from `.env.example` and fill your `DB_DSN` in there.

Install Go dependencies with `go mod download`.

## Useful make commands:

- `make dev`
- `make api-docs`
- `make unit-test`
- `make coverage`

## Running locally

Run `make api-docs; make dev`

## Access the API
http://127.0.0.1:3000/docs/index.htm