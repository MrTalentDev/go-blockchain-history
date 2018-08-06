package hashing

import (
	"github.com/miguelmota/go-solidity-sha3"
	"strconv"
)

func StringToKeccak(data string) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String(data),
	)

	return hash
}

func IntToKeccak(data int) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String(strconv.FormatInt(int64(data), 10)),
	)

	return hash
}

func BytesToKeccak(data []byte) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(data),
	)

	return hash
}

func CellCreation(assetID int, txNumber int) []byte {

	a := IntToKeccak(txNumber)

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
