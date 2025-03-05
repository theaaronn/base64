package base64

import (
	"testing"
)

var base64Alphabet = [64]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
}
var base64AlphabetStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var base64AlphabetMap = map[int]byte{
	0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G', 7: 'H', 8: 'I', 9: 'J', 10: 'K', 11: 'L', 12: 'M', 13: 'N', 14: 'O', 15: 'P',
	16: 'Q', 17: 'R', 18: 'S', 19: 'T', 20: 'U', 21: 'V', 22: 'W', 23: 'X', 24: 'Y', 25: 'Z', 26: 'a', 27: 'b', 28: 'c', 29: 'd', 30: 'e', 31: 'f',
	32: 'g', 33: 'h', 34: 'i', 35: 'j', 36: 'k', 37: 'l', 38: 'm', 39: 'n', 40: 'o', 41: 'p', 42: 'q', 43: 'r', 44: 's', 45: 't', 46: 'u', 47: 'v',
	48: 'w', 49: 'x', 50: 'y', 51: 'z', 52: '0', 53: '1', 54: '2', 55: '3', 56: '4', 57: '5', 58: '6', 59: '7', 60: '8', 61: '9', 62: '+', 63: '/',
}

func stringLookup(char byte) int {
	for i := 0; i < 64; i++ {
		if base64Alphabet[i] == char {
			return i
		}
	}
	return -1
}

func arrayLookup(char byte) int {
	for i := 0; i < 64; i++ {
		if base64AlphabetStr[i] == char {
			return i
		}
	}
	return -1
}

func mapLookup(char byte) int {
	for k, v := range base64AlphabetMap {
		if v == char {
			return k
		}
	}
	return -1
}

var (
	idx = 0
	mapp = 0
	arr  = 0
)

// This one seems to be the fastest, followed by Arr
func BenchmarkStr(b *testing.B) {
	for char := range base64Alphabet {
		idx = stringLookup(base64AlphabetStr[char])
	}
}
func BenchmarkMap(b *testing.B) {
	for char := range base64Alphabet {
		mapp = mapLookup(base64AlphabetStr[char])
	}
}
func BenchmarkArr(b *testing.B) {
	for char := range base64Alphabet {
		arr = arrayLookup(base64AlphabetStr[char])
	}
}
