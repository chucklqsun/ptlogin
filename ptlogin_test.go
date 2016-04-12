package ptlogin

import (
	"testing"
)

func TestSetInput(t *testing.T) {
	var pt Ptlogin

	expect := map[string]string{
		"username": "123456",
		"password": "abcdef",
	}

	pt.SetInput(expect["username"], expect["password"])

	if pt.uInput != expect["username"] {
		t.Errorf("SetInput username: expected %s, got %s", expect["username"], pt.uInput)
	}

	if pt.pInput != expect["password"] {
		t.Errorf("SetInput password: expected %s, got %s", expect["password"], pt.pInput)
	}
}

func TestSetCookieName(t *testing.T) {
	var pt Ptlogin
	expect := "HelloPT"
	pt.SetCookieName(expect)

	if pt.cookieName != expect {
		t.Errorf("SetCookieName: expected %s, got %s", expect, pt.cookieName)
	}
	if pt.cookieArr == nil {
		t.Error("SetCookieName cookieArr is nil ptr")
	}
}

func TestMd5(t *testing.T) {
	var pt Ptlogin
	expect := map[string]string{
		"empty": "D41D8CD98F00B204E9800998ECF8427E",
		"ab":    "187EF4436122D1CC2F40DC2B92F0EBA0",
		"abc":   "900150983CD24FB0D6963F7D28E17F72",
	}
	got := pt.md5([]byte("ab"), []byte{})
	if got != expect["ab"] {
		t.Errorf("Md5 1: expected %s, got %s", expect["ab"], got)
	}
	got = pt.md5([]byte("ab"), []byte("c"))
	if got != expect["abc"] {
		t.Errorf("Md5 2: expected %s, got %s", expect["abc"], got)
	}

	got = pt.md5([]byte{}, []byte{})
	if got != expect["empty"] {
		t.Errorf("Md5 3: expected %s, got %s", expect["empty"], got)
	}

	got = pt.md5([]byte{}, []byte("abc"))
	if got != expect["abc"] {
		t.Errorf("Md5 4: expected %s, got %s", expect["abc"], got)
	}
}
