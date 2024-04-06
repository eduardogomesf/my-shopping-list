package webserver

import (
	"fmt"
	"net/http"
	"strings"

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

func (ws *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	if method == "" || path == "" || handler == nil {
		panic("method, path and handler are required")
	}

	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" && method != "PATCH" {
		panic("method must be GET, POST, PUT, PATCH or DELETE")
	}

	key := fmt.Sprintf("%s->%s", method, path)

	ws.Handlers[key] = handler
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)

	ws.AddHandler("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Running on port: %s", ws.WebServerPort)))
	})

	for path, handler := range ws.Handlers {
		keyParts := strings.Split(path, "->")

		method := keyParts[0]
		path = keyParts[1]

		fmt.Println("Adding handler", method, path)

		if method == "GET" {
			ws.Router.Get(path, handler)
		} else if method == "POST" {
			ws.Router.Post(path, handler)
		} else if method == "PUT" {
			ws.Router.Put(path, handler)
		} else if method == "DELETE" {
			ws.Router.Delete(path, handler)
		} else if method == "PATCH" {
			ws.Router.Patch(path, handler)
		}
	}

	http.ListenAndServe(fmt.Sprintf(":%s", ws.WebServerPort), ws.Router)
}
