package diaspora

import "golang.org/x/oauth2"

type Client struct {
	Issuer string
	Token  *oauth2.Token
}

func NewClient(issuer string, token *oauth2.Token) *Client {
	return &Client{issuer, token}
}
