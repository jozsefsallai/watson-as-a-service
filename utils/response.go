package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	Data  string `json:"data,omitempty"`
}

// SendJSON will send a JSON response based on the provided parameters.
func SendJSON(w http.ResponseWriter, code int, isError bool, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	var res response
	res.OK = !isError

	if isError {
		res.Error = string(data)
		w.WriteHeader(code)
	} else {
		res.Data = string(data)
	}

	buf, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(buf)
}
