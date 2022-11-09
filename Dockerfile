FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod init web
#RUN go build ./cmd/main.go -o main .

#CMD ["./cmd/main"]

#EXPOSE 8000