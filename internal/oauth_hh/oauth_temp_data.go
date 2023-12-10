package oauth_hh

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type OAuthTempData struct {
	UserId          string
	Code            string
	redirectUri     string
	clientId        string
	clientSecret    string
	authRefreshLink string
}

func (oat *OAuthTempData) GetAccessRefreshTkn() IOAuthPersistData {
	client := http.Client{Timeout: 3 * time.Second}

	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", oat.clientId)
	data.Add("client_secret", oat.clientSecret)
	data.Add("redirect_uri", oat.redirectUri)
	data.Add("code", oat.Code)

	req, err := http.NewRequest(http.MethodPost, oat.authRefreshLink, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %s", err.Error())
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %s", err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Статус запроса выполнился с ошибкой: %s", resp.Status)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Не могу прочитать тело запроса: %s", err.Error())
		return nil
	}

	var tokenData TokenData
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		log.Fatalf("Не могу получить JSON: %s", err.Error())
		return nil
	}

	return NewOAuthPersonal(oat.UserId, tokenData, oat.authRefreshLink)
}

func NewOAuthTempData(code string, state string, clientId string, clientSecret string, redirectUri string, authRefreshLink string) IOAuthTempData {
	return &OAuthTempData{
		UserId:          state,
		Code:            code,
		redirectUri:     redirectUri,
		clientId:        clientId,
		clientSecret:    clientSecret,
		authRefreshLink: authRefreshLink,
	}
}
