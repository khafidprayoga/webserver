FROM golang:1.19 as build-env

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY . .

RUN go mod tidy -v
RUN go mod download && go mod verify
RUN go build -v -installsuffix cgo -o /webserver main.go
RUN chmod +x /webserver


FROM alpine:latest
COPY --from=build-env /dist/webserver webserver
COPY ./public /public
ENTRYPOINT ["./webserver"]