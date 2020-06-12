package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
	"gopkg.in/go-playground/validator.v9"
)

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
