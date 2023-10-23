##
## Build
##
FROM golang:1.21-alpine as dev

WORKDIR /app
COPY go.mod go.sum /app/

RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY . .

RUN task build

##
##  Production
##
FROM alpine:latest as prod

WORKDIR /app



COPY --from=dev /app/configs/conf.docker.yml configs/conf.yml
COPY --from=dev /app/internal/database/migrations migrations
COPY --from=dev /app/build-out ./app
CMD ["./app"]