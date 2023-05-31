FROM golang:1.20-alpine

WORKDIR /usr/src/app

COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

ENV CONFIG_FILE=config/app.docker.yaml

COPY src .
RUN mkdir -p /usr/local/bin/
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app

CMD ["app"]
