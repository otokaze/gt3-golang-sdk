package service

import (
	"crypto/md5"
	"encoding/hex"
	"gt3-golang-sdk/conf"
	"gt3-golang-sdk/dao/geetest"
	mdl "gt3-golang-sdk/model/geetest"
	"math/rand"
	"strconv"
)

// Service Geetest
type Service struct {
	// config
	c *conf.Config
	// dao
	d *geetest.Dao
}

// New new a service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		d: geetest.New(c),
	}
	return s
}

// PreProcess preprocessing the geetest and get to challenge
func (s *Service) PreProcess(mid int64, ip, clientType string, newCaptcha int) (res *mdl.ProcessRes, err error) {
	var pre string
	res = &mdl.ProcessRes{}
	res.CaptchaID = s.c.Secret.CaptchaID
	res.NewCaptcha = newCaptcha
	if pre, err = s.d.PreProcess(mid, ip, clientType, newCaptcha); err != nil || pre == "" {
		randOne := md5.Sum([]byte(strconv.Itoa(rand.Intn(100))))
		randTwo := md5.Sum([]byte(strconv.Itoa(rand.Intn(100))))
		challenge := hex.EncodeToString(randOne[:]) + hex.EncodeToString(randTwo[:])[0:2]
		res.Challenge = challenge
		return
	}
	res.Success = 1
	slice := md5.Sum([]byte(pre + s.c.Secret.PrivateKey))
	res.Challenge = hex.EncodeToString(slice[:])
	return
}

// Validate recheck the challenge code and get to seccode
func (s *Service) Validate(challenge, validate, seccode, clientType, ip string, success int, mid int64) (stat bool) {
	if len(validate) != 32 {
		return
	}
	if success != 1 {
		slice := md5.Sum([]byte(challenge))
		stat = hex.EncodeToString(slice[:]) == validate
		return
	}
	slice := md5.Sum([]byte(s.c.Secret.PrivateKey + "geetest" + challenge))
	if hex.EncodeToString(slice[:]) != validate {
		return
	}
	res, err := s.d.Validate(challenge, seccode, clientType, ip, s.c.Secret.CaptchaID, mid)
	if err != nil {
		return
	}
	slice = md5.Sum([]byte(seccode))
	stat = hex.EncodeToString(slice[:]) == res.Seccode
	return
}
