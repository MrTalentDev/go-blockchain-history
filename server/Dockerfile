FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/miguelmota/go-solidity-sha3
RUN go get github.com/ethereum/go-ethereum
RUN go get github.com/ethereum/go-ethereum/accounts/abi
RUN go get github.com/ethereum/go-ethereum/accounts/abi/bind
RUN go get github.com/ethereum/go-ethereum/common
RUN go get github.com/ethereum/go-ethereum/core/types
RUN go get github.com/ethereum/go-ethereum/event
RUN go get github.com/ethereum/go-ethereum/ethclient
RUN go get github.com/ethereum/go-ethereum/crypto

ENV PVT_KEY=fb4346d4d7518c159c809f62fb6ab0152f7ca878bf7fddc056d6b02a3c917653
ENV CONTRACT_ADDRESS=0x0293a80682dc2a192c683baf434dd67343cebf70
ENV LOGIN_DB=bankex2
ENV PASSWORD_DB=ObshiDostup1
ENV IP=18.209.40.150
ENV LOGIN=bankex
ENV PASSWORD=ObshiDostup1

RUN go build -o main .

EXPOSE 8080
EXPOSE 27017

CMD ["/app/main"]