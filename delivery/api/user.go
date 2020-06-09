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
	token, errMsg, status := api.service.SignIn(r.Context(), params.Email, params.Password, api.AppID)
	if status != http.StatusOK {
		responses.ERROR(w, status, "Failed, "+errMsg)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (api *API) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	var (
		Scopes []entity.Scopes
		uid    string
	)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}

	uid, isAdmin := auth.IsAdmin(r)
	if isAdmin {
		uid = ""
	}
	user, status := api.service.GetUserByID(r.Context(), false, uid, id, api.AppID)
	if status != http.StatusOK {
		responses.ERROR(w, status, "Failed Get User")
		return
	}

	// display array scope
	if user.Scope != nil {
		arrScope := strings.Split(*user.Scope, ",")
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
	responses.JSON(w, http.StatusOK, res)

}

func (api *API) handleSelectUsers(w http.ResponseWriter, r *http.Request) {
	var (
		Scopes []entity.Scopes
		uid    string
	)
	uid, isAdmin := auth.IsAdmin(r)
	if isAdmin {
		uid = ""
	}
	user, status := api.service.SelectUsers(r.Context(), false, uid, api.AppID)
	if status != http.StatusOK {
		responses.ERROR(w, status, "Failed Get User")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(user))
	for _, i := range user {
		if i.Scope != nil {
			arrScope := strings.Split(*i.Scope, ",")
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

	responses.JSON(w, http.StatusOK, res)

}
