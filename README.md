# HTTP咩

A next generation HTTP client for Golang. Like Python httpx/requests.

~~次时代Golang HTTP客户端~~  

因为Go上没有像Python中`httpx`或是`requests`那样完美好用的HTTP客户端，并且目前现有的多个名为`requests`都有一些跨平台等的兼容性问题不好处理，并且已经多年未维护（可能已经停止维护），迫不得已来造轮子（因为我也要用的），实现类似Python库`httpx`简单易用的特性

可以向下兼容`github.com/asmcos/requests`方便升級，为了避免产生混淆，库名为`httpme`既http咩（罗马音），就不叫`requests`混入其中了

可以在Issues提出需求，会尽可能的实现

# Installation 安装

```
go get -u github.com/zanjie1999/httpme
```

# Use 使用 
## Get

``` go
package main

import (
	"fmt"

	"github.com/zanjie1999/httpme"
)

func main (){

        resp,err := httpme.Get("http://google.cn")
        if err != nil{
          return
        }
        // raw Body
        fmt.println(resp.Content())
        // string
        fmt.println(resp.Text())
        // save to file
        resp.SaveFile("file.html")
}
```

## Post

``` go
package main

import (
	"fmt"

	"github.com/zanjie1999/httpme"
)

func main (){

        data := httpme.Datas{
          "name":"httpme_post_test",
        }
        resp,_ := httpme.Post("https://www.httpbin.org/post",data)
        // raw Body
        fmt.println(resp.Content())
        // string
        fmt.println(resp.Text())
}

```

     Server return data...

``` json
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "name": "httpme_post_test"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "23",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "www.httpbin.org",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
  },
  "json": null,
  "origin": "114.242.34.110",
  "url": "https://www.httpbin.org/post"
}

```


## PostJson

``` go
package main

import "github.com/zanjie1999/httpme"


func main (){

        jsonStr := "{\"name\":\"httpme_post_test\"}"
        resp,_ := httpme.PostJson("https://www.httpbin.org/post",jsonStr)
        println(resp.Text())
}

```

     Server return data...

``` json
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "name": "httpme_post_test"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "23",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "www.httpbin.org",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
  },
  "json": null,
  "origin": "114.242.34.110",
  "url": "https://www.httpbin.org/post"
}

```

## On Android
Android必须设置一个dns才能正常解析域名，只需要在生命周期开始前设置一次
```
package main

import (
	"fmt"

	"github.com/zanjie1999/httpme"
)

func main (){
        httpme.SetDns("223.6.6.6:53")
        resp,err := httpme.Get("http://google.cn")
        if err != nil{
          return
        }
        // raw Body
        fmt.println(resp.Content())
        // string
        fmt.println(resp.Text())
        // save to file
        resp.SaveFile("file.html")
}
```
```

# Feature Support
  - Set headers
  - Set params
  - Multipart File Uploads
  - Sessions with Cookie Persistence
  - Proxy
  - Authentication
  - JSON
  - Chunked Httpme
  - Debug
  - SetTimeout


# Set header

### example 1

``` go
req := httpme.Httpme()

resp,err := req.Get("http://google.cn",httpme.Header{"Referer":"http://www.jeapedu.com"})
if (err == nil){
  println(resp.Text())
}
```

### example 2

``` go
req := httpme.Httpme()
req.Header.Set("accept-encoding", "gzip, deflate, br")
resp,_ := req.Get("http://google.cn",httpme.Header{"Referer":"http://www.jeapedu.com"})
println(resp.Text())

```

### example 3

``` go
h := httpme.Header{
  "Referer":         "http://www.jeapedu.com",
  "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
}
resp,_ := req.Get("http://wgoogle.cn",h)

h2 := httpme.Header{
  ...
  ...
}
h3,h4 ....
// two or more headers ...
resp,_ = req.Get("http://google.cn",h,h2,h3,h4)
```


# Set params

``` go
p := httpme.Params{
  "title": "The blog",
  "name":  "file",
  "id":    "12345",
}
resp,_ := req.Get("http://www.cpython.org", p)

```


# Auth

Test with the `correct` user information.

``` go
req := httpme.Httpme()
resp,_ := req.Get("https://api.github.com/user",httpme.Auth{"zanjie1999","password...."})
println(resp.Text())
```

github return

```
{"login":"zanjie1999","id":xxxxx,"node_id":"Mxxxxxxxxx==".....
```

# JSON

``` go
req := httpme.Httpme()
req.Header.Set("Content-Type","application/json")
resp,_ = req.Get("https://httpbin.org/json")

var json map[string]interface{}
resp.Json(&json)

for k,v := range json{
  fmt.Println(k,v)
}
```


# SetTimeout

```
req := Httpme()
req.Debug = 1

// 20 Second
req.SetTimeout(20)
req.Get("http://golang.org")
```

# Get Cookies

``` go
resp,_ = req.Get("https://www.httpbin.org")
coo := resp.Cookies()
// coo is [] *http.Cookies
println("********cookies*******")
for _, c:= range coo{
  fmt.Println(c.Name,c.Value)
}
```
