package main

import "github.com/zanjie1999/httpme"

func main() {
	data := httpme.Datas{
		"name": "httpme_post_test",
	}
	resp, _ := httpme.Post("https://www.httpbin.org/post", data)
	println(resp.Text())
}
