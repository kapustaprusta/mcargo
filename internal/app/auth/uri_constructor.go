package auth

import (
	"fmt"
)

var (
	baseUri                 = "https://accounts.spotify.com/authorize"
	clientIdTmpl            = "client_id=%s"
	responseTypeTmpl        = "response_type=%s"
	redirectUriTmpl         = "redirect_uri=%s"
	codeChallengeMethodTmpl = "code_challenge_method=%s"
	codeChallengeTmpl       = "code_challenge=%s"
	stateTmpl               = "state=%s"
)

type uriParams struct {
	clientId            string
	responseType        string
	redirectUri         string
	codeChallengeMethod string
	codeChallenge       string
	state               string
}

func constructURI(params uriParams) string {
	clientIdParam := fmt.Sprintf(clientIdTmpl, params.clientId)
	responseTypeParam := fmt.Sprintf(responseTypeTmpl, params.responseType)
	redirectUriParam := fmt.Sprintf(redirectUriTmpl, params.redirectUri)
	codeChallengeMethodParam := fmt.Sprintf(codeChallengeMethodTmpl, params.codeChallengeMethod)
	codeChallengeParam := fmt.Sprintf(codeChallengeTmpl, params.codeChallenge)
	stateParam := fmt.Sprintf(stateTmpl, params.state)

	return baseUri + "?" + clientIdParam +
		"&" + responseTypeParam +
		"&" + redirectUriParam +
		"&" + codeChallengeMethodParam +
		"&" + codeChallengeParam +
		"&" + stateParam
}
