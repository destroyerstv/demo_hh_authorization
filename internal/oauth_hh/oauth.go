package oauth_hh

type TokenData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type transform func(oa *OAuthInitData) string

type IOAuthInitData interface {
	GetAuthLink(fn transform) string
}

type IOAuthPersistData interface {
	GetUserId() string
	GetAccessToken() (string, error)
	GetJSON() string
}

type IOAuthTempData interface {
	GetAccessRefreshTkn() IOAuthPersistData
}
