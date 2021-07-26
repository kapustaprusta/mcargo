package auth

import (
	"fmt"
	"testing"
)

func TestShuffleRunes(t *testing.T) {
	runes := []rune("abcdefghijklmnop")
	fmt.Println(string(shuffleRunes(runes)))
}

func TestGetVerifierLen(t *testing.T) {
	fmt.Println(getCodeVerifierLen())
}

func TestMakeVerifier(t *testing.T) {
	fmt.Println(string(MakeCodeVerifier()))
}

func BenchmarkMakeVerifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeCodeVerifier()
	}
}
