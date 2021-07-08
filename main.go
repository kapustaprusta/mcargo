package main

import (
	"fmt"
)

func main() {
	codeChallenge := makeCodeChallenge(makeCodeVerifier())

	fmt.Println(string(codeChallenge))
}
