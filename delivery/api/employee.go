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

func (api *API) handleSelectEmployees(w http.ResponseWriter, r *http.Request) {
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

	employees, status := api.service.SelectEmployees(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Employees")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(employees))
	for _, i := range employees {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "Employee",
			Attributes: entity.Employee{
				ID:           i.ID,
				PositionID:   i.PositionID,
				Name:         i.Name,
				DateOfBirth:  i.DateOfBirth,
				Address:      i.Address,
				Gender:       i.Gender,
				Email:        i.Email,
				Salary:       i.Salary,
				Bonus:        i.Bonus,
				FromDate:     i.FromDate,
				FinishDate:   i.FinishDate,
				ShowEmployee: i.ShowEmployee,
				AppID:        i.AppID,
				CreatedAt:    i.CreatedAt,
				CreatedBy:    i.CreatedBy,
				UpdatedAt:    i.UpdatedAt,
				LastUpdateBy: i.LastUpdateBy,
				DeletedAt:    i.DeletedAt,
				ImageUrl:     i.ImageUrl,
				ImageID:      i.ImageID,
				IsActive:     i.IsActive,
			},
		})
	}
	responses.OK(w, res)
}

func (api *API) handleGetEmployeeById(w http.ResponseWriter, r *http.Request) {
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
	employee, status := api.service.GetEmployeeByID(r.Context(), id, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Employe", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   employee.ID,
		Type: "Employee",
		Attributes: entity.Employee{
			ID:           employee.ID,
			PositionID:   employee.PositionID,
			Name:         employee.Name,
			DateOfBirth:  employee.DateOfBirth,
			Address:      employee.Address,
			Gender:       employee.Gender,
			Email:        employee.Email,
			Salary:       employee.Salary,
			Bonus:        employee.Bonus,
			FromDate:     employee.FromDate,
			FinishDate:   employee.FinishDate,
			ShowEmployee: employee.ShowEmployee,
			AppID:        employee.AppID,
			CreatedAt:    employee.CreatedAt,
			CreatedBy:    employee.CreatedBy,
			UpdatedAt:    employee.UpdatedAt,
			LastUpdateBy: employee.LastUpdateBy,
			DeletedAt:    employee.DeletedAt,
			ImageUrl:     employee.ImageUrl,
			ImageID:      employee.ImageID,
			IsActive:     employee.IsActive,
		},
	}
	responses.OK(w, res)
}

func (api *API) handlePostEmployees(w http.ResponseWriter, r *http.Request) {
	var (
		params reqEmployee
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
	employee := &entity.Employee{
		Name:        params.Name,
		DateOfBirth: params.DateOfBirth,
		Address:     params.Address,
		Gender:      params.Gender,
		Email:       params.Email,
		Salary:      params.Salary,
		Bonus:       params.Bonus,
		FromDate:    params.FromDate,
		FinishDate:  params.FinishDate,
	}
	employee, status := api.service.InsertEmployee(r.Context(), uid, employee)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Employee", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   employee.ID,
		Type: "Employee",
		Attributes: entity.Employee{
			ID:           employee.ID,
			PositionID:   employee.PositionID,
			Name:         employee.Name,
			DateOfBirth:  employee.DateOfBirth,
			Address:      employee.Address,
			Gender:       employee.Gender,
			Email:        employee.Email,
			Salary:       employee.Salary,
			Bonus:        employee.Bonus,
			FromDate:     employee.FromDate,
			FinishDate:   employee.FinishDate,
			ShowEmployee: employee.ShowEmployee,
			AppID:        employee.AppID,
			CreatedAt:    employee.CreatedAt,
			CreatedBy:    employee.CreatedBy,
			UpdatedAt:    employee.UpdatedAt,
			LastUpdateBy: employee.LastUpdateBy,
			DeletedAt:    employee.DeletedAt,
			ImageUrl:     employee.ImageUrl,
			ImageID:      employee.ImageID,
			IsActive:     employee.IsActive,
		},
	}
	responses.OK(w, res)

}

func (api *API) handlePatchEmployee(w http.ResponseWriter, r *http.Request) {
	var (
		params reqEmployee
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
	employee := &entity.Employee{
		ID:          id,
		Name:        params.Name,
		DateOfBirth: params.DateOfBirth,
		Address:     params.Address,
		Gender:      params.Gender,
		Email:       params.Email,
		Salary:      params.Salary,
		Bonus:       params.Bonus,
		FromDate:    params.FromDate,
		FinishDate:  params.FinishDate,
	}
	employee, status := api.service.UpdateEmployee(r.Context(), isAdmin, uid, employee)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Employee", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   employee.ID,
		Type: "Employee",
		Attributes: entity.Employee{
			ID:           employee.ID,
			PositionID:   employee.PositionID,
			Name:         employee.Name,
			DateOfBirth:  employee.DateOfBirth,
			Address:      employee.Address,
			Gender:       employee.Gender,
			Email:        employee.Email,
			Salary:       employee.Salary,
			Bonus:        employee.Bonus,
			FromDate:     employee.FromDate,
			FinishDate:   employee.FinishDate,
			ShowEmployee: employee.ShowEmployee,
			AppID:        employee.AppID,
			CreatedAt:    employee.CreatedAt,
			CreatedBy:    employee.CreatedBy,
			UpdatedAt:    employee.UpdatedAt,
			LastUpdateBy: employee.LastUpdateBy,
			DeletedAt:    employee.DeletedAt,
			ImageUrl:     employee.ImageUrl,
			ImageID:      employee.ImageID,
			IsActive:     employee.IsActive,
		},
	}
	responses.OK(w, res)

}
func (api *API) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	uid, isAdmin := auth.IsAdmin(r)
	status := api.service.DeleteEmployee(r.Context(), id, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete User", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
