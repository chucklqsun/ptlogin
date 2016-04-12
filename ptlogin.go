package ptlogin

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/robertkrimen/otto"
	"strings"
)

type Ptlogin struct {
	m_type        string
	uInput        string
	pInput        string
	vcode         string
	verifysession string
	salt          []byte
	isRandSalt    string
	submit_o      map[string]string
	cookieName    string
	cookieArr     map[string]string
}

var pt Ptlogin

func (pt *Ptlogin) SetInput(uin, pwd string) {
	pt.uInput = uin
	pt.pInput = pwd
}

func (pt *Ptlogin) SetCookieName(name string) {
	pt.cookieName = name
	pt.cookieArr = make(map[string]string)
}

func (pt *Ptlogin) Ptui_checkVC() {
	pt.check()
	pt.cb_checkVC()
}

func (pt Ptlogin) md5(b []byte, e []byte) string {
	hash := md5.New()
	for _, value := range e {
		b = append(b, value)
	}
	hash.Write(b)
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum(b)))
}

func (pt *Ptlogin) RSAencrypt(o string) string {
	vm := otto.New()
	vm.Set("plaint_text", o)
	_, err := vm.Run(algorithm["RSA"])
	if err != nil {
		pt.log("JS runtime error:", err)
	}
	if value, err := vm.Get("p"); err == nil {
		if value_string, err := value.ToString(); err == nil {
			return value_string
		} else {
			pt.log("get RSA string failed:", err)
		}
	} else {
		pt.log("get RSA result failed:", err)
	}
	return ""
}

func (pt *Ptlogin) TEAencrypt(key, plain_text string) string {
	//use Javascript engine to get TX modified TEA result
	vm := otto.New()
	vm.Set("key", key)
	vm.Set("plaint_text", plain_text)
	vm.Run(algorithm["TEA"])
	if value, err := vm.Get("result"); err == nil {
		if value_string, err := value.ToString(); err == nil {
			result := base64.StdEncoding.EncodeToString(pt.hexchar2bin(value_string))
			return result
		} else {
			pt.log("get TEA string failed:", err)
		}
	} else {
		pt.log("get TEA result failed:", err)
	}
	return ""
}

func (pt *Ptlogin) getEncryption(t string, e []byte, n string, i bool) string {
	var o, a string
	var r []byte
	if i {
		o = t
	} else {
		o = pt.md5([]byte(t), []byte{})
		r = pt.hexchar2bin(o)
		a = pt.md5(r, e)
	}

	pp := pt.RSAencrypt(o)
	s := fmt.Sprintf("%x", len(pp)/2)
	c := fmt.Sprintf("%x", []byte(strings.ToUpper(n)))
	u := fmt.Sprintf("%x", len(c)/2)
	for len(u) < 4 {
		u = "0" + u
	}

	for len(s) < 4 {
		s = "0" + s
	}

	l := pt.TEAencrypt(a, s+pp+fmt.Sprintf("%x", e)+u+c)

	l = strings.Replace(l, `/`, `-`, -1)
	l = strings.Replace(l, `+`, `*`, -1)
	l = strings.Replace(l, `=`, `_`, -1)

	return l

	return ""
}

func (pt Ptlogin) hexchar2bin(str string) []byte {
	arr := make([]byte, 0)
	for i := 0; i < len(str); i += 2 {
		bb, _ := hex.DecodeString(str[i : i+2])
		arr = append(arr, bb[0])
	}
	return arr
}

func (pt *Ptlogin) submit(t string) {
	pt.submit_o = map[string]string{
		"aid":                 "21000115",
		"daid":                "8",
		"device":              "2",
		"fp":                  "loginerroralert",
		"from_ui":             "1",
		"g":                   "1",
		"h":                   "1",
		"low_login_enable":    "1",
		"low_login_hour":      "720",
		"pt_3rd_aid":          "0",
		"pt_randsalt":         pt.isRandSalt,
		"pt_ttype":            "1",
		"pt_uistyle":          "9",
		"pt_vcode_v1":         "0",
		"pt_verifysession_v1": pt.verifysession,
		"ptlang":              "2052",
		"ptredirect":          "1",
		"u":                   pt.uInput,
		"u1":                  "http%3A%2F%2Fdaoju.qq.com",
		"verifycode":          pt.vcode,
	}
	r := false
	pt.submit_o["p"] = pt.getEncryption(pt.pInput, pt.salt, pt.vcode, r)
	pt.login()
}

func (pt *Ptlogin) cb_checkVC() {
	pt.submit(pt.vcode)
}

func (pt *Ptlogin) log(v ...interface{}) {
	fmt.Println(v)
	return
}
