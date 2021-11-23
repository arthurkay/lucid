package routes

import (
	"net/http/pprof"

	"github.com/arthurkay/lucid/middleware"
	"github.com/gorilla/mux"
)

// ApiRoutes return a router instance, meant for api
func ApiRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.DefaultApiHeaders)
	r.StrictSlash(true)
	return r
}

// WebRoutes returns a router instance, meant for web
func WebRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.DefaultWebHeaders)
	r.StrictSlash(true)

	// Perfomance measuring metrics
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	r.Handle("/debug/pprof/block", pprof.Handler("block"))
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	return r
}
