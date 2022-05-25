package apiServer

import (
	apilogger "GoTogether/apiLogger"
	"GoTogether/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *zap.Logger
	store store.Store

}

func  NewServer(store store.Store) *server  {
	 s := &server{
		router: mux.NewRouter(),
		logger: apilogger.CreateLogger(),
		store: store,
	}
	s.configureRouter()
	 return  s
}
func (s *server) ServeHTTP (w http.ResponseWriter,r *http.Request) {
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){
   s.router.HandleFunc("/users", s.UserCreate()).Methods("POST")
   s.router.HandleFunc("/", s.HandleRoot()).Methods("GET")
}