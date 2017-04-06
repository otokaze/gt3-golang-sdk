package route

import (
	"net/http"

	"gt3-golang-sdk/conf"
	"gt3-golang-sdk/service"
)

var (
	config  *conf.Config
	servive *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	config = c
	servive = s
	http.HandlerFunc("/geetest/preproccess", gtPreProcess)
	http.HandleFunc("/geetest/revalidate", gtValidate)
	if err := http.ListenAndServe(":2233", nil); err != nil {
		panic(err)
	}
}
