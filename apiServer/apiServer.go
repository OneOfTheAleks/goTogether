package apiServer

import (
	"GoTogether/store"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiServer struct {
	config *Config
	//logger *zap.Logger
	router *mux.Router
	Store *store.Store

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

	if err:= s.ConfigureStore(); err!=nil{

	}

	return http.ListenAndServe(s.config.BindAdrr, s.router)
}

func (s *ApiServer) ConfigureRouter() {
	s.router.HandleFunc("/", s.HandleRoot())
}

func (s *ApiServer)ConfigureStore() error {
  st:= store.New(s.config.Store)
  if err:= st.Open(); err !=nil {
  	return err
  }
  s.Store = st
	return nil
}

