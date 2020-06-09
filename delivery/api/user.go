package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/entity"
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
	token, errMsg, status := api.service.SignIn(r.Context(), params.Email, params.Password)
	if status != http.StatusOK {
		responses.ERROR(w, status, "Failed, "+errMsg)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (api *API) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	var (
		Scopes []entity.Scopes
	)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	user, status := api.service.GetByID(r.Context(), id)
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
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Scope:     Scopes,
		},
	}
	responses.JSON(w, http.StatusOK, res)

}
