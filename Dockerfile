FROM golang:1.19.0 
WORKDIR /chat-app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-app
EXPOSE 8080
CMD ["/chat-app"]