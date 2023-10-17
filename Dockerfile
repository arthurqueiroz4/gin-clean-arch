 # Etapa de build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# RUN --mount=type=cache,target=/go/pkg/mod go mod download

ADD . /app

RUN go build -o main cmd/main.go



FROM builder as runner

# RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o main ./cmd/main.go

WORKDIR /app

COPY --from=builder /app/main ./

EXPOSE 8080

CMD ["/app/main"]

