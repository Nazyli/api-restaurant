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

func (api *API) HandleSelectMenues(w http.ResponseWriter, r *http.Request) {
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

	menu, status := api.Service.SelectMenues(r.Context(), all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Menues")
		return
	}

	// display array scope
	res := make([]DataResponse, 0, len(menu))
	for _, i := range menu {
		res = append(res, DataResponse{
			ID:   i.ID,
			Type: "Menu",
			Attributes: entity.Menu{
				ID:           i.ID,
				CategoryID:   i.CategoryID,
				Name:         i.Name,
				Price:        i.Price,
				Discount:     i.Discount,
				ShowMenu:     i.ShowMenu,
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

func (api *API) HandleGetMenuById(w http.ResponseWriter, r *http.Request) {
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
	menu, status := api.Service.GetMenuByID(r.Context(), id, all, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Menu", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   menu.ID,
		Type: "Menu",
		Attributes: entity.Menu{
			ID:           menu.ID,
			CategoryID:   menu.CategoryID,
			Name:         menu.Name,
			Price:        menu.Price,
			Discount:     menu.Discount,
			ShowMenu:     menu.ShowMenu,
			AppID:        menu.AppID,
			CreatedAt:    menu.CreatedAt,
			CreatedBy:    menu.CreatedBy,
			UpdatedAt:    menu.UpdatedAt,
			LastUpdateBy: menu.LastUpdateBy,
			DeletedAt:    menu.DeletedAt,
			IsActive:     menu.IsActive,
		},
	}
	responses.OK(w, res)
}

func (api *API) HandlePostMenu(w http.ResponseWriter, r *http.Request) {
	var (
		params reqMenu
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
	menu := &entity.Menu{
		CategoryID: params.CategoryID,
		Name:       params.Name,
		Price:      params.Price,
		Discount:   params.Discount,
	}
	menu, status := api.Service.InsertMenu(r.Context(), uid, menu)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Menu", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   menu.ID,
		Type: "Menu",
		Attributes: entity.Menu{
			ID:           menu.ID,
			CategoryID:   menu.CategoryID,
			Name:         menu.Name,
			Price:        menu.Price,
			Discount:     menu.Discount,
			ShowMenu:     menu.ShowMenu,
			AppID:        menu.AppID,
			CreatedAt:    menu.CreatedAt,
			CreatedBy:    menu.CreatedBy,
			UpdatedAt:    menu.UpdatedAt,
			LastUpdateBy: menu.LastUpdateBy,
			DeletedAt:    menu.DeletedAt,
			IsActive:     menu.IsActive,
		},
	}
	responses.OK(w, res)

}

func (api *API) HandlePatchMenu(w http.ResponseWriter, r *http.Request) {
	var (
		params reqMenu
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
	menu := &entity.Menu{
		ID:         id,
		CategoryID: params.CategoryID,
		Name:       params.Name,
		Price:      params.Price,
		Discount:   params.Discount,
	}
	menu, status := api.Service.UpdateMenu(r.Context(), isAdmin, uid, menu)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Menu", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   menu.ID,
		Type: "Menu",
		Attributes: entity.Menu{
			ID:           menu.ID,
			CategoryID:   menu.CategoryID,
			Name:         menu.Name,
			Price:        menu.Price,
			Discount:     menu.Discount,
			ShowMenu:     menu.ShowMenu,
			AppID:        menu.AppID,
			CreatedAt:    menu.CreatedAt,
			CreatedBy:    menu.CreatedBy,
			UpdatedAt:    menu.UpdatedAt,
			LastUpdateBy: menu.LastUpdateBy,
			DeletedAt:    menu.DeletedAt,
			IsActive:     menu.IsActive,
		},
	}
	responses.OK(w, res)

}
func (api *API) HandleDeleteMenu(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	uid, isAdmin := auth.IsAdmin(r)
	status := api.Service.DeleteMenu(r.Context(), id, isAdmin, uid)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Delete User", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
