package router

import (
	"github.com/gorilla/mux"

	"github.com/alifradityar/votr/votr-api/handler"
)

// CreateRouter create new router object with all routing definition
func CreateRouter(rh handler.Root) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/topic", rh.HelloHandler).Methods("GET")
	return router
}
