package auth

import (
	"math/rand"
	"time"
)

var (
	verifierMinLength = 43
	verifierMaxLength = 128

	allowedSymbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_.-~")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffleRunes(in []rune) []rune {
	out := make([]rune, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}

	rand.Shuffle(len(out), func(i int, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}

func getCodeVerifierLen() int {
	return verifierMinLength + rand.Intn(verifierMaxLength-verifierMinLength)
}

func MakeCodeVerifier() []rune {
	verifierLen := getCodeVerifierLen()
	verifier := make([]rune, verifierLen)
	shuffledAllowedSymbols := shuffleRunes(allowedSymbols)

	for i := 0; i < len(verifier); i++ {
		symbolIdx := rand.Intn(len(shuffledAllowedSymbols))
		verifier[i] = shuffledAllowedSymbols[symbolIdx]
	}

	return verifier
}
