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

func (s *svc) GetEmployeeByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (employee *entity.Employee, status Status) {
	employee, err := s.employee.GetByID(ctx, s.AppID, id, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return employee, Status{http.StatusNotFound, "Employee"}
	}
	if err != nil {
		log.Println(err)
		return employee, Status{http.StatusInternalServerError, "Employee"}
	}
	return employee, Status{http.StatusOK, ""}
}
func (s *svc) SelectEmployees(ctx context.Context, all bool, isAdmin bool, uid string) (employees entity.Employees, status Status) {
	employee, err := s.employee.Select(ctx, s.AppID, all, isAdmin, uid)
	if err != nil {
		log.Println(err)
		return employee, Status{http.StatusInternalServerError, ""}
	}
	return employee, Status{http.StatusOK, ""}
}

func (s *svc) InsertEmployee(ctx context.Context, uid string, employee *entity.Employee) (employeeData *entity.Employee, status Status) {
	// Add
	employee.CreatedAt = null.TimeFrom(time.Now())
	employee.CreatedBy = uid
	employee.IsActive = 1
	employee.AppID = s.AppID
	employee.ShowEmployee = 1
	err := s.employee.Insert(ctx, employee)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Employee"}
	}

	return employee, Status{http.StatusOK, ""}
}

func (s *svc) UpdateEmployee(ctx context.Context, isAdmin bool, uid string, employee *entity.Employee) (employeeData *entity.Employee, status Status) {
	getEmployee, status := s.GetEmployeeByID(ctx, employee.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return employee, status
	}
	employee.UpdatedAt = null.TimeFrom(time.Now())
	employee.AppID = s.AppID
	employee.CreatedBy = getEmployee.CreatedBy
	employee.LastUpdateBy = &uid
	employee.ShowEmployee = 1
	err := s.employee.Update(ctx, isAdmin, employee)
	if err != nil {
		log.Println(err)
		return employee, Status{http.StatusInternalServerError, "Employee"}
	}

	// kirim response
	employeeData, status = s.GetEmployeeByID(ctx, employee.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return employee, status
	}
	return employeeData, Status{http.StatusOK, ""}
}

func (s *svc) DeleteEmployee(ctx context.Context, id int64, isAdmin bool, uid string) (status Status) {
	getEmployee, status := s.GetEmployeeByID(ctx, s.AppID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return status
	}

	employee := &entity.Employee{
		ID:           id,
		LastUpdateBy: &uid,
		AppID:        s.AppID,
		DeletedAt:    null.TimeFrom(time.Now()),
		CreatedBy:    getEmployee.CreatedBy,
	}
	err := s.employee.Delete(ctx, isAdmin, employee)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "Employee"}
	}
	return Status{http.StatusOK, ""}
}
