package main

import (
	oauth_hh "demo_hh_authorization/internal/oauth_hh"
	service "demo_hh_authorization/internal/srv"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	const configPath = "/home/test/go_work/src/demo_hh_authorization/config.yaml"

	viper.SetConfigFile(configPath)
	fmt.Println(viper.AllKeys())
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config; %s", err.Error())
	}

	hhOAuth := oauth_hh.NewOAuthInitData(
		viper.GetString("OAuthHH.Service"),
		viper.GetString("OAuthHH.RedirectURI"),
		viper.GetString("OAuthHH.ClientID"),
		viper.GetString("OAuthHH.ClientSecret"),
		"sessionId_string",
	)

	AuthLink := hhOAuth.GetAuthLink(nil)
	fmt.Println(AuthLink)

	exit := make(chan os.Signal, 1)
	s := service.NewService(
		viper.GetString("OAuthHH.ClientID"),
		viper.GetString("OAuthHH.ClientSecret"),
		viper.GetString("OAuthHH.RedirectURI"),
		viper.GetString("OAuthHH.AuthRefreshLink"),
	)
	go func() {
		err := s.Run("8080")
		if err != nil {
			fmt.Println(err)
			exit <- syscall.SIGTERM
		}
	}()

	// graceful shutdown
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}
