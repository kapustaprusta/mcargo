package auth

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

func BenchmarkMakeChallenger(b *testing.B) {
	in := []rune("abcsdbcjksndsjnsdf")

	for i := 0; i < b.N; i++ {
		makeCodeChallenge(in)
	}
}
