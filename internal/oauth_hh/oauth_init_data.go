package oauth_hh

import "strings"

type OAuthInitData struct {
	addrAuth     string
	redirectUri  string
	clientId     string
	clientSecret string
	sessionId    string
}

func (oa *OAuthInitData) GetAuthLink(fn transform) string {

	if fn != nil {
		return fn(oa)
	}

	var replacer *strings.Replacer

	if oa.sessionId == "" {
		replacer = strings.NewReplacer("{client_id}", oa.clientId, "state={state}&", "", "{redirect_uri}", oa.redirectUri)
	} else {
		replacer = strings.NewReplacer("{client_id}", oa.clientId, "{state}", oa.sessionId, "{redirect_uri}", oa.redirectUri)
	}

	return replacer.Replace(oa.addrAuth)
}

func NewOAuthInitData(addrAuth string, redirectUri string, clientId string, clientSecret string, sessionId string) IOAuthInitData {
	return &OAuthInitData{
		addrAuth,
		redirectUri,
		clientId,
		clientSecret,
		sessionId,
	}
}
