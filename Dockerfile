FROM golang:1.26

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

COPY . .

CMD ["air"]
