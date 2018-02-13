// Copyright 2018 CodeX
// Go SDK for Capella
// license that can be found in the LICENSE file.

package capella

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/url"
	"os"
	"bytes"
	"mime/multipart"
	"io"
)

const (
	API_URL = "https://capella.pics/upload"
)

type CapellaError struct {
	Message string
}

func (err *CapellaError) Error() string {
	return fmt.Sprintf("%s", err.Message)
}

type Response struct {
	ID string `json:"id"`
	URL string `json:"url"`
	Success bool `json:"success"`
	Message string `json:"message"`
}

func UploadFile(path string) (response Response, error CapellaError) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", path)
	if err != nil {
		log.Println("error writing to buffer", err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("Can't open file", err)
	}
	defer file.Close()

	//iocopy
	_, err = io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	postResp, err := http.Post(API_URL, contentType, bodyBuf)

	if err != nil {
		panic(err)
	}

	defer postResp.Body.Close()

	message, _ := ioutil.ReadAll(postResp.Body)
	err = json.Unmarshal(message, &response)

	if err != nil {
		log.Fatal("Can't parse file because: ", err)
	}

	if response.Success != true {
		error.Message = response.Message
	}

	return
}

func Upload(uri string) (response Response, error CapellaError) {

	// sending post request to the Capella server
	postResp, postErr := http.PostForm(API_URL, url.Values{"link" : {uri}})
	if postErr != nil {
		log.Println("Error while sending request:", postErr)
	}

	// close body when we got response
	defer postResp.Body.Close()

	body, _ := ioutil.ReadAll(postResp.Body)
	postErr = json.Unmarshal(body, &response)
	if postErr != nil {
		log.Fatal("Can't parse file because: ", postErr)
	}

	if response.Success != true {
		error.Message = response.Message
	}

	return
}