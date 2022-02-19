package apiServer

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type ApiServer struct {
	config *Config
	//logger *zap.Logger
	router *mux.Router
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		//	logger: zap.Logger{}
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {

	s.ConfigureRouter()

	return http.ListenAndServe(s.config.BindAdrr, s.router)
}

func (s *ApiServer) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.HandleHallo())
}

func (s *ApiServer) HandleHallo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
