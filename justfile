set dotenv-load

alias d := dev

dev:
    air

build:
    go build -o ./build/event-glance ./cmd/event-glance

start: build
    ./build/event-glance