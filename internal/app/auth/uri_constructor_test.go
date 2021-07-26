package auth

import (
	"fmt"
	"testing"
)

func TestConstructUri(t *testing.T) {
	params := uriParams{
		clientId:            "77e602fc63fa4b96acff255ed33428d3",
		responseType:        "code",
		redirectUri:         "localhost:10100",
		codeChallengeMethod: "S256",
		codeChallenge:       "KADwyz1X~HIdcAG20lnXitK6k51xBP4pEMEZHmCneHD1JhrcHjE1P3yU_NjhBz4TdhV6acGo16PCd10xLwMJJ4uCutQZHw",
		state:               "e21392da45dbf4",
	}

	fmt.Println(constructURI(params))
}
