FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod init graylog
RUN go mod tidy
RUN go build ./cmd/main.go

#EXPOSE 8000