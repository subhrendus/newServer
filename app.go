package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"


	"github.com/subhrendus/newServer/config"
	"github.com/subhrendus/newServer/controllers"
)

// PORT - The default port no, used config doesn't have a port no defined.
const PORT = 8800
const AppName = "uploadServer"