package auth

import (
	"fmt"
	"testing"
)

func TestHashRunes(t *testing.T) {

}

func TestMakeChallenger(t *testing.T) {
	in := []rune("abcsdbcjksndsjnsdf")
	out := string(MakeCodeChallenge(in))

	fmt.Println(len(in), len(out))

	fmt.Println(string(MakeCodeChallenge(in)))
}

func BenchmarkMakeChallenger(b *testing.B) {
	in := []rune("abcsdbcjksndsjnsdf")

	for i := 0; i < b.N; i++ {
		MakeCodeChallenge(in)
	}
}
