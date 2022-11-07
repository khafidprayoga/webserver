FROM golang:1.19 as build-env

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY . .

RUN go mod tidy -v
RUN go mod download && go mod verify
RUN go build -v -installsuffix cgo -o /dist/todo-app-fiber main.go
RUN chmod +x /dist/todo-app-fiber


FROM alpine:latest
RUN apk add --no-cache dumb-init

COPY --from=build-env /dist/todo-app-fiber app
COPY ./public /public
ENTRYPOINT ["/usr/bin/dumb-init", "./app"]