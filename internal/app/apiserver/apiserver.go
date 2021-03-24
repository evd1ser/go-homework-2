package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/evd1ser/go-homework-2/store"
)

var (
	prefix string = ""
)

// type for APIServer object for instancing server
type APIServer struct {
	//Unexported field
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//APIServer constructor
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server and connection to db and logger confs
func (api *APIServer) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port :", api.config.BindAddr)

	api.configureStore()
	api.configureRouter()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

//func for configureate logger, should be unexported
func (api *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil {
		return nil
	}
	api.logger.SetLevel(level)

	return nil
}

//func for configure Router
func (api *APIServer) configureRouter() {
	api.router.HandleFunc(prefix+"/grab", api.PostGrab).Methods(http.MethodPost)
	api.router.HandleFunc(prefix+"/solve", api.GetSolve).Methods(http.MethodGet)
}

//configureStore method
func (api *APIServer) configureStore() {
	st := store.New()
	api.store = st
	return
}
