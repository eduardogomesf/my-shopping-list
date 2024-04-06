package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)

	ws.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Running on port: %s", ws.WebServerPort)))
	})

	for path, handler := range ws.Handlers {
		ws.Router.Handle(path, handler)
	}
	http.ListenAndServe(fmt.Sprintf(":%s", ws.WebServerPort), ws.Router)
}
