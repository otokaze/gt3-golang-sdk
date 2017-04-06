package dao

import (
	"gt3-golang-sdk/conf"
)

const (
	_register = "/register.php"
	_validate = "/validate.php"
)

type Dao struct {
	c *conf.Config
	// url
	regURI  string
	valiURI string
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:       c,
		regURI:  c.Host.Geetest + _register,
		valiURI: c.Host.Geetest + _validate,
	}
	return
}
