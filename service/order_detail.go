package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/nazyli/api-restaurant/entity"
	null "gopkg.in/guregu/null.v3"
)

func (s *svc) SelectOrderDetailByInv(ctx context.Context, inv string, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, status Status) {
	orderDetail, err := s.orderDetail.SelectByInv(ctx, s.AppID, inv, all, isAdmin, uid)
	if err != nil {
		log.Println(err)
		return orderDetail, Status{http.StatusInternalServerError, "Failed Get Order Details"}
	}
	return orderDetail, Status{http.StatusOK, ""}
}
func (s *svc) GetOrderDetailByID(ctx context.Context, id int64, all, isAdmin bool, uid string) (orderDetail *entity.OrderDetail, status Status) {
	orderDetail, err := s.orderDetail.GetByID(ctx, s.AppID, id, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return orderDetail, Status{http.StatusNotFound, "OrderDetail"}
	}
	if err != nil {
		log.Println(err)
		return orderDetail, Status{http.StatusInternalServerError, "OrderDetail"}
	}
	return orderDetail, Status{http.StatusOK, ""}
}

func (s *svc) InsertOrderDetail(ctx context.Context, uid string, orderDetail *entity.OrderDetail) (orderData *entity.OrderDetail, status Status) {
	var (
		subTotal, price float64
		paymentStatus   int8
	)
	paymentStatus = 1
	order, status := s.GetOrderByInv(ctx, orderDetail.InvoiceNum, false, false, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	if order.PaymentStatus != nil || order.PaymentStatus == &paymentStatus {
		return nil, Status{http.StatusBadRequest, "Order has been paid"}
	}
	menu, status := s.GetMenuByID(ctx, orderDetail.MenuID, false, true, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}

	price = menu.Price - s.CalculatePercent(menu.Price, menu.Discount)

	subTotal = orderDetail.Amount * price
	if orderDetail.Discount != nil {
		subTotal = subTotal - s.CalculatePercent(subTotal, orderDetail.Discount)
	}
	orderDetail.Price = price
	orderDetail.SubTotal = subTotal
	orderDetail.CreatedAt = null.TimeFrom(time.Now())
	orderDetail.CreatedBy = uid
	orderDetail.IsActive = 1
	orderDetail.AppID = s.AppID

	err := s.orderDetail.Insert(ctx, orderDetail)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Order Detail"}
	}
	return orderDetail, Status{http.StatusOK, ""}
}

func (s *svc) UpdateOrderDetail(ctx context.Context, isAdmin bool, uid string, orderDetail *entity.OrderDetail) (orderDetailData *entity.OrderDetail, status Status) {
	dt := time.Now()

	getOrderDetail, status := s.GetOrderDetailByID(ctx, orderDetail.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return nil, status
	}
	orderDetail.UpdatedAt = null.TimeFrom(dt)
	orderDetail.AppID = s.AppID
	orderDetail.CreatedBy = getOrderDetail.CreatedBy
	orderDetail.LastUpdateBy = &uid
	err := s.orderDetail.Update(ctx, isAdmin, orderDetail)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Order Detail"}
	}

	// kirim response
	orderDetailData, status = s.GetOrderDetailByID(ctx, orderDetail.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return orderDetailData, status
	}
	return orderDetailData, Status{http.StatusOK, ""}
}
