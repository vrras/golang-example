package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpFiras NoKtp = "8723978238234"
	var marriedStatus Married = true

	fmt.Println(noKtpFiras)
	fmt.Println(marriedStatus)
}
