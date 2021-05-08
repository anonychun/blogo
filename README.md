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
- CRUD operations `MySQL / MariaDB (raw sql)`
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
- MySQL / MariaDB (included in docker compose)
- Redis (included in docker compose)

## Running

Setting up all containers

```
make compose.up
```

Applying all up migrations and start the server

```
make launch
```

## Destroy

Applying all down migrations

```
make rollbacks
```

Destroy running containers

```
make compose.down
```

Delete all volumes that doesn't used

```
docker volume prune
```

## Environment Variables

| **Key**                 | **Type** | **Value (Example)** |
| :---------------------- | :------- | ------------------: |
| APP_PORT                | int      |                1401 |
| HTTP_RATE_LIMIT_REQUEST | int      |                 100 |
| HTTP_RATE_LIMIT_TIME    | duration |                  1s |
| JWT_SECRET_KEY          | string   |         go_blog_api |
| JWT_TTL                 | duration |                 48h |
| PAGINATION_LIMIT        | int      |                 100 |
| MYSQL_USER              | string   |         go_blog_api |
| MYSQL_PASSWORD          | string   |         go_blog_api |
| MYSQL_HOST              | string   |           localhost |
| MYSQL_PORT              | int      |                3306 |
| MYSQL_DATABASE          | string   |         go_blog_api |
| MYSQL_MAX_IDLE_CONNS    | int      |                   5 |
| MYSQL_MAX_OPEN_CONNS    | int      |                  10 |
| MYSQL_CONN_MAX_LIFETIME | duration |                 30m |
| REDIS_PASSWORD          | string   |         go_blog_api |
| REDIS_HOST              | string   |           localhost |
| REDIS_PORT              | int      |                6379 |
| REDIS_DATABASE          | int      |                   0 |
| REDIS_POOL_SIZE         | int      |                  10 |
| REDIS_TTL               | duration |                  1h |
