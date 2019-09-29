package route

import (
	"net/http"

	"github.com/otokaze/gt3-golang-sdk/conf"
	"github.com/otokaze/gt3-golang-sdk/service"
)

var (
	config  *conf.Config
	servive *service.Service
)

// Init init route
func Init(c *conf.Config, s *service.Service) {
	config, servive = c, s
	// init static file handler
	fsh := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
	http.HandleFunc("/x/gt/preprocess", gtPreProcess)
	http.HandleFunc("/x/gt/validate", gtValidate)
	// listen port 2233
	if err := http.ListenAndServe(":2233", nil); err != nil {
		panic(err)
	}
}
