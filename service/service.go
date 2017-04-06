package service

import (
	"gt3-golang-sdk/conf"
	"gt3-golang-sdk/dao"
)

// Service Geetest
type Service struct {
	// config
	c *conf.Config
	// dao
	d *dao.Dao
}

// New new a service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		d: dao.New(c),
	}
	return s
}
