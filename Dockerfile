FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN go build -v -o /usr/local/bin/app ./...

ENV TELEGRAM_BOT_KEY ${TELEGRAM_BOT_KEY}

CMD ["go", "run", "main.go"]
