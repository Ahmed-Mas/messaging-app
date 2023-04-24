FROM golang:1.20-alpine as base
WORKDIR /app
COPY . .

FROM base as build
RUN go build -o main /app 

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
COPY --from=build /app/main .
ENTRYPOINT ["/app/main"]
