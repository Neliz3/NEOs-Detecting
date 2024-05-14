FROM golang:latest

ARG API_KEY

COPY ./ ./


ENV HOME /

ENTRYPOINT ["{API_KEY}"]

WORKDIR /NEOs/
CMD ["go", "run", "./cmd/main.go"]