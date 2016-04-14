# ptlogin

[![Build Status](https://travis-ci.org/chucklqsun/ptlogin.svg?branch=master)](https://travis-ci.org/chucklqsun/ptlogin)

Use to login account.  
All algorithm and process are open and extract from public JS  

Usage:  
```go
package main

import (
	"fmt"
	"github.com/chucklqsun/ptlogin"
)

var _ = fmt.Println

func main() {
	var pt ptlogin.Ptlogin  
	//username,password  
	pt.SetInput("1122233344", "aaabbbccdd")  
	pt.SetCookieName("temp.cookie")  
	pt.Ptui_checkVC()  
}
```
