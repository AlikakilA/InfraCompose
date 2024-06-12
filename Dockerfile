FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY hangman ./hangman

COPY main1.go .

RUN go build -o main ./

EXPOSE 8080

CMD ["./main"]