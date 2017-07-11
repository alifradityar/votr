package topic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Handler handle HTTP request regarding Topic domain
type Handler struct {
	Service *Service `inject:""`
}

// GetAllTopicHandler for Hello world!
func (handler *Handler) GetAllTopicHandler(w http.ResponseWriter, r *http.Request) {
	topics, err := handler.Service.GetAll()
	if err != nil {
		badRequest(w, topics, err)
		return
	}
	ok(w, topics)
}

func (handler *Handler) GetTopicPageHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(r.FormValue("size"))
	if size < 1 {
		size = 1
	}
	topics, err := handler.Service.GetTopicPage(page, size)
	if err != nil {
		badRequest(w, topics, err)
		return
	}
	ok(w, topics)
}

func (handler *Handler) CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	var topicRequest TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&topicRequest); err != nil {
		badRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.CreateTopic(topicRequest.Title)
	if err != nil {
		badRequest(w, topic, err)
		return
	}
	ok(w, topic)
}

// response utility
func ok(w http.ResponseWriter, data interface{}) {
	resp := map[string]interface{}{
		"data":  data,
		"error": nil,
	}
	js, err := json.Marshal(resp)
	if err != nil {
		resp := map[string]interface{}{
			"data":  nil,
			"error": fmt.Sprintf("%s", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}

func badRequest(w http.ResponseWriter, data interface{}, err error) {
	resp := map[string]interface{}{
		"data":  data,
		"error": fmt.Sprintf("%s", err),
	}
	js, err := json.Marshal(resp)
	if err != nil {
		resp := map[string]interface{}{
			"data":  nil,
			"error": fmt.Sprintf("%s", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(js)
}

func internalServerError(w http.ResponseWriter, data interface{}, err error) {
	resp := map[string]interface{}{
		"data":  data,
		"error": fmt.Sprintf("%s", err),
	}
	js, err := json.Marshal(resp)
	if err != nil {
		resp := map[string]interface{}{
			"data":  nil,
			"error": fmt.Sprintf("%s", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(js)
}
