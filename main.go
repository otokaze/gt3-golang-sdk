package main

import (
	"gt3-golang-sdk/conf"
	"gt3-golang-sdk/route"
	"gt3-golang-sdk/service"
)

func main() {
	// conf init
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// service init
	serv := service.New(conf.Conf)
	// http init
	route.Init(conf.Conf, serv)
}
