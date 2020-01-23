FROM golang AS build-env
ADD . /code

WORKDIR /code
