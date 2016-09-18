package cryptopals

import "testing"

func TestChallenge1(t *testing.T) {
	hexToEncode := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual := Challenge1(hexToEncode)

	if actual != expected {
		t.Errorf("Expected: %s \n Actual: %s", expected, actual)
	}
}
