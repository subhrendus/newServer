package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type SystemController struct {
	controller
	semVer []byte
}

func NewSystemController(buildPath string) (*SystemController, error) {
	buildInfo, err := ioutil.ReadFile(buildPath)
	if err != nil {
		return nil, fmt.Errorf("Error in instantiating a new system controller. Error: %+v", err.Error())
	}

	return &SystemController{controller: &baseController{}, semVer: buildInfo}, nil
}

// SystemHealth  returns a heart-beat response with status 200
func (sc *SystemController) Health(resp http.ResponseWriter, req *http.Request) {
	sc.successEmpty(resp)
}

// SystemBuild  returns the content of BUILD_INFO (semantic version) with status 200
func (sc *SystemController) Build(resp http.ResponseWriter, req *http.Request) {
	sc.success(resp, req, sc.semVer)
}
