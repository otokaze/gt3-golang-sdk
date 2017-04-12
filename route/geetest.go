package route

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func gtPreProcess(w http.ResponseWriter, req *http.Request) {
	ip := req.RemoteAddr
	mid := int64(2233)
	res := make(map[string]interface{})
	process, err := servive.PreProcess(mid, ip, "web", 1)
	if err != nil {
		res["code"] = -400
		res["data"] = err
	} else {
		res["code"] = 0
		res["data"] = *process
	}
	bs, _ := json.Marshal(res)
	w.Write(bs)
}

func gtValidate(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	res := make(map[string]interface{})
	challenge := req.Form.Get("geetest_challenge")
	validate := req.Form.Get("geetest_validate")
	seccode := req.Form.Get("geetest_seccode")
	success := req.Form.Get("geetest_success")
	successi, err := strconv.Atoi(success)
	ip := req.RemoteAddr
	mid := int64(2233)
	if err != nil {
		successi = 1
	}
	status := servive.Validate(challenge, validate, seccode, "web", ip, successi, mid)
	if !status {
		res["code"] = -400
		res["msg"] = "Failed"
	} else {
		res["code"] = 0
		res["msg"] = "Success"
	}
	bs, _ := json.Marshal(res)
	w.Write(bs)
}
