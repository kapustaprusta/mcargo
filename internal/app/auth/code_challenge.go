package auth

import (
	"crypto/sha256"
	"encoding/base64"
)

func hashRunes(in []rune) []byte {
	hashSum := sha256.Sum256([]byte(string(in)))

	return hashSum[:]
}

func MakeCodeChallenge(in []rune) []byte {
	inAsHashSum := hashRunes(in)
	inAsBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(inAsHashSum)))
	base64.StdEncoding.Encode(inAsBase64, inAsHashSum)

	return inAsBase64
}
