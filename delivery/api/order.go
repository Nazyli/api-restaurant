package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
	"gopkg.in/go-playground/validator.v9"
)

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
