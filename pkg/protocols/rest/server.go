package rest

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bungysheep/news-api/pkg/configs"
	"github.com/bungysheep/news-api/pkg/protocols/rest/routes"
)

// Server type
type Server struct {
	*http.Server
}

// NewRestServer creates new rest server
func NewRestServer() *Server {
	return &Server{}
}

// RunServer runs rest server
func (s *Server) RunServer() error {

	port := resolvePortNumber()

	s.Server = &http.Server{
		Addr:         ":" + port,
		Handler:      routes.APIV1RouteHandler(),
		ReadTimeout:  configs.READTIMEOUT * time.Second,
		WriteTimeout: configs.WRITETIMEOUT * time.Second,
	}

	log.Printf("Listening on port %s...\n", port)

	return s.Server.ListenAndServe()
}

func resolvePortNumber() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}

	return configs.PORT
}
