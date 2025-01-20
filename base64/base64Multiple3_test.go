package base64

import (
	"fmt"
	"testing"
)

var expectationsMultiple3 = map[string]string{
	"mex":             "bWV4",
	"not my problem!": "bm90IG15IHByb2JsZW0h",
	"abcdef":          "YWJjZGVm",
	"deez nuts":       "ZGVleiBudXRz",
	"":                "",
}

func AssureChars(first, second string) error {
	if len(first) != len(second) {
		return fmt.Errorf("Mismatch of length: %s and %s", first, second)
	}

	for index := 0; index < len(first); index++ {
		if first[index] != second[index] {
			return fmt.Errorf("Mismatch of characters in index %d of %s and %s", index, first, second)
		}
	}
	return nil
}

func Test64EncodeMultiple3(t *testing.T) {
	for input, encoded := range expectationsMultiple3 {
		outcome, err := Encode64(input)
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		if outcome != encoded {
			t.Fatalf("Incorrect output from base64 encoding: %s, expected %s and got %s", input, encoded, outcome)
		}
		charsAssurement := AssureChars(outcome, encoded)
		if charsAssurement != nil {
			t.Fatalf("%s", charsAssurement.Error())
		}
	}
}

func Test64DecodeMultiple3(t *testing.T) {
	for decoded, encoded := range expectationsMultiple3 {
		outcome, err := Decode64(encoded)
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		if outcome != decoded {
			t.Fatalf("Incorrect output from base64 decoding: %s, expected '%s' and got '%s'", encoded, decoded, outcome)
		}
		charsAssurement := AssureChars(outcome, decoded)
		if charsAssurement != nil {
			t.Fatalf("%s", charsAssurement.Error())
		}
	}
}
