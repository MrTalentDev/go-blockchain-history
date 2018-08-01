package main

import (
	"github.com/miguelmota/go-solidity-sha3"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {

	res := stringToKeccak("cc")

	resTwo := cellCreation(1, 1)

	fmt.Println(hex.EncodeToString(res))
	fmt.Println(hex.EncodeToString(resTwo))

	return
}

func stringToKeccak(data string) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String("cc"),
	)

	return hash
}

func intToKeccak(data int) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String(strconv.FormatInt(int64(data), 10)),
	)

	return hash
}

func cellCreation(assetID int, txNumber int) []byte {

	a := intToKeccak(txNumber)

	b := solsha3.SoliditySHA3(
		solsha3.Bytes32(a),
		solsha3.String(strconv.FormatInt(int64(assetID), 10)),
	)

	c := solsha3.SoliditySHA3(
		solsha3.Bytes32(b),
	)

	return c
	//address(keccak256(assetId + keccak256(txNumber)))
}
