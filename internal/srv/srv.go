package service

import (
	api_hh "demo_hh_authorization/internal/api_hh"
	oauth_hh "demo_hh_authorization/internal/oauth_hh"
	"net/http"
)

type IService interface {
	Run(port string) error
}

type Service struct {
	storage         map[string]string
	clientId        string
	clientSecret    string
	redirect_uri    string
	authRefreshLink string
}

func (s *Service) Run(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/hh", s.oAuth)
	mux.HandleFunc("/oauth/hh/create", s.resumCreate)

	return http.ListenAndServe(":"+port, mux)
}

func (s *Service) oAuth(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	tmpCode := oauth_hh.NewOAuthTempData(
		params.Get("code"),
		params.Get("state"),
		s.clientId,
		s.clientSecret,
		s.redirect_uri,
		s.authRefreshLink,
	)

	oAuthPers := tmpCode.GetAccessRefreshTkn()

	s.storage[oAuthPers.GetUserId()] = oAuthPers.GetJSON()
}

func (s *Service) resumCreate(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userId := params.Get("state")
	data := s.storage[userId]
	oAuthPers := oauth_hh.OAuthPersonalFromJson(data)
	HHApi := api_hh.NewHHApi(oAuthPers)
	dict := HHApi.Dictionaries()

	resume := api_hh.NewMinimalResume()
	HHApi.Resumes(resume, dict)

	// if err != nil {
	// 	log.Fatalf("Ошибка создания резюме: %s", err.Error())
	// }

	// PostResume(tkn)

}

func NewService(clientId string, clientSecret string, redirect_uri string, authRefreshLink string) IService {
	storage := make(map[string]string)
	return &Service{
		storage,
		clientId,
		clientSecret,
		redirect_uri,
		authRefreshLink,
	}
}
