package router

import (
	"github.com/gorilla/mux"

	"github.com/alifradityar/votr/votr-api/handler"
)

// CreateRouter create new router object with all routing definition
func CreateRouter(rh handler.Root) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/topic/all", rh.GetAllTopicHandler).Methods("GET")
	router.HandleFunc("/topic", rh.GetTopicPageHandler).Methods("GET")
	router.HandleFunc("/topic", rh.CreateTopicHandler).Methods("POST")
	// router.HandleFunc("/topic/{id}/upvote")
	return router
}
