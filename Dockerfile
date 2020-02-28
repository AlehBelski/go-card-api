FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

CMD ["go", "run", "main.go"]
