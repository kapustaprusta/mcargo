package main

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
	fmt.Println(string(makeCodeVerifier()))
}
