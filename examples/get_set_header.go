package main

import "github.com/zanjie1999/httpme"

func main() {

	req := httpme.Httpme()

	resp, _ := req.Get("http://go.xiulian.net.cn", httpme.Header{"Referer": "http://www.jeapedu.com"})

	println(resp.Text())

}
