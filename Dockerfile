FROM phusion/baseimage:latest

LABEL maintainer "galih.pratama2207@gmail.com"

# Add binary and assets
COPY ./go-backend-ggalihpp /apps/
COPY .env /apps/

WORKDIR /apps


ENTRYPOINT /apps/go-backend-ggalihpp

EXPOSE 5050