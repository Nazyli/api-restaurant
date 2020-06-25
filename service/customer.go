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

func (s *svc) GetCustomerByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (customer *entity.Customer, status Status) {
	customer, err := s.customer.GetByID(ctx, s.AppID, id, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return customer, Status{http.StatusNotFound, "Customer"}
	}
	if err != nil {
		log.Println(err)
		return customer, Status{http.StatusInternalServerError, "Customer"}
	}
	return customer, Status{http.StatusOK, ""}
}
func (s *svc) SelectCustomers(ctx context.Context, all bool, isAdmin bool, uid string) (customers entity.Customers, status Status) {
	customer, err := s.customer.Select(ctx, s.AppID, all, isAdmin, uid)
	if err != nil {
		log.Println(err)
		return customer, Status{http.StatusInternalServerError, ""}
	}
	return customer, Status{http.StatusOK, ""}
}

func (s *svc) InsertCustomer(ctx context.Context, uid string, customer *entity.Customer) (customerData *entity.Customer, status Status) {
	// Add
	customer.CreatedAt = null.TimeFrom(time.Now())
	customer.CreatedBy = uid
	customer.IsActive = 1
	customer.AppID = s.AppID
	err := s.customer.Insert(ctx, customer)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Customer"}
	}

	return customer, Status{http.StatusOK, ""}
}

func (s *svc) UpdateCustomer(ctx context.Context, isAdmin bool, uid string, customer *entity.Customer) (customerData *entity.Customer, status Status) {
	getCustomer, status := s.GetCustomerByID(ctx, customer.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return customer, status
	}
	customer.UpdatedAt = null.TimeFrom(time.Now())
	customer.AppID = s.AppID
	customer.CreatedBy = getCustomer.CreatedBy
	customer.LastUpdateBy = &uid
	err := s.customer.Update(ctx, isAdmin, customer)
	if err != nil {
		log.Println(err)
		return customer, Status{http.StatusInternalServerError, "Customer"}
	}

	// kirim response
	customerData, status = s.GetCustomerByID(ctx, customer.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return customer, status
	}
	return customerData, Status{http.StatusOK, ""}
}

func (s *svc) DeleteCustomer(ctx context.Context, id int64, isAdmin bool, uid string) (status Status) {
	getCustomer, status := s.GetCustomerByID(ctx, s.AppID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return status
	}

	customer := &entity.Customer{
		ID:           id,
		LastUpdateBy: &uid,
		AppID:        s.AppID,
		DeletedAt:    null.TimeFrom(time.Now()),
		CreatedBy:    getCustomer.CreatedBy,
	}
	err := s.customer.Delete(ctx, isAdmin, customer)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "Customer"}
	}
	return Status{http.StatusOK, ""}
}
