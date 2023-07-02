package main

import "github.com/zanjie1999/httpme"

func main() {
	resp, _ := httpme.Get("http://go.xiulian.net.cn")
	println(resp.Text())
}
