## Build
# Use uma imagem base mínima e específica para seu aplicativo Go
#FROM golang:1.22.0 AS builder
# FROM golang:1.22.0-alpine AS builder 
FROM golang:1.22.0-alpine3.19 AS builder 
# Define o diretório de trabalho dentro do contêiner
WORKDIR /app
#  install CA certificates
RUN apk update 
RUN update-ca-certificates
# Limpa o cache do módulo Go
RUN go clean --modcache
# Copia o arquivo go.mod e go.sum para baixar as dependências
COPY go.mod go.sum ./
# # Baixa e instala as dependências do módulo Go
# # Tenta resolver as dependências antes de executar o go build
RUN go mod download && go mod verify
# Copia o restante do código-fonte para o diretório de trabalho
COPY . .
# Compila o aplicativo Go para um executável estático
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ${MAIN_PATH}

## Deploy
# Cria uma imagem mínima sem dependências para executar o aplicativo
FROM scratch AS final
#FROM alpine:3.19.1 AS final
# Define o diretório de trabalho dentro do contêiner
WORKDIR /app
# Copia o executável compilado do estágio de construção para o estágio final de u  contêiner mínimo
# Copy the binary from the build stage
COPY --from=builder /app/main .
#EXPOSE 80
# Define o comando padrão a ser executado quando o contêiner é iniciado
CMD ["./main"] 