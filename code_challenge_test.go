package main

import (
	"fmt"
	"testing"
)

func TestHashRunes(t *testing.T) {

}

func TestMakeChallenger(t *testing.T) {
	in := []rune("abcsdbcjksndsjnsdf")
	out := string(makeCodeChallenge(in))

	fmt.Println(len(in), len(out))

	fmt.Println(string(makeCodeChallenge(in)))
}
