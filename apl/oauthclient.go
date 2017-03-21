package apl

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// getOauth2HTTPClient gets an oauth2 http.Client for sling to use
func getOauth2HTTPClient() *http.Client {
	tokenPath := fmt.Sprintf("%srequest_token", conf.API)
	oauthConfig := clientcredentials.Config{
		ClientID:     conf.Svc_Username,
		ClientSecret: conf.Svc_Password,
		TokenURL:     tokenPath,
		Scopes:       []string{},
	}

	return oauthConfig.Client(oauth2.NoContext)
}
