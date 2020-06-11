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
	CDN     CloudinaryConfig
	service _service.Service
}

func New(CDN CloudinaryConfig, service _service.Service) *API {
	return &API{
		CDN:     CDN,
		service: service,
	}
}
func (api *API) Register(r *mux.Router) {
	r.HandleFunc("/ping", middlewares.SetMiddlewareJSON(api.handleGetPing)).Methods("GET")
	r.HandleFunc("/login", middlewares.SetMiddlewareJSON(api.Login)).Methods("POST")
	// User
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetUserById, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.handleSelectUsers, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.handlePostUsers, "create:user")).Methods("POST")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchUsers, "update:user")).Methods("PATCH")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeleteUsers, "delete:user")).Methods("DELETE")

	// Position
	r.HandleFunc("/position", middlewares.SetMiddlewareAuthentication(api.handleSelectPositions, "read:position")).Methods("GET")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetPositionById, "read:position")).Methods("GET")
	r.HandleFunc("/position", middlewares.SetMiddlewareAuthentication(api.handlePostPositions, "create:position")).Methods("POST")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchPositions, "update:position")).Methods("PATCH")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeletePositions, "delete:position")).Methods("DELETE")

	// Category
	r.HandleFunc("/category", middlewares.SetMiddlewareAuthentication(api.handleSelectCategorys, "read:category")).Methods("GET")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetCategoryById, "read:category")).Methods("GET")
	r.HandleFunc("/category", middlewares.SetMiddlewareAuthentication(api.handlePostCategorys, "create:category")).Methods("POST")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchCategorys, "update:category")).Methods("PATCH")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeleteCategorys, "delete:category")).Methods("DELETE")

	// Menu
	r.HandleFunc("/menu", middlewares.SetMiddlewareAuthentication(api.handleSelectMenues, "read:menu")).Methods("GET")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetMenuById, "read:menu")).Methods("GET")
	r.HandleFunc("/menu", middlewares.SetMiddlewareAuthentication(api.handlePostMenus, "create:menu")).Methods("POST")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchMenu, "update:menu")).Methods("PATCH")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeleteMenu, "delete:menu")).Methods("DELETE")

	// Customer
	r.HandleFunc("/customer", middlewares.SetMiddlewareAuthentication(api.handleSelectCustomeres, "read:customer")).Methods("GET")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.handleGetCustomerById, "read:customer")).Methods("GET")
	r.HandleFunc("/customer", middlewares.SetMiddlewareAuthentication(api.handlePostCustomers, "create:customer")).Methods("POST")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.handlePatchCustomer, "update:customer")).Methods("PATCH")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.handleDeleteCustomer, "delete:customer")).Methods("DELETE")

}

func (api *API) handleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}
