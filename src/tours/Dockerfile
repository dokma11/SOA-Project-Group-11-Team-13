FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o tours main.go

FROM alpine
COPY --from=build /app/tours /usr/bin
ENTRYPOINT [ "tours" ]