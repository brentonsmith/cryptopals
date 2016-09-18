package cryptopals

import "testing"

func TestChallenge1(t *testing.T) {
	hexToEncode := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual := Challenge1(hexToEncode)

	if expected != actual {
		t.Errorf("Expected: %s \n Actual: %s", expected, actual)
	}
}

func TestChallenge2(t *testing.T) {
	in1 := "1c0111001f010100061a024b53535009181c"
	in2 := "686974207468652062756c6c277320657965"

	expected := "746865206b696420646f6e277420706c6179"
	actual, err := Challenge2(in1, in2)

	if err != nil {
		t.Error("Expected no error!")
	}

	if expected != actual {
		t.Errorf("Expected: %s \n Actual: %s", expected, actual)
	}
}

func TestChallenge2_HandlesUnevenInputs(t *testing.T) {
	in1 := "1c0"
	in2 := "68697"

	_, err1 := Challenge2(in1, in2)

	if err1 == nil {
		t.Error("Expected an error!")
	}

	_, err2 := Challenge2(in2, in1)

	if err2 == nil {
		t.Error("Expected an error!")
	}
}
