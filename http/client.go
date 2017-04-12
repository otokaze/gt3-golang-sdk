package http

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gt3-golang-sdk/conf"
)

// Client struct
type Client struct {
	conf   *conf.HTTPClient
	client *http.Client
}

// NewClient new a http client.
func NewClient(c *conf.HTTPClient) (client *Client) {
	client = &Client{}
	var (
		transport *http.Transport
		dialer    *net.Dialer
	)
	dialer = &net.Dialer{
		Timeout:   time.Duration(c.Dial * int64(time.Second)),
		KeepAlive: time.Duration(c.KeepAlive * int64(time.Second)),
	}
	transport = &http.Transport{
		DialContext:     dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.client = &http.Client{
		Transport: transport,
	}
	return
}

// NewRequest new a http request.
func NewRequest(method, uri string, params url.Values) (req *http.Request, err error) {
	if method == "GET" {
		req, err = http.NewRequest(method, uri+"?"+params.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, uri, strings.NewReader(params.Encode()))
	}
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return
}

// Do handler request
func (client *Client) Do(req *http.Request) (body []byte, err error) {
	var res *http.Response
	if res, err = client.client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusInternalServerError {
		err = errors.New("http status code 5xx")
		return
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}
	return
}

// Get client.Get send GET request
func (client *Client) Get(uri string, params url.Values) (body []byte, err error) {
	req, err := NewRequest("GET", uri, params)
	if err != nil {
		return
	}
	body, err = client.Do(req)
	return
}

// Post client.Get send POST request
func (client *Client) Post(uri string, params url.Values) (body []byte, err error) {
	req, err := NewRequest("POST", uri, params)
	if err != nil {
		return
	}
	body, err = client.Do(req)
	return
}
