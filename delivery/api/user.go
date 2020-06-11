package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
	"gopkg.in/go-playground/validator.v9"
)

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	var (
		params reqLogin
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
	token, status := api.service.SignIn(r.Context(), params.Email, params.Password)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed", status.ErrMsg)
		return
	}
	responses.OK(w, token)
}

func (api *API) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		Scopes   []entity.Scopes
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
	user, status := api.service.GetUserByID(r.Context(), id, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get User", status.ErrMsg)
		return
	}

	// display array scope
	if user.Scope != "" {
		arrScope := strings.Split(user.Scope, ",")
		Scopes = make([]entity.Scopes, 0, len(arrScope))
		for _, arr := range arrScope {
			Scopes = append(Scopes, entity.Scopes{
				RoleAcess: arr,
			})
		}
	}
	res := DataResponse{
		ID:   user.ID,
		Type: "User",
		Attributes: entity.UserByScope{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			UserHash:     user.UserHash,
			EmployeeID:   user.EmployeeID,
			CreatedAt:    user.CreatedAt,
			CreatedBy:    user.CreatedBy,
			UpdatedAt:    user.UpdatedAt,
			LastUpdateBy: user.LastUpdateBy,
			DeletedAt:    user.DeletedAt,
			IsActive:     user.IsActive,
			AppID:        user.AppID,
			Scope:        Scopes,
		},
	}
	responses.OK(w, res)

}

func (api *API) handleSelectUsers(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		Scopes   []entity.Scopes
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

	user, status := api.service.SelectUsers(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Users")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(user))
	for _, i := range user {
		if i.Scope != "" {
			arrScope := strings.Split(i.Scope, ",")
			Scopes = make([]entity.Scopes, 0, len(arrScope))
			for _, arr := range arrScope {
				Scopes = append(Scopes, entity.Scopes{
					RoleAcess: arr,
				})
			}
		}
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "User",
			Attributes: entity.UserByScope{
				ID:           i.ID,
				Username:     i.Username,
				Email:        i.Email,
				UserHash:     i.UserHash,
				EmployeeID:   i.EmployeeID,
				CreatedAt:    i.CreatedAt,
				CreatedBy:    i.CreatedBy,
				UpdatedAt:    i.UpdatedAt,
				LastUpdateBy: i.LastUpdateBy,
				DeletedAt:    i.DeletedAt,
				IsActive:     i.IsActive,
				AppID:        i.AppID,
				Scope:        Scopes,
			},
		})
	}

	responses.OK(w, res)

}

func (api *API) handlePostUsers(w http.ResponseWriter, r *http.Request) {
	var (
		params reqUser
		Scopes []entity.Scopes
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
	if params.Password == "" {
		responses.ERROR(w, http.StatusBadRequest, "Password cannot be null")
		return
	}
	uid, _ := auth.IsAdmin(r)
	user := &entity.User{
		Username:   params.Username,
		Email:      params.Email,
		Password:   params.Password,
		EmployeeID: params.EmployeeID,
		Scope:      params.Scope,
		CreatedBy:  uid,
	}
	user, status := api.service.InsertUser(r.Context(), user)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert User", status.ErrMsg)
		return
	}
	if user.Scope != "" {
		arrScope := strings.Split(user.Scope, ",")
		Scopes = make([]entity.Scopes, 0, len(arrScope))
		for _, arr := range arrScope {
			Scopes = append(Scopes, entity.Scopes{
				RoleAcess: arr,
			})
		}
	}
	res := DataResponse{
		ID:   user.ID,
		Type: "User",
		Attributes: entity.UserByScope{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			UserHash:     user.UserHash,
			EmployeeID:   user.EmployeeID,
			CreatedAt:    user.CreatedAt,
			CreatedBy:    user.CreatedBy,
			UpdatedAt:    user.UpdatedAt,
			LastUpdateBy: user.LastUpdateBy,
			DeletedAt:    user.DeletedAt,
			IsActive:     user.IsActive,
			AppID:        user.AppID,
			Scope:        Scopes,
		},
	}
	responses.OK(w, res)

}

func (api *API) handlePatchUsers(w http.ResponseWriter, r *http.Request) {
	var (
		params reqUser
		Scopes []entity.Scopes
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
	user := &entity.User{
		Username:   params.Username,
		Email:      params.Email,
		EmployeeID: params.EmployeeID,
		Scope:      params.Scope,
	}
	user, status := api.service.UpdateUser(r.Context(), id, isAdmin, uid, user)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update User", status.ErrMsg)
		return
	}
	if user.Scope != "" {
		arrScope := strings.Split(user.Scope, ",")
		Scopes = make([]entity.Scopes, 0, len(arrScope))
		for _, arr := range arrScope {
			Scopes = append(Scopes, entity.Scopes{
				RoleAcess: arr,
			})
		}
	}
	res := DataResponse{
		ID:   user.ID,
		Type: "User",
		Attributes: entity.UserByScope{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			UserHash:     user.UserHash,
			EmployeeID:   user.EmployeeID,
			CreatedAt:    user.CreatedAt,
			CreatedBy:    user.CreatedBy,
			UpdatedAt:    user.UpdatedAt,
			LastUpdateBy: user.LastUpdateBy,
			DeletedAt:    user.DeletedAt,
			IsActive:     user.IsActive,
			AppID:        user.AppID,
			Scope:        Scopes,
		},
	}
	responses.OK(w, res)

}

func (api *API) handleDeleteUsers(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	uid, isAdmin := auth.IsAdmin(r)
	status := api.service.DeleteUser(r.Context(), id, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete User", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
