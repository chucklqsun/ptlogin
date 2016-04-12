package ptlogin

import (
	"bufio"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var _, _ = url.Parse("")

type pt_request struct {
	account    string
	url        string
	proxy      string
	method     string
	head       map[string]string
	body       string
	query      string
	cookieName string
	cookie     string
	result     interface{}
}

func (pt *Ptlogin) check() {
	req := pt_request{
		url:    "http://check.ptlogin2.qq.com/check?",
		proxy:  "",
		method: "GET",
		body:   "",
		query:  fmt.Sprintf("pt_tea=1&uin=%s&appid=21000115&ptlang=2052&r=%f", pt.uInput, rand.Float32()),
		cookie: "",
	}

	if pt.sendRequest(&req) {
		callback := string(req.result.([]byte))
		reg := regexp.MustCompile(`(?U)'(.*)'`)
		params := reg.FindAllString(callback, -1)

		pt.m_type = strings.Trim(params[0], `'`)
		pt.vcode = strings.Trim(params[1], `'`)
		params[2] = strings.Replace(params[2], `\x`, ``, -1)
		pt.salt, _ = hex.DecodeString(strings.Trim(params[2], `'`))

		pt.verifysession = strings.Trim(params[3], `'`)
		pt.isRandSalt = strings.Trim(params[4], `'`)

	}
	return
}

func (pt *Ptlogin) login() bool {
	var query = ""
	for key, value := range pt.submit_o {
		query += fmt.Sprintf("%s=%s&", key, value)
	}
	req := pt_request{
		url:    "http://ptlogin2.qq.com/login?",
		proxy:  "",
		method: "GET",
		body:   "",
		query:  query,
		cookie: "",
	}
	var callback = ""
	if pt.sendRequest(&req) {
		callback = string(req.result.([]byte))
	}

	reg := regexp.MustCompile(`(?U)'(.*)'`)
	params := reg.FindAllString(callback, -1)

	flag := strings.Trim(params[0], `'`)
	if flag != "0" {
		return false
	}
	req.url = strings.Trim(params[2], `'`)
	req.query = ""
	if pt.sendRequest(&req) {
		callback = string(req.result.([]byte))
	}
	//fmt.Println(callback)

	return true

}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (pt *Ptlogin) readCookie() {
	var infile, str string

	if pt.cookieName == "" {
		infile = "tmp_cookie"
	} else {
		infile = pt.cookieName
	}
	file, err := os.Open(infile)
	if err != nil {
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			return
		}
		str = string(line)
	}
	param := strings.Split(str, ";")
	for _, value := range param {
		f := strings.Split(value, "=")
		if len(f) == 2 {
			pt.cookieArr[strings.Trim(f[0], " ")] = strings.Trim(f[1], " ")
		} else {
			pt.cookieArr[strings.Trim(f[0], " ")] = ""
		}
	}

	return
}

func (pt *Ptlogin) writeCookie(input map[string]string) {
	var filename, str string
	var f *os.File
	if pt.cookieName == "" {
		filename = "tmp_cookie"
	} else {
		filename = pt.cookieName
	}

	if checkFileIsExist(filename) {
		pt.readCookie()
	}
	f, _ = os.Create(filename)
	for key, value := range input {
		pt.cookieArr[key] = value
	}

	for key, value := range pt.cookieArr {
		str += fmt.Sprintf("%s=%s;", key, value)
	}
	_, _ = io.WriteString(f, str)

}

func (pt *Ptlogin) sendRequest(ar *pt_request) bool {
	tr := &http.Transport{}

	//ignore https verify
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//proxy setup,left empty if no use
	if !strings.EqualFold(ar.proxy, "") {
		proxy, _ := url.Parse(ar.proxy)
		tr.Proxy = http.ProxyURL(proxy)
	}

	client := &http.Client{Transport: tr}

	//get vendor method
	method := ar.method
	var (
		req  *http.Request
		err  error
		body io.Reader //used for POST only
	)

	//setup request body
	switch {
	case method == "POST":
		body = strings.NewReader(ar.body)
	case method == "GET":
		ar.url += ar.query
		body = nil
	default:
		body = nil
	}
	req, err = http.NewRequest(method, ar.url, body)

	resp, err := client.Do(req)
	if err != nil {
		pt.log("Resp err:", err)
		return false
	}

	defer resp.Body.Close()

	cookieArr := resp.Header[http.CanonicalHeaderKey("Set-Cookie")]
	var cookieArrFilted map[string]string = make(map[string]string)
	for _, item := range cookieArr {
		param := strings.Split(item, ";")
		f := strings.Split(param[0], "=")
		cookieArrFilted[strings.Trim(f[0], " ")] = strings.Trim(f[1], " ")
	}
	pt.writeCookie(cookieArrFilted)

	respBody, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		pt.log("ReadAll err:", err)
		return false
	} else {
		ar.result = respBody
		return true
	}

	return false

}
