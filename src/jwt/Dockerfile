FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o jwt main.go

FROM alpine
COPY --from=build /app/jwt /usr/bin
ENTRYPOINT ["jwt"]