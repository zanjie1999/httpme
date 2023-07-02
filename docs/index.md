# Installation


# example 1

```
package main

import "github.com/zanjie1999/httpme"


func main (){

        resp,err := httpme.Get("http://go.xiulian.net.cn")
        if err != nil {
          return 
        }
        println(resp.Text())
}

```
