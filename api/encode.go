package api

import (
	"io/ioutil"
	"net/http"

	"github.com/jozsefsallai/watson-as-a-service/parser"
	"github.com/jozsefsallai/watson-as-a-service/utils"
)

// EncodeHandler will take the request's body and turn it into a WATSON string.
func EncodeHandler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	var dataType parser.DataType = parser.NewDataType(req.URL.Query().Get("type"))

	if len(body) == 0 {
		utils.SendJSON(res, 400, true, []byte("Request body is empty!"))
		return
	}

	buffer, err := parser.EncodeAs(body, dataType)
	if err != nil {
		utils.SendJSON(res, 500, true, []byte(err.Error()))
		return
	}

	utils.SendJSON(res, 200, false, buffer)
}
