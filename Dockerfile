FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /steam-game-version-manager cmd/app/main.go

CMD [ "/steam-game-version-manager" ]