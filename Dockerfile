FROM golang:1-alpine
ADD . /code

WORKDIR /code

RUN apk add ipmitool~=1 make~=4
