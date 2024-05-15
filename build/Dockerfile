FROM golang:latest

ARG API_KEY

COPY ./ /app

WORKDIR /app

ENV HOME=/
ENV API_KEY=${API_KEY}

CMD ["go", "run", "./cmd/main.go"]
