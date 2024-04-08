FROM golang:alpine as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o followers main.go

FROM alpine
COPY --from=build /app/followers /usr/bin
EXPOSE 8084
ENTRYPOINT ["followers"]