FROM golang

WORKDIR /app

COPY . .

RUN go build -o blogs main.go

CMD ["./blogs"]