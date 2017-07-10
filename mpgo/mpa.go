package mpgo

import (
	"encoding/json"
	"errors"
	// "log"
	"net/http"
	"strings"
	// "net/url"
)

type ClientMP struct { 
	Access_token 	string 	`json:"access_token"`
	Refresh_token 	string 	`json:"refresh_token"`
	Live_mode 		bool 	`json:"live_mode"`
	User_id 		int64 	`json:"user_id"`
	Token_type 		string 	`json:"token_type"`
	Expires_in 		int 	`json:"expires_in"`
	Scope 			string 	`json:"scope"`
}

type CauseE struct {
	Code 		string 	`json:"code,omitempty"`
	Description string 	`json:"description,omitempty"`
}

const (
	urlBase string = "https://api.mercadopago.com/" 
)

func GetTokenMP(client_id, client_secret string, sandbox bool) (ClientMP, error) {

	var clientMP ClientMP

	if client_id == "" {
		return clientMP, errors.New("Client ID cannot be empty")
	} else if client_secret == "" {
		return clientMP, errors.New("Client secret cannot be empty")
	}

	url := urlBase + "oauth/token?grant_type=client_credentials&client_id=" + client_id + "&client_secret=" + client_secret

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return clientMP, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return clientMP, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&clientMP); err != nil {
		return clientMP, errors.New("Invalid credentials")
	}

	if sandbox {
		clientMP.Access_token = strings.Replace(clientMP.Access_token, "APP_USR", "TEST", 1)
	}

	return clientMP, nil
}

