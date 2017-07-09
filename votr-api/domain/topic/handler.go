package topic

import (
	"encoding/json"
	"net/http"
)

// Handler handle HTTP request regarding Topic domain
type Handler struct{}

// HelloHandler for Hello world!
func (handler *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"data":  "Hello world!",
		"error": nil,
	}
	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(respJSON)
}
