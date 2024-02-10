package auth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

type Session struct {
	Session struct {
		Name       string `json:"name"`
		Key        string `json:"key"`
		Subscriber int    `json:"subscriber"`
	} `json:"session"`
}

func GetToken(apiKey string, res *Token) error {
	resp, err := http.Get(fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=auth.gettoken&api_key=%s&format=json", apiKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetchRequestToken: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}
	return nil
}

func AuthToken(apiKey, requestToken string) {
	authorizationURL := fmt.Sprintf("https://www.last.fm/api/auth/?api_key=%s&token=%s", apiKey, requestToken)

	fmt.Println("Please visit the following URL to authorize the application:")
	fmt.Println(authorizationURL)
	fmt.Println("Press Enter after authorization")
	fmt.Scanln()
}

func GetSessionKey(apiKey, requestToken, secretKey string, res *Session) error {
	methodSig := "api_key" + apiKey + "methodauth.getSession" + "token" + requestToken + secretKey
	apiSig := getMD5Hash(methodSig)

	resp, err := http.Get(fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=auth.getSession&api_key=%s&token=%s&api_sig=%s&format=json", apiKey, requestToken, apiSig))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetchSessionKey: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}

	return nil
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
