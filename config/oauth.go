package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetOauthConfigGmail() oauth2.Config {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return *oauthConfig
}

func GetOauthConfigApple() oauth2.Config {
	oauthConfig := &oauth2.Config{
		ClientID:    os.Getenv("CLIENT_ID_APPLE"),
		RedirectURL: os.Getenv("REDIRECT"),
		Scopes: []string{
			"email", "name",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("APPLE_AUTH_URL"),
			TokenURL: os.Getenv("APPLE_TOKEN_URL"),
		},
	}
	return *oauthConfig
}
