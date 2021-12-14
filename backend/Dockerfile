FROM golang:alpine
WORKDIR /app
EXPOSE 8080
COPY go.mod ./
RUN go get github.com/google/uuid
COPY *.go ./
RUN go build -o /main main.go
CMD ["/main"]