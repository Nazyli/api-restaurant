package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
	"gopkg.in/go-playground/validator.v9"
)

func (api *API) HandleSelectCustomeres(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		uid      string
		all      = false
		err      error
	)
	uid, isAdmin := auth.IsAdmin(r)
	if isAdmin {
		allParams := getParam.Get("is_active")
		if allParams != "" {
			all, err = strconv.ParseBool(allParams)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, "Is Active must boolean")
				return
			}
		}
	}

	customer, status := api.Service.SelectCustomers(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Customers")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(customer))
	for _, i := range customer {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "Customer",
			Attributes: entity.Customer{
				ID:           i.ID,
				Name:         i.Name,
				Email:        i.Email,
				Addreas:      i.Addreas,
				AppID:        i.AppID,
				CreatedAt:    i.CreatedAt,
				CreatedBy:    i.CreatedBy,
				UpdatedAt:    i.UpdatedAt,
				LastUpdateBy: i.LastUpdateBy,
				DeletedAt:    i.DeletedAt,
				IsActive:     i.IsActive,
			},
		})
	}
	responses.OK(w, res)
}

func (api *API) HandleGetCustomerById(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		uid      string
		all      = false
	)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}

	uid, isAdmin := auth.IsAdmin(r)
	if isAdmin {
		allParams := getParam.Get("is_active")
		if allParams != "" {
			all, err = strconv.ParseBool(allParams)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, "Is Active must boolean")
				return
			}
		}
	}
	customer, status := api.Service.GetCustomerByID(r.Context(), id, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Customer", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   customer.ID,
		Type: "Customer",
		Attributes: entity.Customer{
			ID:           customer.ID,
			Name:         customer.Name,
			Email:        customer.Email,
			Addreas:      customer.Addreas,
			AppID:        customer.AppID,
			CreatedAt:    customer.CreatedAt,
			CreatedBy:    customer.CreatedBy,
			UpdatedAt:    customer.UpdatedAt,
			LastUpdateBy: customer.LastUpdateBy,
			DeletedAt:    customer.DeletedAt,
			IsActive:     customer.IsActive,
		},
	}
	responses.OK(w, res)
}

func (api *API) HandlePostCustomers(w http.ResponseWriter, r *http.Request) {
	var (
		params reqCustomer
	)
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}
	v := validator.New()
	err = v.Struct(params)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}
	uid, _ := auth.IsAdmin(r)
	customer := &entity.Customer{
		Name:    params.Name,
		Email:   params.Email,
		Addreas: params.Addreas,
	}
	customer, status := api.Service.InsertCustomer(r.Context(), uid, customer)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Customer", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   customer.ID,
		Type: "Customer",
		Attributes: entity.Customer{
			ID:           customer.ID,
			Name:         customer.Name,
			Email:        customer.Email,
			Addreas:      customer.Addreas,
			AppID:        customer.AppID,
			CreatedAt:    customer.CreatedAt,
			CreatedBy:    customer.CreatedBy,
			UpdatedAt:    customer.UpdatedAt,
			LastUpdateBy: customer.LastUpdateBy,
			DeletedAt:    customer.DeletedAt,
			IsActive:     customer.IsActive,
		},
	}
	responses.OK(w, res)

}

func (api *API) HandlePatchCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		params reqCustomer
	)
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}
	v := validator.New()
	err = v.Struct(params)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}
	uid, isAdmin := auth.IsAdmin(r)
	customer := &entity.Customer{
		ID:      id,
		Name:    params.Name,
		Email:   params.Email,
		Addreas: params.Addreas,
	}
	customer, status := api.Service.UpdateCustomer(r.Context(), isAdmin, uid, customer)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Customer", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   customer.ID,
		Type: "Customer",
		Attributes: entity.Customer{
			ID:           customer.ID,
			Name:         customer.Name,
			Email:        customer.Email,
			Addreas:      customer.Addreas,
			AppID:        customer.AppID,
			CreatedAt:    customer.CreatedAt,
			CreatedBy:    customer.CreatedBy,
			UpdatedAt:    customer.UpdatedAt,
			LastUpdateBy: customer.LastUpdateBy,
			DeletedAt:    customer.DeletedAt,
			IsActive:     customer.IsActive,
		},
	}
	responses.OK(w, res)

}
func (api *API) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	uid, isAdmin := auth.IsAdmin(r)
	status := api.Service.DeleteCustomer(r.Context(), id, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete User", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
