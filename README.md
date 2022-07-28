# pprof

> [goclub/http](https://github.com/goclub/http) 的 pprof 分析工具

```go
package main

import xhttp "github.com/goclub/http"
import xpprof "github.com/goclub/pprof"
import log "log"
import http "net/http"

func main () {
    r := xhttp.NewRouter(xhttp.RouterOption{})
    server := &http.Server{
        Addr: ":5241",
        Handler: r,
    }
    xpprof.Register(r.NetHttpHandleFunc)
	log.Print("http://localhost" + server.Addr + "/debug/")
	log.Printf("%+v", server.ListenAndServe())
}
```