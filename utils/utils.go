package utils

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/miguelmota/go-solidity-sha3"
	"strconv"
)

func ToKeccak(input interface{}) ([]byte, error) {
	switch v := input.(type) {
	case int:
		return keccak256([]byte(strconv.Itoa(v)))
	case int64:
		return keccak256([]byte(strconv.Itoa(int(v))))
	case string:
		return keccak256([]byte(v))
	case []byte:
		return keccak256(v)
	default:
		return nil, errors.New("Not supported type")
	}
}

func keccak256(d []byte) ([]byte, error) {
	var result []byte
	hash := sha3.NewKeccak256()
	hash.Write(d)
	result = hash.Sum(result)

	return result, nil
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
