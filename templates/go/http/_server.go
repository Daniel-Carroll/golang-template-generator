package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"{{ module }}/domain"
)

const (
	//Environments
	local = "local"
	//Errors
	errorMissingEnvironmentVariable = "Missing Variable For Environment"
)

// Server : configurability options
type Server struct {
	ln net.Listener

	// Services
	ExampleService domain.ExampleService
	// Server options
	Environment string // "local" | "dev" | "cert" | "uat" | "prod"
	Recoverable bool   // panic recovery

	LogOutput io.Writer
}

// NewServer : returns a new instance of Server.
func NewServer(env string) *Server {
	return &Server{
		Recoverable: true,
		LogOutput:   ioutil.Discard,
		Environment: env,
	}
}

// Open : opens the server connection on specified port
func (s *Server) Open() error {
	fmt.Println("Server listening on localhost:3000")
	http.ListenAndServe(":3000", s.router())
	return nil
}

// Close : closes the socket
func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/health", s.healthCheckHandler())
		r.Mount("/example", s.exampleHandler())
	})

	return r
}

// healthCheckHandler : verifies the http server connection and returns a success.
func (s *Server) healthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, http.StatusOK, "status: ok")
	}
}

// Manage Examples
func (s *Server) exampleHandler() *exampleHandler {
	log.Info("Creating exampleHandler")
	h := newExampleHandler()
	h.exampleService = s.Example
	return h
}
