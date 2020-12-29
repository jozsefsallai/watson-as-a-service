package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/genkami/watson"
	"github.com/jozsefsallai/watson-as-a-service/utils"
)

// DecodeHandler will take the request's body (WATSON string) and decode it.
func DecodeHandler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.SendJSON(res, 500, true, []byte(err.Error()))
		return
	}

	if len(body) == 0 {
		utils.SendJSON(res, 400, true, []byte("Request body is empty!"))
		return
	}

	var buffer interface{}
	err = watson.Unmarshal(body, &buffer)
	if err != nil {
		utils.SendJSON(res, 500, true, []byte(err.Error()))
		return
	}

	output, err := json.Marshal(&buffer)
	if err != nil {
		utils.SendJSON(res, 500, true, []byte(err.Error()))
		return
	}

	utils.SendJSON(res, 200, false, output)
}
