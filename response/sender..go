package response

import (
	"encoding/json"
	"net/http"
)

// SendJSON json data to client
func SendJSON(w http.ResponseWriter, msg interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}
