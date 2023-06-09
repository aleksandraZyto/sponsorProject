FROM golang:1.19.0 

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

EXPOSE 3000

CMD go run main.go routes.go -b 0.0.0.0