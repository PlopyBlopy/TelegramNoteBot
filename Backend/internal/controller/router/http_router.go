package router

import (
	"fmt"
	"net/http"
)

type ServeMux struct {
	httpServeMux *http.ServeMux
}

func NewServeMux() *ServeMux {
	sm := http.NewServeMux()

	sm.HandleFunc("GET /live", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	sm.HandleFunc("GET /ready", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	return &ServeMux{
		httpServeMux: sm,
	}
}

func (sm *ServeMux) InitRouter(apiVer int) http.Handler {
	r := http.NewServeMux()

	pattern := fmt.Sprintf("/api/v%d/", apiVer)
	prefix := fmt.Sprintf("/api/v%d", apiVer)

	r.Handle(pattern, http.StripPrefix(prefix, sm.httpServeMux))

	return r
}

func (sm *ServeMux) AddHandler(pattern string, handler http.HandlerFunc) {
	sm.httpServeMux.HandleFunc(pattern, handler)
}
