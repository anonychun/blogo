# Go Blog API

Real-world examples implementing back-end services for blog application using `Go` programming language to build `RESTful API`, all routes and API documentation can be found at `{{base_url}}/swagger/index.html`

## Architecture

This project using `Clean Architecture` with 4 domain layers:

- Model
- Repository
- Service
- Handler

## Features

- API Documentation `Swagger (auto generate)`
- Command line options
- Authentication `Json Web Token`
- CRUD operations `Postgres (raw sql)`
- Caching `Redis`
- Pagination, URL query search, etc
- Environment variables config
- Database `Migrations, Rollbacks, Steps, Drop, etc`
- Validation data request
- Middlewares `CORS, Rate Limit, Logger, Recover, Custom, etc`
- Graceful shutdown
- Many more...

## System Requirements

- Golang
- Docker
- Postgres (included in docker compose)
- Redis (included in docker compose)

## Running

Setting up all containers

```console
$ make compose.up
```

Applying all up migrations and start the server

```console
$ make launch
```

## Destroy

Applying all down migrations

```console
$ make rollbacks
```

Destroy running containers

```console
$ make compose.down
```

Delete all volumes that doesn't used

```console
$ docker volume prune
```

## Environment Variables

| **Key**                    | **Type** | **Value (Example)** |
| :------------------------- | :------- | :------------------ |
| APP_PORT                   | int      | 1401                |
| HTTP_RATE_LIMIT_REQUEST    | int      | 100                 |
| HTTP_RATE_LIMIT_TIME       | duration | 1s                  |
| JWT_SECRET_KEY             | string   | secret              |
| JWT_TTL                    | duration | 48h                 |
| PAGINATION_LIMIT           | int      | 100                 |
| POSTGRES_USER              | string   | admin               |
| POSTGRES_PASSWORD          | string   | secret              |
| POSTGRES_HOST              | string   | localhost           |
| POSTGRES_PORT              | int      | 3306                |
| POSTGRES_DATABASE          | string   | blog                |
| POSTGRES_MAX_IDLE_CONNS    | int      | 5                   |
| POSTGRES_MAX_OPEN_CONNS    | int      | 10                  |
| POSTGRES_CONN_MAX_LIFETIME | duration | 30m                 |
| REDIS_PASSWORD             | string   | secret              |
| REDIS_HOST                 | string   | localhost           |
| REDIS_PORT                 | int      | 6379                |
| REDIS_DATABASE             | int      | 0                   |
| REDIS_POOL_SIZE            | int      | 10                  |
| REDIS_TTL                  | duration | 1h                  |
