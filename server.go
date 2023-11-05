package rego

import (
	"fmt"
	"net/http"
)

type PageHandler struct {
	Handler func(w http.ResponseWriter, r *http.Request)
	Page    Page
}

type RegoServer struct {
	Handlers      map[string]PageHandler
	StaticDir     string
	StaticURL     string
	StaticEnabled bool
}

func NewRegoServer() *RegoServer {
	return &RegoServer{
		Handlers:      make(map[string]PageHandler),
		StaticDir:     "static",
		StaticURL:     "/static/",
		StaticEnabled: false,
	}
}

func (r *RegoServer) SetStatic(enable bool, path, dir string) {
	r.StaticDir = dir
	r.StaticURL = path
	r.StaticEnabled = enable

	if enable {
		r.Handlers[r.StaticURL] = PageHandler{
			Handler: func(w http.ResponseWriter, r *http.Request) {
				http.StripPrefix(path, http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
			},
		}
	}
}

func (r *RegoServer) Handle(path string, template Page) {
	r.Handlers[path] = PageHandler{Page: template}
}
func (r *RegoServer) CustomHandle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Handlers[path] = PageHandler{Handler: handler}
}

func (r *RegoServer) HandleHTML(path, HTML string) {
	template := Page{}
	template.SetTemplate(HTML)
	r.Handlers[path] = PageHandler{Page: template}
}

func (R *RegoServer) ServeHTML(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	template, exists := R.Handlers[path]
	if !exists {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(NotFoundTemplate))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(template.Page.GetTemplate()))
}

func (r *RegoServer) RegisterStaticFile(dir string) http.Handler {
	return http.StripPrefix(r.StaticURL, http.FileServer(http.Dir(dir)))
}

func (r *RegoServer) StartServer(port int) {
	mux := http.NewServeMux()
	for path, pageHandler := range r.Handlers {
		if path != r.StaticURL {
			if pageHandler.Handler != nil {
				mux.HandleFunc(path, pageHandler.Handler)
			} else {
				mux.HandleFunc(path, r.ServeHTML)
			}
		}
	}

	if r.StaticEnabled {
		fs := http.FileServer(http.Dir(r.StaticDir))
		mux.Handle(r.StaticURL, http.StripPrefix(r.StaticURL, fs))
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server started on port %d...\n", port)
	http.ListenAndServe(addr, mux)
}

