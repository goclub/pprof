package xpprof

import (
	"net/http"
	"net/http/pprof"
)

func RegisterHttp(handleFunc func(path string, handler func(w http.ResponseWriter, r *http.Request))) {
	handleFunc("/debug/", pprof.Index)
	handleFunc("/debug/cmdline", pprof.Cmdline)
	handleFunc("/debug/profile", pprof.Profile)
	handleFunc("/debug/symbol", pprof.Symbol)
	handleFunc("/debug/trace", pprof.Trace)
	handleFunc("/debug/allocs", wrapHandler(pprof.Handler("allocs")))
	handleFunc("/debug/block", wrapHandler(pprof.Handler("block")))
	handleFunc("/debug/goroutine", wrapHandler(pprof.Handler("goroutine")))
	handleFunc("/debug/heap", wrapHandler(pprof.Handler("heap")))
	handleFunc("/debug/mutex", wrapHandler(pprof.Handler("mutex")))
	handleFunc("/debug/threadcreate", wrapHandler(pprof.Handler("threadcreate")))
}
func wrapHandler(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
