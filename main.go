package main

import "base32and64/base64"

func main() {
	secret := "THS"
	encodedSecret, err := base64.Encode64(secret)
	if err != nil {
		println(err.Error())
	}else {
		println(encodedSecret)
	}
	
	result, err := base64.Decode64(encodedSecret)
	if err != nil {
		println(err.Error())
	}else {
		println(result)
	}

}
