# Use a imagem oficial do Golang como imagem base
FROM golang:1.21 AS builder

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para baixar as dependências
COPY go.mod go.sum ./app

# Execute o comando go mod download para baixar as dependências
RUN go mod download

# Copie todo o código-fonte da aplicação, incluindo a pasta cmd, para o contêiner
COPY . ./app

# Compile a aplicação (assumindo que o arquivo main.go esteja em cmd/)
RUN go build -o ./cmd

CMD ["/app/myapp"]
