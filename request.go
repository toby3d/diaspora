package diaspora

import (
	"errors"
	"fmt"
	"strings"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

const prefix = "/api/v1"

func (c *Client) get(url string) ([]byte, error) {
	return c.request(nil, url, "GET")
}

func (c *Client) post(dst []byte, url string) ([]byte, error) {
	return c.request(dst, url, "POST")
}

func (c *Client) patch(dst []byte, url string) ([]byte, error) {
	return c.request(dst, url, "PATCH")
}

func (c *Client) delete(url string) ([]byte, error) {
	return c.request(nil, url, "PATCH")
}

func (c *Client) request(dst []byte, url, method string) ([]byte, error) {
	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.SetMethod(method)
	req.Header.SetRequestURI(url)
	if c.Token.TokenType != "" && c.Token.AccessToken != "" {
		req.Header.Set(
			"Authorization",
			fmt.Sprint(strings.Title(c.Token.TokenType), " ", c.Token.AccessToken),
		)
	}
	req.SetBody(dst)

	log.Ln("Request:")
	log.D(req)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	err := http.Do(req, resp)
	log.Ln("Resp:")
	log.D(resp)
	if err != nil {
		return resp.Body(), err
	}

	if resp.StatusCode() < 200 && resp.StatusCode() >= 300 {
		var data Error
		if err = json.Unmarshal(resp.Body(), &data); err != nil {
			return resp.Body(), err
		}
		return resp.Body(), errors.New(data.Error)
	}

	return resp.Body(), err
}
