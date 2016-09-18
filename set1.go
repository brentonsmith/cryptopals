package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func Challenge1(h string) string {
	hexStringAsBytes, _ := hex.DecodeString(h)
	return base64.StdEncoding.EncodeToString(hexStringAsBytes)
}