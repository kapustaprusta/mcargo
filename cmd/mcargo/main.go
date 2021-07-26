package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kapustaprusta/mcargo/internal/app/auth"
)

var (
	url = "localhost:10100"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestUri := r.RequestURI
	switch requestUri {
	case "/auth":
		codeChallenge := auth.MakeCodeChallenge(auth.MakeCodeVerifier())
		authUri := auth.ConstructURI(auth.UriParams{
			ClientId:            "97b7d3dc266245eea2c9a4df4cc01e45",
			ResponseType:        "code",
			RedirectUri:         "localhost:10100/auth_callback",
			CodeChallengeMethod: "S256",
			CodeChallenge:       string(codeChallenge),
			State:               "helloworld",
		})

		rawAuthUri, err := json.Marshal(authUri)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(authUri)

		w.Write(rawAuthUri)
	case "/auth_callback":
		fmt.Println(r.RequestURI)
	}

}

func main() {
	fmt.Printf("Listening at %s\n", url)
	http.ListenAndServe(url, &myHandler{})
}
