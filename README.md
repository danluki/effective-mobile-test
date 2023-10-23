### Effective mobile test task

## To run app in docker

This will run migrations, and application will be accessable on port 8000

```sh
docker compose up
```

## Technologies

- Golang 1.21. Cus I already like slog, and some of new features
- Goose for migrations. I think it's better then atlas/go-migrate cus of handwritten migrations and simple api
- Gin as web framework. Not the fastest, but with tons of examples and ready to go solutions.
- Gorm as ORM. I decided to try gorm for the first time and it looks cool. But I prefer more sqlc.
- Taskfile instead of Makefile.
- Golangci-lint

## Why this is cool

- Deployment ready.
- Dockerized with small image size(not so haha).
- Github actions.
- Clean architecture/DDD.
- Ready for unit/integrations tests.
- Mocks.
- Linter.
- Slog as logger with custom implementation.
- Hot reload.
