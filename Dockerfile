FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN go build -o /test-cloud-run
#EXPOSE 8080
CMD ["/test-cloud-run"]