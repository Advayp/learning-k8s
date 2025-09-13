FROM golang:1.24.0

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello

EXPOSE 5051 

CMD ["./hello"]