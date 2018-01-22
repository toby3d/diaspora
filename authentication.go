package diaspora

import "golang.org/x/oauth2"

type Client struct{ *oauth2.Token }

func NewClient(token *oauth2.Token) *Client {
	return &Client{token}
}
