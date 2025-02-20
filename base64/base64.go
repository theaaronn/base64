package base64

import (
	"fmt"
	"math"
	"strconv"
)

// Support other formats than only ascii?

var base64Map = map[int]string{
	0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H",
	8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P",
	16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X",
	24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d", 30: "e", 31: "f",
	32: "g", 33: "h", 34: "i", 35: "j", 36: "k", 37: "l", 38: "m", 39: "n",
	40: "o", 41: "p", 42: "q", 43: "r", 44: "s", 45: "t", 46: "u", 47: "v",
	48: "w", 49: "x", 50: "y", 51: "z", 52: "0", 53: "1", 54: "2", 55: "3",
	56: "4", 57: "5", 58: "6", 59: "7", 60: "8", 61: "9", 62: "+", 63: "/",
}

/*
* When converting to binary zeros at the left are deleted, which if missing, can corrupt the data
* Example: 001010 turns into 1010
* Thats this function reason to be
 */
func format6BitBinary(strToPad string) string {
	finalStr := strToPad
	strLen := len(strToPad)

	for strLen != 6 {
		finalStr = "0" + finalStr
		strLen = len(finalStr)
	}

	return finalStr
}
func pad6BitBinary(str *string, howMuch int) {
	for range howMuch {
		*str += "0"
	}
}
func StripPaddingBits(str string, howMuch int) string {
	return str[:len(str)-howMuch]
}
func AddPaddingEquals(str *string, howMuch int) {
	for range howMuch {
		*str += "="
	}
}

func mapKey(value string) int {
	for k, v := range base64Map {
		if v == value {
			return k
		}
	}
	return 0
}

func Encode64(initialString string) (string, error) {
	if len(initialString) == 0 {
		return "", nil
	}
	var (
		finalString, chunk, binString string
		paddingToAdd                  = 0
		index                         = 0
	)

	// If the module isn't 0 (then it can only be 1 or 2), substract 3 and get absolute value (make it positive) to get the padding corresponding to the length of the string
	module := len(initialString) % 3
	if module != 0 {
		paddingToAdd = int(math.Abs(float64((module) - 3)))
	}

	for _, char := range initialString {
		ascii := int(char)
		bin8 := strconv.FormatInt(int64(ascii), 2)
		if len(bin8) != 8 {
			// Fill with zeros at the right till eight characters
			bin8 = fmt.Sprintf("%08s", bin8)	
		}
		binString += bin8
	}
	pad6BitBinary(&binString, paddingToAdd*2)

	for range len(binString) / 6 {
		chunk = binString[index : index+6]
		index += 6
		decimalChunk, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			return "", err
		}
		finalString += base64Map[int(decimalChunk)]
	}
	AddPaddingEquals(&finalString, paddingToAdd)

	return finalString, nil
}

func Decode64(initialString string) (string, error) {
	if len(initialString) == 0 {
		return "", nil
	}
	var (
		mapIndexes     = make([]int, 0)
		binString      = ""
		index          = 0
		finalString    = ""
		paddingToStrip = 0
	)

	for _, char := range initialString {
		if char == '=' {
			paddingToStrip += 2
		} else {
			mapIndexes = append(mapIndexes, mapKey(string(char)))
		}
	}
	for _, index := range mapIndexes {
		bin6 := strconv.FormatInt(int64(index), 2)
		binString += format6BitBinary(bin6)
	}

	binString = StripPaddingBits(binString, paddingToStrip)

	for range len(binString) / 8 {
		bin8 := binString[index : index+8]
		index += 8
		decimalNum, err := strconv.ParseInt(bin8, 2, 64)
		if err != nil {
			return "", err
		}
		finalString += string(rune(decimalNum))
	}
	return finalString, nil
}
