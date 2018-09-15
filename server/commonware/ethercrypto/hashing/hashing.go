package hashing

import (
	"github.com/miguelmota/go-solidity-sha3"
	"encoding/hex"
	"strconv"
)

//StringToKeccak converts string to Kecckak hash of Ethereum
func StringToKeccak(data string) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String(data),
	)

	return hash
}

//IntToKeccak converts int to Kecckak hash of Ethereum
func IntToKeccak(data int64) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.String(data),
	)
	return hash
}

//BytesToKeccak converts byte to Kecckak hash of Ethereum
func BytesToKeccak(data []byte) []byte {

	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(data),
	)
	return hash
}

//CellCreation create a specific cell
func CellCreation(assetID string, txNumber int64) []byte {

	a := solsha3.SoliditySHA3(
		solsha3.String(strconv.Itoa(int(txNumber))),
	)
	b := solsha3.SoliditySHA3(
		solsha3.String(assetID),
	)
	c := solsha3.SoliditySHA3(
		solsha3.String(hex.EncodeToString(a) + hex.EncodeToString(b)),
	)

	return c
}

