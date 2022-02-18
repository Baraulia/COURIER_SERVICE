FROM golang:alpine AS builder

COPY . /stlab.itechart-group.com/go/food_delivery/courier_service/
WORKDIR /stlab.itechart-group.com/go/food_delivery/courier_service/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /stlab.itechart-group.com/go/food_delivery/courier_service/.bin/service .
COPY --from=0 /stlab.itechart-group.com/go/food_delivery/courier_service/configs configs/

EXPOSE 81

CMD ["./service"]

