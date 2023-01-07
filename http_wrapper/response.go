package http_wrapper

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, status int, responseBody interface{}) {
	w.WriteHeader(status)
	if responseBody != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"data": responseBody})
	}
}

func Failure(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(http.StatusText(status)))
}
