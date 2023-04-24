# Define a imagem base do Docker
FROM golang:latest

# Define o diretório de trabalho do Docker
WORKDIR /app

# Copia os arquivos do projeto para o diretório de trabalho
COPY . .

# Compila o projeto
RUN go build -o main .

# Expõe a porta em que a API será executada
EXPOSE 8080

# Define o comando para executar a API
CMD ["./main"]