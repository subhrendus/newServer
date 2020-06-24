package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	

	"github.com/subhrendus/newServer/config"
	"github.com/subhrendus/newServer/controllers"
	"github.com/subhrendus/newServer/logging"
)

// PORT - The default port no, used config doesn't have a port no defined.
const PORT = 8800
const AppName = "newServer"

// App - a struct to hold the entire application context
type App struct {
	Configs  *config.AppConfig
	Logger   *logging.Logger

	router *mux.Router
}


// Initialize - start the app with a path to config yaml
func (a *App) Initialize(configpath string) {
	var err error

	// load the configs
	a.Configs, err = config.LoadConfiguration(configpath)
	if err != nil {
		fmt.Printf("%s failed to start due to invalid config. error: %+v, config-path: %s", AppName, err.Error(), configpath)
		os.Exit(1) // kill the app
	}

	// initializing logging
	logConfig := logging.LogConfig{
		AppName:    a.Configs.Logging.AppName,
		AppVersion: a.Configs.Logging.AppVersion,
		Level:      a.Configs.Logging.Level,
	}

	logger, err := logging.New(&logConfig)
	if err != nil {
		fmt.Printf("%s failed to init logger. Error: %+v", AppName, err.Error())
		os.Exit(1) // kill the app
	}

	// attach the logger instrument to the app struct
	a.Logger = &logger

	a.router = mux.NewRouter()
	a.initializeRoutes()
}

// Run - run the application with loaded App struct
func (a *App) Run() {
	addr := fmt.Sprintf(":%d", a.getPort())

	a.Logger.Info("starting application", logging.DataFields{"port": addr})
	err := http.ListenAndServe(addr,a.router)
	if err != nil {
		a.Logger.Error("An error occurred starting the server. Going to call exit", logging.DataFields{"Error": err.Error()})
		os.Exit(1)
	}
}

func (a *App) initializeRoutes() {
	a.Logger.Debug("registering routes and preparing controllers")

	// init controllers
	sc, _ := controllers.NewSystemController("BUILD_INFO")
	uc, _ := controllers.NewUploadsController()

	// system specific endpoints, like health-check and build
	a.router.HandleFunc("/system/health", sc.Health).Methods("GET")
	a.router.HandleFunc("/system/build", sc.Build).Methods("GET")
	a.router.HandleFunc("/", uc.GetUpload).Methods("GET")
	a.router.HandleFunc("/upload/{id}", uc.PutUpload).Methods("PUT")


}

// getPort - returns the port number from the config. If no port no
// is defined in the config, returns a default (PORT)
func (a *App) getPort() int {
	if a.Configs.Service.Listen == 0 {
		a.Logger.Debug("no port specified, using default port", logging.DataFields{"port": PORT})
		return PORT
	}
	return a.Configs.Service.Listen
}
