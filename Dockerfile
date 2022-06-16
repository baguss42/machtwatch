FROM alpine:latest

COPY deploy/_output/service /usr/local/bin
COPY env.sample .env

RUN apk --no-cache add ca-certificates

EXPOSE 8080
ENTRYPOINT ["service"]