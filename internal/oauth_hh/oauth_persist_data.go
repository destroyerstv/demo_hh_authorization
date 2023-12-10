package oauth_hh

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type OAuthPersistData struct {
	UserId       string
	AccessToken  string
	ExpiresIn    time.Time
	RefreshToken string
	RefreshLink  string
	client       *http.Client `json:"-"`
}

func (oap *OAuthPersistData) GetUserId() string {
	return oap.UserId
}

func (oap *OAuthPersistData) CheckAccessToken() bool {
	return time.Until(oap.ExpiresIn) > 0
}

func (oap *OAuthPersistData) RefreshAccessToken() bool {
	data := url.Values{}
	data.Add("grant_type", "refresh_token")
	data.Add("refresh_token", oap.RefreshToken)

	req, err := http.NewRequest(http.MethodPost, oap.RefreshLink, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %s", err.Error())
		return false
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := oap.client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %s", err.Error())
		return false
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Статус запроса выполнился с ошибкой: %s", resp.Status)
		return false
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Не могу прочитать тело запроса: %s", err.Error())
		return false
	}

	var tokenData TokenData
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		log.Fatalf("Не могу получить JSON: %s", err.Error())
		return false
	}

	oap.AccessToken = tokenData.AccessToken
	oap.RefreshToken = tokenData.RefreshToken
	oap.ExpiresIn = time.Now().Add(time.Second * time.Duration(tokenData.ExpiresIn))

	return true
}

func (oap *OAuthPersistData) GetAccessToken() (string, error) {

	if !oap.CheckAccessToken() {
		if !oap.RefreshAccessToken() {
			return "", errors.New("произошла ошибка при обновлении access_token")
		}
	}

	return oap.AccessToken, nil
}

func (oap *OAuthPersistData) GetJSON() string {
	data, err := json.Marshal(oap)
	if err != nil {
		log.Fatalf("Не могу получить JSON из OAuthPersonal: %s", err.Error())
		return ""
	}

	return string(data)
}

func OAuthPersonalFromJson(jsonData string) IOAuthPersistData {
	var oAuthPersonal OAuthPersistData

	err := json.Unmarshal([]byte(jsonData), &oAuthPersonal)

	if err != nil {
		log.Fatalf("Не могу получить OAuthPersonal из JSON: %s", err.Error())
		return nil
	}

	client := http.Client{Timeout: 3 * time.Second}
	oAuthPersonal.client = &client

	return &oAuthPersonal
}

func NewOAuthPersonal(userId string, tokenData TokenData, refreshLink string) IOAuthPersistData {

	client := http.Client{Timeout: 3 * time.Second}

	accessToken := tokenData.AccessToken
	refreshToken := tokenData.RefreshToken
	expiresIn := time.Now().Add(time.Second * time.Duration(tokenData.ExpiresIn))

	return &OAuthPersistData{
		userId,
		accessToken,
		expiresIn,
		refreshToken,
		refreshLink,
		&client,
	}
}
