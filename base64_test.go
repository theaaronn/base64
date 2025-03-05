package base64

import (
	"fmt"
	"testing"
)

var expectations = map[string]string{
	"":                "",
	" ":               "IA==",
	"=":               "PQ==",
	"A":               "QQ==",
	"XO":              "WE8=",
	"mex":             "bWV4",
	"water":           "d2F0ZXI=",
	"abcdef":          "YWJjZGVm",
	"cowshed":         "Y293c2hlZA==",
	"deez nuts":       "ZGVleiBudXRz",
	"not my problem!": "bm90IG15IHByb2JsZW0h",
	"is not like I want to break the encoding": "aXMgbm90IGxpa2UgSSB3YW50IHRvIGJyZWFrIHRoZSBlbmNvZGluZw==",
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

func Test64Encode(t *testing.T) {
	for input, encoded := range expectations {
		outcome, err := Encode(input)
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

func Test64Decode(t *testing.T) {
	for decoded, encoded := range expectations {
		outcome, err := Decode(encoded)
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
