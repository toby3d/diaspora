package diaspora

import (
	"fmt"

	log "github.com/kirillDanshin/dlog"
	http "github.com/valyala/fasthttp"
)

func (c *Client) get(url string) ([]byte, error) {
	return c.request(nil, url, "GET")
}

func (c *Client) post(dst []byte, url string) ([]byte, error) {
	return c.request(dst, url, "POST")
}

func (c *Client) request(dst []byte, url, method string) ([]byte, error) {
	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.SetMethod(method)
	req.Header.SetRequestURI(url)
	if c.TokenType != "" && c.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprint(c.TokenType, " ", c.AccessToken))
	}
	req.SetBody(dst)

	log.Ln("Request:")
	log.D(req)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	err := http.Do(req, resp)
	log.Ln("Resp:")
	log.D(resp)
	return resp.Body(), err
}
