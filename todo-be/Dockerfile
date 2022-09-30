# STAGE 1
FROM golang:alpine as build
RUN apk update && apk add --no-cache git
WORKDIR /src
COPY . .
RUN go mod tidy
RUN go build -o sample-golang

# STAGE 2
FROM alpine
WORKDIR /app
COPY --from=build /src/sample-golang /app
ENTRYPOINT [ "/app/sample-golang" ]