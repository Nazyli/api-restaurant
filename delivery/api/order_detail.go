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

func (api *API) handleSelectOrderDetails(w http.ResponseWriter, r *http.Request) {
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

	orderDetail, status := api.service.SelectOrderDetail(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Order Details")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(orderDetail))
	for _, i := range orderDetail {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "OrderDetail",
			Attributes: entity.OrderDetail{
				ID:           i.ID,
				InvoiceNum:   i.InvoiceNum,
				MenuID:       i.MenuID,
				Amount:       i.Amount,
				Price:        i.Price,
				Discount:     i.Discount,
				SubTotal:     i.SubTotal,
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
func (api *API) handleGetOrderDetailById(w http.ResponseWriter, r *http.Request) {
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
	orderDetail, status := api.service.GetOrderDetailByID(r.Context(), id, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Order Detail", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   orderDetail.ID,
		Type: "Order Detail",
		Attributes: entity.OrderDetail{
			ID:           orderDetail.ID,
			InvoiceNum:   orderDetail.InvoiceNum,
			MenuID:       orderDetail.MenuID,
			Amount:       orderDetail.Amount,
			Price:        orderDetail.Price,
			Discount:     orderDetail.Discount,
			SubTotal:     orderDetail.SubTotal,
			AppID:        orderDetail.AppID,
			CreatedAt:    orderDetail.CreatedAt,
			CreatedBy:    orderDetail.CreatedBy,
			UpdatedAt:    orderDetail.UpdatedAt,
			LastUpdateBy: orderDetail.LastUpdateBy,
			DeletedAt:    orderDetail.DeletedAt,
			IsActive:     orderDetail.IsActive,
		},
	}
	responses.OK(w, res)
}

func (api *API) handleGetOrderDetailByInv(w http.ResponseWriter, r *http.Request) {
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

	orderDetail, status := api.service.SelectOrderDetailByInv(r.Context(), inv, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Order Details")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(orderDetail))
	for _, i := range orderDetail {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "OrderDetail",
			Attributes: entity.OrderDetail{
				ID:           i.ID,
				InvoiceNum:   i.InvoiceNum,
				MenuID:       i.MenuID,
				Amount:       i.Amount,
				Price:        i.Price,
				Discount:     i.Discount,
				SubTotal:     i.SubTotal,
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
func (api *API) handlePostOrderDetail(w http.ResponseWriter, r *http.Request) {
	var (
		params reqOrderDetail
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
	if params.InvoiceNum == "" {
		responses.ERROR(w, http.StatusBadRequest, "Invalid Parameter, Invoice Number could not be nil")
		return
	}
	uid, _ := auth.IsAdmin(r)
	orderDetail := &entity.OrderDetail{
		InvoiceNum: params.InvoiceNum,
		MenuID:     params.MenuID,
		Amount:     params.Amount,
		Discount:   params.Discount,
	}
	orderDetail, status := api.service.InsertOrderDetail(r.Context(), uid, orderDetail)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Order Detail", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   orderDetail.ID,
		Type: "Order Detail",
		Attributes: entity.OrderDetail{
			ID:           orderDetail.ID,
			InvoiceNum:   orderDetail.InvoiceNum,
			MenuID:       orderDetail.MenuID,
			Amount:       orderDetail.Amount,
			Price:        orderDetail.Price,
			Discount:     orderDetail.Discount,
			SubTotal:     orderDetail.SubTotal,
			AppID:        orderDetail.AppID,
			CreatedAt:    orderDetail.CreatedAt,
			CreatedBy:    orderDetail.CreatedBy,
			UpdatedAt:    orderDetail.UpdatedAt,
			LastUpdateBy: orderDetail.LastUpdateBy,
			DeletedAt:    orderDetail.DeletedAt,
			IsActive:     orderDetail.IsActive,
		},
	}
	responses.OK(w, res)

}

func (api *API) handlePatchOrderDetail(w http.ResponseWriter, r *http.Request) {
	var (
		params reqOrderDetail
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
	orderDetail := &entity.OrderDetail{
		ID:       id,
		MenuID:   params.MenuID,
		Amount:   params.Amount,
		Discount: params.Discount,
	}
	orderDetail, status := api.service.UpdateOrderDetail(r.Context(), isAdmin, uid, orderDetail)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update OrderDetail", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   orderDetail.ID,
		Type: "Order Detail",
		Attributes: entity.OrderDetail{
			ID:           orderDetail.ID,
			InvoiceNum:   orderDetail.InvoiceNum,
			MenuID:       orderDetail.MenuID,
			Amount:       orderDetail.Amount,
			Price:        orderDetail.Price,
			Discount:     orderDetail.Discount,
			SubTotal:     orderDetail.SubTotal,
			AppID:        orderDetail.AppID,
			CreatedAt:    orderDetail.CreatedAt,
			CreatedBy:    orderDetail.CreatedBy,
			UpdatedAt:    orderDetail.UpdatedAt,
			LastUpdateBy: orderDetail.LastUpdateBy,
			DeletedAt:    orderDetail.DeletedAt,
			IsActive:     orderDetail.IsActive,
		},
	}
	responses.OK(w, res)
}

func (api *API) handleDeleteOrderDetail(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}

	uid, isAdmin := auth.IsAdmin(r)

	status := api.service.DeleteOrderDetail(r.Context(), id, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete OrderDetail", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
