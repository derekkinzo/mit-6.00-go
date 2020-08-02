FROM golang:1.14.3-alpine
RUN mkdir /code
WORKDIR /code
COPY . /code/