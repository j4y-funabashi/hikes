FROM golang:1.25-alpine AS builder

RUN apk update && \
	apk add \
	build-base

WORKDIR /api

COPY . .
RUN go mod download
RUN go build -o app cmd/web/main.go

FROM alpine:3.23

RUN apk update && \
	apk add --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/ \
	bash \
    gpsbabel

WORKDIR /
COPY --from=builder /api/app /app
EXPOSE 8080

CMD ["/app"]