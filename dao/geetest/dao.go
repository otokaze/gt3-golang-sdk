package geetest

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"gt3-golang-sdk/conf"
	"gt3-golang-sdk/http"
	"gt3-golang-sdk/model/geetest"
)

const (
	_register = "/register.php"
	_validate = "/validate.php"
)

// Dao is account dao.
type Dao struct {
	c *conf.Config
	// url
	registerURI string
	validateURI string
	// http client
	client *http.Client
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:           c,
		registerURI: c.Host.Geetest + _register,
		validateURI: c.Host.Geetest + _validate,
		// http client
		client: http.NewClient(c.HTTPClient),
	}
	return
}

// PreProcess preprocessing the geetest and get to challenge
func (d *Dao) PreProcess(mid int64, ip, clientType string, newCaptcha int) (challenge string, err error) {
	var (
		bs     []byte
		params url.Values
	)
	params = url.Values{}
	params.Set("user_id", strconv.FormatInt(mid, 10))
	params.Set("new_captcha", strconv.Itoa(newCaptcha))
	params.Set("client_type", clientType)
	params.Set("ip_address", ip)
	params.Set("gt", d.c.Secret.CaptchaID)
	if bs, err = d.client.Get(d.registerURI, params); err != nil {
		return
	}
	if len(bs) != 32 {
		return
	}
	challenge = string(bs)
	return
}

// Validate recheck the challenge code and get to seccode
func (d *Dao) Validate(challenge, seccode, clientType, ip, captchaID string, mid int64) (res *geetest.ValidateRes, err error) {
	var (
		bs     []byte
		params url.Values
	)
	params = url.Values{}
	params.Set("seccode", seccode)
	params.Set("challenge", challenge)
	params.Set("captchaid", captchaID)
	params.Set("client_type", clientType)
	params.Set("ip_address", ip)
	params.Set("json_format", "1")
	params.Set("sdk", "golang_3.0.0")
	params.Set("user_id", strconv.FormatInt(mid, 10))
	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	if bs, err = d.client.Post(d.validateURI, params); err != nil {
		return
	}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	return
}
