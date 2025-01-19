# Utiliza a imagem oficial do Golang como base
FROM golang:1.20-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa todas as dependências necessárias
RUN go mod download

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila a aplicação
RUN go build -o main .

# Define a porta que a aplicação irá escutar
EXPOSE 8080

# Comando para rodar a aplicação quando o container iniciar
CMD ["./main"]
