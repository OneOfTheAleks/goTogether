package apiServer

import (
	apilogger "GoTogether/apiLogger"
	"GoTogether/store"
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type ApiServer struct {
	config *Config
	//logger *apilogger.ApiLogger
	router *mux.Router
	Store *store.Store
	loger *zap.Logger

}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		//	apiLogger: zap.Logger{}
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
    ctx:= context.Background()
	s.ConfigureRouter()

	if err:=s.ConfigureLogger(); err!= nil{
		return err
	}

	if err:= s.ConfigureStore(ctx); err!=nil{
      return err
	}



	return http.ListenAndServe(s.config.BindAdrr, s.router)
}

func (s *ApiServer) ConfigureRouter() {
	s.router.HandleFunc("/", s.HandleRoot())
}

func (s *ApiServer)ConfigureStore(ctx context.Context) error {
  st:= store.New(s.config.Store)
  if err:= st.Open(ctx); err !=nil {
  	return err
  }
  s.Store = st
	return nil
}

func (s *ApiServer)ConfigureLogger() error  {
   lg := apilogger.NewApiLogger(s.config.Logger)
   if err :=lg.Create(); err != nil{
   	return err
	}
	s.loger = lg.Logger
	return nil
}