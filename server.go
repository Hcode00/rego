package rego

import (
	"fmt"
	"net/http"
)

type PageHandler struct {
	Handler    func(w http.ResponseWriter, r *http.Request) // Handler function for the specific path.
	Concurrent bool                                         // A flag indicating whether to handle requests concurrently.
	Page       Page                                         // Page object associated with the path.
}

type RegoServer struct {
	Handlers               map[string]PageHandler // A map of registered page handlers for different paths.
	StaticDir              string                 // The directory path for static files.
	StaticURL              string                 // The URL path for static files.
	StaticEnabled          bool                   // A flag to enable or disable serving static files.
	StaticPagesConcurrency bool                   // A flag to determine whether static file handling should be concurrent.
}

// - NewRegoServer creates a new RegoServer instance with default values.
// - StaticDir:              "static",
// - StaticURL:              "/static/",
// - StaticEnabled:          false,
// - StaticPagesConcurrency: true,
func NewRegoServer() *RegoServer {
	return &RegoServer{
		Handlers:               make(map[string]PageHandler),
		StaticDir:              "static",
		StaticURL:              "/static/",
		StaticEnabled:          false,
		StaticPagesConcurrency: true,
	}
}

// SetStatic configures the static file serving settings.
// - enable: A flag indicating whether to enable static file serving.
// - path: The URL path for serving static files.
// - dir: The directory path where static files are located.
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

// SetConcurrencyForStatic enables or disables concurrent handling of static Pages.
func (r *RegoServer) SetConcurrencyForStatic(concurrent bool) {
	r.StaticPagesConcurrency = concurrent
}

// Handle registers a Page object for a specific path.
// - path: The URL path to associate with the Page template.
// - template: The Page object representing the HTML template.
func (r *RegoServer) Handle(path string, template Page) {
	r.Handlers[path] = PageHandler{Page: template}
}

// CustomHandle registers a custom handler function for a specific path.
// - path: The URL path to associate with the custom handler.
// - handler: The custom handler function for the specified path.
// - concurrent: A flag indicating whether to handle requests to this path concurrently.
func (r *RegoServer) CustomHandle(path string, handler func(w http.ResponseWriter, r *http.Request), concurrent bool) {
	r.Handlers[path] = PageHandler{
		Handler:    handler,
		Concurrent: concurrent,
	}
}

// HandleHTML registers an HTML string as a template for a specific path.
// - path: The URL path to associate with the HTML template.
// - HTML: The HTML content to be served for the specified path.
func (r *RegoServer) HandleHTML(path, HTML string) {
	template := Page{}
	template.SetTemplate(HTML)
	r.Handlers[path] = PageHandler{Page: template}
}

// ServeHTML is the default HTML serving handler.
// It serves the registered Page templates for specified paths.
// - w: The http.ResponseWriter for the HTTP response.
// - r: The http.Request for the incoming HTTP request.
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

// RegisterStaticFile creates an http.Handler to serve static files from a directory.
// - dir: The directory path where static files are located.
func (r *RegoServer) RegisterStaticFile(dir string) http.Handler {
	return http.StripPrefix(r.StaticURL, http.FileServer(http.Dir(dir)))
}

// StartServer starts the HTTP server to listen for incoming requests.
// - port: The port on which the server should listen for incoming requests.
func (r *RegoServer) StartServer(port int) {
	mux := http.NewServeMux()
	for path, pageHandler := range r.Handlers {
		if path != r.StaticURL {
			if pageHandler.Handler != nil {
				if pageHandler.Concurrent {
					go mux.HandleFunc(path, pageHandler.Handler)
				} else {
					mux.HandleFunc(path, pageHandler.Handler)
				}
			} else {
				if r.StaticPagesConcurrency {
					go mux.HandleFunc(path, r.ServeHTML)
				} else {
					mux.HandleFunc(path, r.ServeHTML)
				}
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

