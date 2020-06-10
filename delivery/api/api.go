package api

import (
	"net/http"

	"github.com/gorilla/mux"
	_service "github.com/nazyli/api-restaurant/service"
	"github.com/nazyli/api-restaurant/util/middlewares"
)

type CloudinaryConfig struct {
	AccountName string
	APIKey      string
	APISecret   string
}

// API struct
type API struct {
	AppID   int64
	CDN     CloudinaryConfig
	service _service.Service
}

func New(AppID int64, CDN CloudinaryConfig, service _service.Service) *API {
	return &API{
		AppID:   AppID,
		CDN:     CDN,
		service: service,
	}
}
func (api *API) Register(r *mux.Router) {
	r.HandleFunc("/ping", middlewares.SetMiddlewareJSON(api.handleGetPing)).Methods("GET")
	r.HandleFunc("/login", middlewares.SetMiddlewareJSON(api.Login)).Methods("POST")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetUserById, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.handleSelectUsers, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.handlePostUsers, "create:user")).Methods("POST")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchUsers, "create:user")).Methods("PATCH")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeleteUsers, "create:user")).Methods("DELETE")
}

func (api *API) handleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}
