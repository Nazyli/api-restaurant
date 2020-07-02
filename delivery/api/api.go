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
	Service _service.Service
}

func New(CDN CloudinaryConfig, service _service.Service) *API {
	return &API{
		CDN:     CDN,
		Service: service,
	}
}
func (api *API) Register(r *mux.Router) {
	r.HandleFunc("/", api.HandleGetHello).Methods("GET")
	r.HandleFunc("/ping", middlewares.SetMiddlewareJSON(api.HandleGetPing)).Methods("GET")
	r.HandleFunc("/login", middlewares.SetMiddlewareJSON(api.Login)).Methods("POST", "OPTIONS")
	// User
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetUserById, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.HandleSelectUsers, "read:user")).Methods("GET")
	r.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(api.HandlePostUsers, "create:user")).Methods("POST")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchUsers, "update:user")).Methods("PATCH")
	r.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteUsers, "delete:user")).Methods("DELETE")

	// Position
	r.HandleFunc("/position", middlewares.SetMiddlewareJSON(api.HandleSelectPositions)).Methods("GET")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetPositionById, "read:position")).Methods("GET")
	r.HandleFunc("/position", middlewares.SetMiddlewareAuthentication(api.HandlePostPositions, "create:position")).Methods("POST")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchPositions, "update:position")).Methods("PATCH")
	r.HandleFunc("/position/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeletePositions, "delete:position")).Methods("DELETE")

	// Category
	r.HandleFunc("/category", middlewares.SetMiddlewareAuthentication(api.HandleSelectCategorys, "read:category")).Methods("GET")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetCategoryById, "read:category")).Methods("GET")
	r.HandleFunc("/category", middlewares.SetMiddlewareAuthentication(api.HandlePostCategorys, "create:category")).Methods("POST")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchCategorys, "update:category")).Methods("PATCH")
	r.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteCategorys, "delete:category")).Methods("DELETE")

	// Menu
	r.HandleFunc("/menu", middlewares.SetMiddlewareAuthentication(api.HandleSelectMenues, "read:menu")).Methods("GET")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetMenuById, "read:menu")).Methods("GET")
	r.HandleFunc("/menu", middlewares.SetMiddlewareAuthentication(api.HandlePostMenu, "create:menu")).Methods("POST")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchMenu, "update:menu")).Methods("PATCH")
	r.HandleFunc("/menu/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteMenu, "delete:menu")).Methods("DELETE")

	// Customer
	r.HandleFunc("/customer", middlewares.SetMiddlewareAuthentication(api.HandleSelectCustomeres, "read:customer")).Methods("GET")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetCustomerById, "read:customer")).Methods("GET")
	r.HandleFunc("/customer", middlewares.SetMiddlewareAuthentication(api.HandlePostCustomers, "create:customer")).Methods("POST")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchCustomer, "update:customer")).Methods("PATCH")
	r.HandleFunc("/customer/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteCustomer, "delete:customer")).Methods("DELETE")

	// Employee
	r.HandleFunc("/employee", middlewares.SetMiddlewareAuthentication(api.HandleSelectEmployees, "read:employee")).Methods("GET")
	r.HandleFunc("/employee/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetEmployeeById, "read:employee")).Methods("GET")
	r.HandleFunc("/employee", middlewares.SetMiddlewareAuthentication(api.HandlePostEmployees, "create:employee")).Methods("POST")
	r.HandleFunc("/employee/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchEmployee, "update:employee")).Methods("PATCH")
	r.HandleFunc("/employee/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteEmployee, "delete:employee")).Methods("DELETE")

	// Order
	r.HandleFunc("/order", middlewares.SetMiddlewareAuthentication(api.HandleSelectOrders, "read:order")).Methods("GET")
	r.HandleFunc("/order/{inv}", middlewares.SetMiddlewareAuthentication(api.HandleGetOrderByInv, "read:order")).Methods("GET")
	r.HandleFunc("/order", middlewares.SetMiddlewareAuthentication(api.HandlePostOrder, "create:order")).Methods("POST")
	r.HandleFunc("/order/{inv}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteOrder, "delete:order")).Methods("DELETE")
	r.HandleFunc("/calculateorder/{inv}", middlewares.SetMiddlewareAuthentication(api.HandleSelectCalculateOrder, "read:order")).Methods("GET")
	r.HandleFunc("/paymentorder/{inv}", middlewares.SetMiddlewareAuthentication(api.HandlePatchPaymentOrder, "create:order")).Methods("PATCH")

	// OrderDetail
	r.HandleFunc("/orderdetail", middlewares.SetMiddlewareAuthentication(api.HandleSelectOrderDetails, "read:orderdetail")).Methods("GET")
	r.HandleFunc("/orderdetail/{id}", middlewares.SetMiddlewareAuthentication(api.HandleGetOrderDetailById, "read:orderdetail")).Methods("GET")
	r.HandleFunc("/orderdetail_by_inv/{inv}", middlewares.SetMiddlewareAuthentication(api.HandleGetOrderDetailByInv, "read:orderdetail")).Methods("GET")
	r.HandleFunc("/orderdetail", middlewares.SetMiddlewareAuthentication(api.HandlePostOrderDetail, "create:orderdetail")).Methods("POST")
	r.HandleFunc("/orderdetail/{id}", middlewares.SetMiddlewareAuthentication(api.HandlePatchOrderDetail, "update:orderdetail")).Methods("PATCH")
	r.HandleFunc("/orderdetail/{id}", middlewares.SetMiddlewareAuthentication(api.HandleDeleteOrderDetail, "delete:orderdetail")).Methods("DELETE")

}

func (api *API) HandleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}
func (api *API) HandleGetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API-RESTAURANT by Nazyli"))
}
