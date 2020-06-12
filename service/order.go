package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nazyli/api-restaurant/entity"
	null "gopkg.in/guregu/null.v3"
)

func (s *svc) GetOrderByInv(ctx context.Context, inv string, all bool, isAdmin bool, uid string) (order *entity.Order, status Status) {
	order, err := s.order.GetByInv(ctx, s.AppID, inv, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return order, Status{http.StatusNotFound, "Order"}
	}
	if err != nil {
		log.Println(err)
		return order, Status{http.StatusInternalServerError, "Order"}
	}
	return order, Status{http.StatusOK, ""}
}
func (s *svc) CalculateOrder(ctx context.Context, inv string, uid string) (orderData *entity.CalculateOrder, status Status) {
	var paymentStatus int8
	paymentStatus = 1
	dt := time.Now()
	saleDate := dt.Format("2006-01-02")
	saleTime := dt.Format("15:04:05")
	var subTotal, total float64
	order, status := s.GetOrderByInv(ctx, inv, false, true, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	if order.PaymentStatus != nil || order.PaymentStatus == &paymentStatus {
		return nil, Status{http.StatusBadRequest, "Order has been paid"}
	}
	orderDetail, status := s.SelectOrderDetailByInv(ctx, inv, false, true, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	// res := make([]entity.OrderDetails, 0, len(orderDetail))
	for _, p := range orderDetail {
		subTotal += p.SubTotal
	}
	total = subTotal - s.CalculatePercent(subTotal, &s.Tax)
	orderData = &entity.CalculateOrder{
		ID:          order.ID,
		InvoiceNum:  order.InvoiceNum,
		SaleDate:    saleDate,
		SaleTime:    saleTime,
		SubTotal:    subTotal,
		Tax:         s.Tax,
		Total:       total,
		EmployeeID:  order.EmployeeID,
		AppID:       order.AppID,
		CreatedAt:   order.CreatedAt,
		CreatedBy:   order.CreatedBy,
		OrderDetail: orderDetail,
	}
	return orderData, Status{http.StatusOK, ""}
}
func (s *svc) PaymentOrder(ctx context.Context, inv string, isAdmin bool, uid string, order *entity.Order) (a *entity.Order, b entity.OrderDetails, status Status) {
	var paymentStatus int8
	paymentStatus = 1
	orderData, status := s.CalculateOrder(ctx, inv, uid)
	if status.Code != http.StatusOK {
		return nil, nil, status
	}
	// cek money
	if orderData.Total > *order.Cash {
		return nil, nil, Status{http.StatusBadRequest, "Payment below the total price"}
	}
	ch := *order.Cash - orderData.Total

	if order.CustomerID != nil {
		_, status = s.GetCustomerByID(ctx, *order.CustomerID, false, true, uid)
		if status.Code != http.StatusOK {
			return nil, nil, status
		}
	}
	// cek employee id
	user, status := s.GetUserByHash(ctx, false, false, uid)
	if status.Code != http.StatusOK {
		return nil, nil, status
	}
	if user.EmployeeID == nil {
		return nil, nil, Status{http.StatusNotFound, "Employe ID tidak terdaftar di User Data, "}
	}

	// Update Order Details update_at and last_update_by
	res := make(entity.OrderDetails, 0, len(orderData.OrderDetail))
	for _, p := range orderData.OrderDetail {
		orderDetail, status := s.UpdateOrderDetail(ctx, isAdmin, uid, &p)
		if status.Code != http.StatusOK {
			return nil, nil, status
		}
		res = append(res, *orderDetail)
	}

	// update data
	order.InvoiceNum = orderData.InvoiceNum
	order.SaleDate = &orderData.SaleDate
	order.SaleTime = &orderData.SaleTime
	order.SubTotal = &orderData.SubTotal
	order.Tax = &orderData.Tax
	order.Total = &orderData.Total
	order.ChangeMoney = &ch
	order.PaymentStatus = &paymentStatus
	order.EmployeeID = *user.EmployeeID
	order.UpdatedAt = null.TimeFrom(time.Now())
	order.LastUpdateBy = &uid
	// Update Order
	order, status = s.UpdateOrder(ctx, isAdmin, uid, order)
	if status.Code != http.StatusOK {
		return nil, nil, status
	}
	return order, res, Status{http.StatusOK, ""}
}
func (s *svc) InsertOrder(ctx context.Context, uid string) (orderData *entity.Order, status Status) {
	user, status := s.GetUserByHash(ctx, false, false, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	if user.EmployeeID == nil {
		return nil, Status{http.StatusNotFound, "Employe ID tidak terdaftar di User Data, "}
	}
	// Add
	order := &entity.Order{
		InvoiceNum: s.invoice(),
		CreatedAt:  null.TimeFrom(time.Now()),
		CreatedBy:  uid,
		IsActive:   1,
		AppID:      s.AppID,
		EmployeeID: *user.EmployeeID,
	}
	err := s.order.Insert(ctx, order)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Order"}
	}

	return order, Status{http.StatusOK, ""}
}

func (s *svc) UpdateOrder(ctx context.Context, isAdmin bool, uid string, order *entity.Order) (orderData *entity.Order, status Status) {
	dt := time.Now()

	getOrder, status := s.GetOrderByInv(ctx, order.InvoiceNum, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	order.UpdatedAt = null.TimeFrom(dt)
	order.AppID = s.AppID
	order.CreatedBy = getOrder.CreatedBy
	order.LastUpdateBy = &uid
	order.ID = getOrder.ID
	err := s.order.Update(ctx, isAdmin, order)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Order Detail"}
	}

	// kirim response
	orderData, status = s.GetOrderByInv(ctx, order.InvoiceNum, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return orderData, status
	}
	return orderData, Status{http.StatusOK, ""}
}

func (s *svc) invoice() string {
	dt := time.Now()
	//Format MM-DD-YYYY
	date := dt.Format("01022006")

	//Format MM-DD-YYYY hh:mm:ss
	time := dt.Format("150405")
	key := fmt.Sprintf("INV-%s-%s", date, time)
	return key
}

func (s *svc) CalculatePercent(price float64, percentage *float64) float64 {
	if percentage != nil {
		disc := *percentage
		calculate := price * disc / 1
		return calculate
	}
	return 0
}
