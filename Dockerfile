# Compile stage
FROM  golang:1.19-bullseye AS build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

ADD . /app
WORKDIR /app

RUN task

ENV SCOPE=go-fiber-app

ENTRYPOINT ["./build/program"]
