package routes

import (
	handler "github.com/arthurkay/lucid-framework/handlers"
	"github.com/arthurkay/lucid-framework/middleware"

	"net/http/pprof"

	"github.com/gorilla/mux"
)

// Api returns a router instance
func Api() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.ApiHomeHandler).Methods("POST", "GET")
	r.HandleFunc("/login", handler.ApiLoginHandler).Methods("POST")
	r.HandleFunc("/register", handler.ApiRegisterHandler).Methods("POST")
	r.HandleFunc("/reset-password", handler.ApiResetPasswordHandler).Methods("POST")
	r.HandleFunc("/new-password", handler.ApiNewPasswordHandler).Methods("POST")
	auth := r.NewRoute().Subrouter()
	auth.Use(middleware.CheckAuth)
	auth.HandleFunc("/service", handler.ApiServiceRequest).Methods("POST")

	// Add the pprof routes
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
