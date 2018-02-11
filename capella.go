package capella

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/url"
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