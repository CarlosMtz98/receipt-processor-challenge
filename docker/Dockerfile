FROM golang:1.21.0

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -v -o main ./cmd

EXPOSE 7070

CMD ["./main"]