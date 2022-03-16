FROM golang:alpine as builder
LABEL maintainer="me@yellowei.com"
COPY . /app
COPY src/views /app/views
COPY src/static /app/static
COPY src/conf /app/conf
WORKDIR /app
RUN apk update && apk add git \
    && export GO111MODULE=on  \
    && cd src \
    && go build -mod vendor -o ../app

# && go get -u github.com/gin-gonic/gin \
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/views .
COPY --from=builder /app/static .
COPY --from=builder /app/conf .
EXPOSE 8080
CMD ["./app"]
