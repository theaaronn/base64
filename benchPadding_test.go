package base64

import (
	"fmt"
	"testing"
)

var result string

// More optimal even tough makes one more allocation than format
func manualPad6(str string) string {
	strLen := len(str)
	finalStr := str
	for strLen != 6 {
		finalStr = "0" + finalStr
		strLen++
	}
	return finalStr
}

func formatPad6(str string) string {
	return fmt.Sprintf("%06s", str)
}

func BenchmarkManualPad6(b *testing.B) {
	var s string
	for b.Loop() {
		s = manualPad6("123")
	}
	result = s
}

func BenchmarkFormatPad6(b *testing.B) {
	var s string
	for b.Loop() {
		s = formatPad6("123")
	}
	result = s
}

func manualPad8(str string) string {
	strLen := len(str)
	finalStr := str
	for strLen != 8 {
		finalStr = "0" + finalStr
		strLen++
	}
	return finalStr
}

// More optimal in every aspect (memory, allocations and swiftness)
func formatPad8(str string) string {
	return fmt.Sprintf("%08s", str)
}

func BenchmarkManualPad8(b *testing.B) {
	var s string
	for b.Loop() {
		s = manualPad8("123")
	}
	result = s
}

func BenchmarkFormatPad8(b *testing.B) {
	var s string
	for b.Loop() {
		s = formatPad8("123")
	}
	result = s
}
