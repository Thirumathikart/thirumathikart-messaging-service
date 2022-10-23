## Build
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN go build -o messaging

## Dev
FROM build AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest


ENTRYPOINT make watch


## Prod
FROM alpine:latest AS prod

WORKDIR /

COPY --from=build /app/messaging  /app/.env  /


ENTRYPOINT ["./messaging"]