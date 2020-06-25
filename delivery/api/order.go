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

func (api *API) handleSelectOrders(w http.ResponseWriter, r *http.Request) {
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

	order, status := api.service.SelectOrder(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Orders")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(order))
	for _, i := range order {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "Order",
			Attributes: entity.Order{
				ID:            i.ID,
				InvoiceNum:    i.InvoiceNum,
				SaleDate:      i.SaleDate,
				SaleTime:      i.SaleTime,
				SubTotal:      i.SubTotal,
				Tax:           i.Tax,
				Total:         i.Total,
				Cash:          i.Cash,
				ChangeMoney:   i.ChangeMoney,
				Other:         i.Other,
				PaymentStatus: i.PaymentStatus,
				CustomerID:    i.CustomerID,
				EmployeeID:    i.EmployeeID,
				AppID:         i.AppID,
				CreatedAt:     i.CreatedAt,
				CreatedBy:     i.CreatedBy,
				UpdatedAt:     i.UpdatedAt,
				LastUpdateBy:  i.LastUpdateBy,
				DeletedAt:     i.DeletedAt,
				IsActive:      i.IsActive,
			},
		})
	}
	responses.OK(w, res)
}

func (api *API) handleGetOrderByInv(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		uid      string
		all      = false
		err      error
	)
	inv := mux.Vars(r)["inv"]
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
	order, status := api.service.GetOrderByInv(r.Context(), inv, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Order ", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   order.ID,
		Type: "Order ",
		Attributes: entity.Order{
			ID:            order.ID,
			InvoiceNum:    order.InvoiceNum,
			SaleDate:      order.SaleDate,
			SaleTime:      order.SaleTime,
			SubTotal:      order.SubTotal,
			Tax:           order.Tax,
			Total:         order.Total,
			Cash:          order.Cash,
			ChangeMoney:   order.ChangeMoney,
			Other:         order.Other,
			PaymentStatus: order.PaymentStatus,
			CustomerID:    order.CustomerID,
			EmployeeID:    order.EmployeeID,
			AppID:         order.AppID,
			CreatedAt:     order.CreatedAt,
			CreatedBy:     order.CreatedBy,
			UpdatedAt:     order.UpdatedAt,
			LastUpdateBy:  order.LastUpdateBy,
			DeletedAt:     order.DeletedAt,
			IsActive:      order.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) handleSelectCalculateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		uid, inv string
	)
	uid, _ = auth.IsAdmin(r)
	params := mux.Vars(r)
	inv = params["inv"]
	data, status := api.service.CalculateOrder(r.Context(), inv, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Calculate Order", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:         data.ID,
		Type:       "Calculate Order",
		Attributes: data,
	}
	responses.OK(w, res)
}

func (api *API) handlePostOrder(w http.ResponseWriter, r *http.Request) {
	uid, _ := auth.IsAdmin(r)
	order, status := api.service.InsertOrder(r.Context(), uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Employee", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   order.ID,
		Type: "Order",
		Attributes: entity.OrderID{
			ID:         order.ID,
			InvoiceNum: order.InvoiceNum,
			EmployeeID: order.EmployeeID,
			AppID:      order.AppID,
			CreatedAt:  order.CreatedAt,
			CreatedBy:  order.CreatedBy,
			IsActive:   order.IsActive,
		},
	}
	responses.OK(w, res)

}

func (api *API) handlePatchPaymentOrder(w http.ResponseWriter, r *http.Request) {
	var (
		params   reqPayment
		paramsID = mux.Vars(r)
	)
	inv := paramsID["inv"]

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
	uid, isAdmin := auth.IsAdmin(r)
	order := &entity.Order{
		Cash:       &params.Cash,
		CustomerID: params.CustomerID,
		Other:      params.Other,
	}
	order, orderDetail, status := api.service.PaymentOrder(r.Context(), inv, isAdmin, uid, order)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Order", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   order.ID,
		Type: "Order",
		Attributes: entity.OrderData{
			ID:            order.ID,
			InvoiceNum:    order.InvoiceNum,
			SaleDate:      order.SaleDate,
			SaleTime:      order.SaleTime,
			SubTotal:      order.SubTotal,
			Tax:           order.Tax,
			Total:         order.Total,
			Cash:          order.Cash,
			ChangeMoney:   order.ChangeMoney,
			Other:         order.Other,
			PaymentStatus: order.PaymentStatus,
			CustomerID:    order.CustomerID,
			EmployeeID:    order.EmployeeID,
			AppID:         order.AppID,
			CreatedAt:     order.CreatedAt,
			CreatedBy:     order.CreatedBy,
			UpdatedAt:     order.UpdatedAt,
			LastUpdateBy:  order.LastUpdateBy,
			DeletedAt:     order.DeletedAt,
			IsActive:      order.IsActive,
			OrderDetail:   orderDetail,
		},
	}
	responses.OK(w, res)
}

func (api *API) handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	inv := mux.Vars(r)["inv"]
	uid, isAdmin := auth.IsAdmin(r)

	status := api.service.DeleteOrder(r.Context(), inv, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete Order", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
