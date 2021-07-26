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

type UriParams struct {
	ClientId            string
	ResponseType        string
	RedirectUri         string
	CodeChallengeMethod string
	CodeChallenge       string
	State               string
}

func ConstructURI(params UriParams) string {
	clientIdParam := fmt.Sprintf(clientIdTmpl, params.ClientId)
	responseTypeParam := fmt.Sprintf(responseTypeTmpl, params.ResponseType)
	redirectUriParam := fmt.Sprintf(redirectUriTmpl, params.RedirectUri)
	codeChallengeMethodParam := fmt.Sprintf(codeChallengeMethodTmpl, params.CodeChallengeMethod)
	codeChallengeParam := fmt.Sprintf(codeChallengeTmpl, params.CodeChallenge)
	stateParam := fmt.Sprintf(stateTmpl, params.State)

	return baseUri + "?" + clientIdParam +
		"&" + responseTypeParam +
		"&" + redirectUriParam +
		"&" + codeChallengeMethodParam +
		"&" + codeChallengeParam +
		"&" + stateParam
}
