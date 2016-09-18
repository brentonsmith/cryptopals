package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func Challenge1(h string) string {
	hexStringAsBytes, _ := hex.DecodeString(h)
	return base64.StdEncoding.EncodeToString(hexStringAsBytes)
}

func Challenge2(i1, i2 string) (string, error) {
	if len(i1) != len(i2) {
		return "", errors.New("Length mismatch!")
	}

	length := hex.DecodedLen(len(i1))
	bytes1 := make([]byte, length)
	bytes2 := make([]byte, length)
	output := make([]byte, length)

	hex.Decode(bytes1, []byte(i1))
	hex.Decode(bytes2, []byte(i2))

	for i := 0; i < len(bytes1); i++ {
		output[i] = bytes1[i] ^ bytes2[i]
	}

	return hex.EncodeToString(output), nil
}
