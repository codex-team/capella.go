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

// Capella uploading URL.
// https://github.com/codex-team/capella
const (
	API_URL = "https://capella.pics/upload"
)

// Returns capella error message
// Implements error interface and releases custom Error function
type CapellaError struct {
	// text message from CodeX capella. Takes value of nil if Response success is true
	Message string
}

// Custom Error function that returns text Message
func (err *CapellaError) Error() string {
	// Formatted message
	return fmt.Sprintf("%s", err.Message)
}

// CodeX Capella formatted response
type Response struct {
	// uploaded image ID
	ID string `json:"id"`
	// full URL to the image
	URL string `json:"url"`
	// success condition.
	Success bool `json:"success"`
	// error message according to the Capella API
	Message string `json:"message"`
}

// Method uploads file from local path
// It is important to use absolute path to file so that os.Open could find the file
// in error case client get's CapellaError
// Success case returns Response struct type that describes all properties
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

// This method uploads file from passed url
// Responses the same struct which were described above
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