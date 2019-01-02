package utils

import (
	"encoding/hex"
	"testing"
)

func TestToKeccak(t *testing.T) {
	t.Log("To Keccak256 test")

	a, err := ToKeccak("test")
	if err != nil {
		t.Fatal("Wrong string hashing")
	}

	if  hex.EncodeToString(a) != "9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658" {
		t.Fatal("Wrong string hashing")
	}
}
