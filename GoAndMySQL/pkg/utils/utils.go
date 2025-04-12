// this is the code to unmarshell the data that we need in the controller that we recieve in the json from the user 

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// this parse bodu function is used to parse especially in the createBook function in the controller 
func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}