package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		codeChallenge := auth.makeCodeChallenge(auth.makeCodeVerifier())
		authUri := auth.constructURI(uriParams{
			clientId:            "97b7d3dc266245eea2c9a4df4cc01e45",
			responseType:        "code",
			redirectUri:         "localhost:10100/auth_callback",
			codeChallengeMethod: "S256",
			codeChallenge:       string(codeChallenge),
			state:               "helloworld",
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
