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
	router.HandleFunc("/topic", rh.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/topic/{id}", rh.GetTopicHandler).Methods("POST")
	router.HandleFunc("/topic/{id}/upvote", rh.UpvoteTopicHandler).Methods("POST")
	router.HandleFunc("/topic/{id}/downvote", rh.DownvoteTopicHandler).Methods("POST")
	return router
}
