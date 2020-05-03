FROM golang:1.14-alpine as GO_BUILD
WORKDIR /app
ADD . /app
RUN go build -i -v ./cmd/main.go

FROM ubuntu:18.04
LABEL maintainer="zhshch<zhshch@athorx.com>"

WORKDIR /app
COPY --from=GO_BUILD /app /app

ENV RELEASE=T
ENV TZ=Asia/Shanghai
EXPOSE 4000

CMD ["/app/main"]