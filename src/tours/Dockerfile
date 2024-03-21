FROM golang

WORKDIR /app

COPY . .

RUN go build -o tours main.go

CMD ["./tours"]