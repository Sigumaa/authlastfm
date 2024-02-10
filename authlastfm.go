// Package authlastfm provides utilities for authenticating
// with Last.fm's API.
package authlastfm

type Client struct {
	baseURL   string
	userName  string
	apiKey    string
	secretKey string
}

func New(userName, apiKey, secretKey string) *Client {
	return &Client{
		baseURL:   "https://ws.audioscrobbler.com/2.0/",
		userName:  userName,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}
