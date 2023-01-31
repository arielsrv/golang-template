# Compile stage
FROM golang:1.19.5-bullseye AS build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

ADD . /app
WORKDIR /app

RUN task

ENV SCOPE=go-fiber-app

EXPOSE 8080

CMD ["./app/build/program"]
