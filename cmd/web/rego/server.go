package rego

import (
	"fmt"
	"net/http"
)

type HTMLServer struct {
	templates map[string]Page
}

func NewHTMLServer() *HTMLServer {
	return &HTMLServer{
		templates: make(map[string]Page),
	}
}

func (s *HTMLServer) RegisterTemplate(path string, template Page) {
	s.templates[path] = template
}

func (s *HTMLServer) ServeHTML(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	template, exists := s.templates[path]
	if !exists {
		// Use the 404 "Not Found" template if the path is not registered
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(NOT_FOUND_TEMPLATE))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(template.GetTemplate()))
}

func (s *HTMLServer) ServeStaticFile(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Path[len("/static/"):]
    http.ServeFile(w, r, STATIC_DIR+"/"+filename)
}

func (s *HTMLServer) RegisterHTML(path string, html string) {
    s.templates[path] = Page{HTML: html}
}


func (s *HTMLServer) StartRouter(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", s.ServeStaticFile)
	mux.HandleFunc("/", s.ServeHTML)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on port %d...\n", port)
	http.ListenAndServe(addr, mux)
}
