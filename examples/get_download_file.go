package main

import (
	"fmt"

	"github.com/zanjie1999/httpme"
)

func main() {
	resp, _ := httpme.Get("https://music.163.com/song/media/outer/url?id=29019010")
	fmt.Println("data len: ", len(resp.Content()))
	resp.SaveFile("神サマといっしょ.mp3")
}
