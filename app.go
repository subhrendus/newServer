package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"


	"bitbucket.org/subhrendu_sarkar/uploadserver/src/config"
	"bitbucket.org/subhrendu_sarkar/uploadserver/src/controllers"
)

// PORT - The default port no, used config doesn't have a port no defined.
const PORT = 8800
const AppName = "uploadServer"