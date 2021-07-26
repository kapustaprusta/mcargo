package auth

import (
	"fmt"
	"testing"
)

func TestConstructUri(t *testing.T) {
	params := UriParams{
		ClientId:            "77e602fc63fa4b96acff255ed33428d3",
		ResponseType:        "code",
		RedirectUri:         "localhost:10100",
		CodeChallengeMethod: "S256",
		CodeChallenge:       "KADwyz1X~HIdcAG20lnXitK6k51xBP4pEMEZHmCneHD1JhrcHjE1P3yU_NjhBz4TdhV6acGo16PCd10xLwMJJ4uCutQZHw",
		State:               "e21392da45dbf4",
	}

	fmt.Println(ConstructURI(params))
}
