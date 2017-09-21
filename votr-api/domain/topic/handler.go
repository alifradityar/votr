package topic

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alifradityar/votr/votr-api/resp"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Handler handle HTTP request regarding Topic domain
type Handler struct {
	Service *Service `inject:""`
}

// GetAllTopicHandler for Hello world!
func (handler *Handler) GetAllTopicHandler(w http.ResponseWriter, r *http.Request) {
	topics, err := handler.Service.GetAll()
	if err != nil {
		resp.BadRequest(w, topics, err)
		return
	}
	resp.OK(w, topics)
}

func (handler *Handler) GetTopicPageHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	page, _ := strconv.Atoi(r.FormValue("page"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(r.FormValue("size"))
	if size < 1 {
		size = 1
	}
	topics, err := handler.Service.GetTopicPage(keyword, page, size)
	if err != nil {
		resp.BadRequest(w, topics, err)
		return
	}
	resp.OK(w, topics)
}

func (handler *Handler) CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	var topicRequest TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&topicRequest); err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.CreateTopic(topicRequest.Title)
	if err != nil {
		resp.BadRequest(w, topic, err)
		return
	}
	resp.OK(w, topic)
}

func (handler *Handler) UpdateTopicHandler(w http.ResponseWriter, r *http.Request) {
	var topicRequest TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&topicRequest); err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.UpdateTopic(id, topicRequest.Title)
	if err != nil {
		resp.BadRequest(w, topic, err)
		return
	}
	resp.OK(w, topic)
}

func (handler *Handler) DeleteTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	err = handler.Service.DeleteTopic(id)
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	resp.OK(w, nil)
}

func (handler *Handler) UpvoteTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.UpvoteTopic(id)
	if err != nil {
		resp.BadRequest(w, topic, err)
		return
	}
	resp.OK(w, topic)
}

func (handler *Handler) DownvoteTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.DownvoteTopic(id)
	if err != nil {
		resp.BadRequest(w, topic, err)
		return
	}
	resp.OK(w, topic)
}

func (handler *Handler) GetTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		resp.BadRequest(w, nil, err)
		return
	}
	topic, err := handler.Service.GetTopic(id)
	if err != nil {
		resp.BadRequest(w, topic, err)
		return
	}
	resp.OK(w, topic)
}

func (handler *Handler) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	resp.OK(w, nil)
}
