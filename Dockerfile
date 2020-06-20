FROM golang:1.14

WORKDIR /app
COPY ./src/ .

RUN go build -o main .

EXPOSE 5500
CMD ["./main"]