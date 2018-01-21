package diaspora

import (
	log "github.com/kirillDanshin/dlog"
	http "github.com/valyala/fasthttp"
)

func (c *Client) get(path string, args *http.Args) ([]byte, error) {
	return c.request("GET", nil, path, args)
}

func (c *Client) post(dst []byte, path string, args *http.Args) ([]byte, error) {
	return c.request("POST", dst, path, args)
}

func (c *Client) request(method string, dst []byte, path string, args *http.Args) ([]byte, error) {
	requestURL := c.Pod
	requestURL.Path += path
	if args != nil {
		requestURL.RawQuery = args.String()
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.SetMethod(method)
	req.Header.SetRequestURI(requestURL.String())
	req.Header.SetHost(requestURL.Host)
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
