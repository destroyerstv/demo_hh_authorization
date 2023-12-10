package api_hh

import (
	oauth_hh "demo_hh_authorization/internal/oauth_hh"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IHHApi interface {
	Dictionaries() IDictionaries
	CreationAvailabilityResume() bool
	Resumes(resume IHHResume, dict IDictionaries) bool
}

type HHApi struct {
	client *http.Client
	oapd   oauth_hh.IOAuthPersistData
}

// Dictionaries implements IHHApi.
func (hh *HHApi) Dictionaries() IDictionaries {
	method_api := http.MethodGet
	url_api := "https://api.hh.ru/dictionaries"

	data := url.Values{}

	req, err := http.NewRequest(method_api, url_api, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %s", err.Error())
	}

	token, err := hh.oapd.GetAccessToken()
	if err != nil {
		log.Fatalf("Ошибка получения токена авторизации: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", "test-agent")

	resp, err := hh.client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Статус запроса выполнился с ошибкой: %s", resp.Status)
	}

	defer resp.Body.Close()

	var dict Dictionaries
	err = json.NewDecoder(resp.Body).Decode(&dict)

	if err != nil {
		log.Fatalf("Не могу прочитать тело json-ответа: %s", err.Error())
	}

	return &dict
}

// CreationAvailabilityResume implements IHHApi.
func (hh *HHApi) CreationAvailabilityResume() bool {

	method_api := http.MethodGet
	url_api := "https://api.hh.ru/resumes/creation_availability"

	data := url.Values{}

	req, err := http.NewRequest(method_api, url_api, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %s", err.Error())
	}

	token, err := hh.oapd.GetAccessToken()
	if err != nil {
		log.Fatalf("Ошибка получения токена авторизации: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", "test-agent")

	resp, err := hh.client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		// log.Fatalf("Статус запроса выполнился с ошибкой: %s", resp.Status)
		log.Printf("Статус запроса выполнился с ошибкой: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Не могу прочитать тело запроса: %s", err.Error())
	}

	fmt.Println(string(body))

	return true
}

func (hh *HHApi) Resumes(resume IHHResume, dict IDictionaries) bool {
	return true
}

func NewHHApi(oapd oauth_hh.IOAuthPersistData) IHHApi {
	client := http.Client{Timeout: 3 * time.Second}
	return &HHApi{
		&client,
		oapd,
	}
}
