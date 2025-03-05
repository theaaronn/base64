package base64

import (
	"encoding/base64"
	"testing"
)

var (
	str = ""
	str2 = [8]byte{}
)

type discardWriter struct{}

func (w discardWriter) Write(p []byte) (n int, err error) {
	for char := range len(p) {
		str2[char] = p[char]
	}
	return len(p), nil
}

func stdEncode() {
	input := "foobar"
	encoder := base64.NewEncoder(base64.RawStdEncoding, discardWriter{})
	encoder.Write([]byte(input))
	encoder.Close()
}

func BenchmarkEncodeStd(b *testing.B) {
	for range b.N {
		stdEncode()
	}
}

func ownEncode() {
	input := "foobar"
	str, _ = Encode(input)
}

func BenchmarkEncodeOwn(b *testing.B) {
	for range b.N {
		ownEncode()
	}
}
