package main

import (
	"fmt"

	"github.com/zanjie1999/httpme"
)

func main() {
	req := httpme.Httpme()
	resp, _ := req.Get("https://api.github.com/user", httpme.Auth{"zanjie1999", "password...."})
	println(resp.Text())
	fmt.Println(resp.R.StatusCode)
	fmt.Println(resp.R.Header["Content-Type"])
}
