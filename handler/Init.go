package handler

import (
	"os"

	"gorm.io/gorm"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


type DatabaseCollections struct {
	MySqlDB *gorm.DB
}


// var (
// 	googleOauthConfig *oauth2.Config
// 	// TODO: randomize it
// 	oauthStateString = "pseudo-random"
// )

// func init() {
// 	googleOauthConfig = &oauth2.Config{
// 		RedirectURL:  os.Getenv("SERVER_IP"),
// 		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
// 		Endpoint:     google.Endpoint,
// 	}
// }

type Config struct {
	GoogleLoginConfig   oauth2.Config
	// FacebookLoginConfig oauth2.Config
}

var AppConfig Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
// const OauthFacebookUrlAPI = "https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token="

func LoadConfig() {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://"+os.Getenv("SERVER_IP")+"/api/googlelogin/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	
}
