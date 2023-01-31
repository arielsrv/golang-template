# Compile stage
FROM golang:1.19.5-bullseye AS build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

ADD . /app
WORKDIR /app

RUN task

ENV SCOPE=docker

EXPOSE 8080 8080

# ENTRYPOINT ["./build/program"]
