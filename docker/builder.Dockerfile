FROM golang:latest

WORKDIR /baumanfinder

ENV PROJ_ROOT=/baumanfinder

COPY . .

RUN apt-get update && \
    make build-app