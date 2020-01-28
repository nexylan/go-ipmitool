FROM golang:1-alpine
COPY . /code

WORKDIR /code

RUN apk add --no-cache ipmitool~=1 make~=4
