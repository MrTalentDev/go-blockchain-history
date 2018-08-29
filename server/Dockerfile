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

#ENV PVT_KEY=
#ENV CONTRACT_ADDRESS=
#ENV LOGIN_DB=
#ENV PASSWORD_DB=
#ENV IP=
#ENV LOGIN=
#ENV PASSWORD=

RUN go build -o main .

EXPOSE 8080
EXPOSE 27017

CMD ["/app/main"]