package controllers

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type UploadsController struct {
	controller
	
}


func NewUploadsController() (*UploadsController, error) {
	

	return &UploadsController{controller: &baseController{}}, nil
}


func (uc *UploadsController) GetUpload(resp http.ResponseWriter, req *http.Request) {

	Filename := req.URL.Query().Get("file")
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(resp, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + Filename)
	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(resp, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	resp.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	resp.Header().Set("Content-Type", FileContentType)
	resp.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(resp, Openfile) //'Copy' the file to the client
	return
	
}


func (uc *UploadsController) PutUpload(resp http.ResponseWriter, req *http.Request) {
	//total := 0
	//maxBytes := 1024 *300 // 300 KB
	//buffer := make([]byte, maxBytes)

	
	params := mux.Vars(req)

	//_, _ := ioutil.ReadAll(req.Body)
	fname := fmt.Sprintf("%s", params["id"] )
	//fmt.Println(str);

	contentType, params, parseErr := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if parseErr != nil || !strings.HasPrefix(contentType, "multipart/") {
		uc.badRequest(resp)
		return
	}

	multipartReader := multipart.NewReader(req.Body, params["boundary"])
	defer req.Body.Close()
	m := make(map[string]int)
	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			uc.internalServerError(resp)
			return
		}
		defer part.Close()
	
		fileBytes, err := ioutil.ReadAll(part)
		for {
			advance, token, err := bufio.ScanLines(fileBytes, true)
			if err != nil {
				uc.internalServerError(resp)
				return
			}
			if advance == 0 {
				break
			}
			if _, ok := m[string(token)]; !ok {
				m[string(token)] = 1
			}
			
			if advance <= len(fileBytes) {
				fileBytes = fileBytes[advance:]
			}
		}
	}

	f, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		uc.internalServerError(resp)
        return
	}
	
	for k, _ := range m {
		_, err = f.WriteString(k)
		if err != nil {
			fmt.Println(err)
			f.Close()
			uc.internalServerError(resp)
			return
		}
		_, err = f.WriteString("\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			uc.internalServerError(resp)
			return
		}
	}

	err = f.Close()
    if err != nil {
		fmt.Println(err)
		uc.internalServerError(resp)
        return
    }


	uc.successEmpty(resp)
	
}
