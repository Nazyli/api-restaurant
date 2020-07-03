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

func (api *API) HandleSelectCategorys(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		all      = false
		err      error
	)
	_, isAdmin := auth.IsAdmin(r)
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

	categorys, status := api.Service.SelectCategory(r.Context(), all)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Categorys")
		return
	}
	// display array scope
	res := make([]DataResponse, 0, len(categorys))
	for _, p := range categorys {
		res = append(res, DataResponse{
			ID:   p.ID,
			Type: "Category",
			Attributes: entity.Category{
				ID:           p.ID,
				CategoryName: p.CategoryName,
				AppID:        p.AppID,
				IsActive:     p.IsActive,
			},
		})
	}

	responses.OK(w, res)
}

func (api *API) HandleGetCategoryById(w http.ResponseWriter, r *http.Request) {
	var (
		getParam = r.URL.Query()
		all      = false
	)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}

	_, isAdmin := auth.IsAdmin(r)
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
	p, status := api.Service.GetCategoryByID(r.Context(), id, all)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Category", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Category",
		Attributes: entity.Category{
			ID:           p.ID,
			CategoryName: p.CategoryName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) HandlePostCategorys(w http.ResponseWriter, r *http.Request) {
	var (
		params reqCategory
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
	position := &entity.Category{
		CategoryName: params.CategoryName,
	}
	p, status := api.Service.InsertCategory(r.Context(), position)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Category", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Category",
		Attributes: entity.Category{
			ID:           p.ID,
			CategoryName: p.CategoryName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) HandlePatchCategorys(w http.ResponseWriter, r *http.Request) {
	var (
		params   reqCategory
		paramsID = mux.Vars(r)
	)
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
	position := &entity.Category{
		ID:           id,
		CategoryName: params.CategoryName,
	}
	p, status := api.Service.UpdateCategory(r.Context(), position)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Category", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Category",
		Attributes: entity.Category{
			ID:           p.ID,
			CategoryName: p.CategoryName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) HandleDeleteCategorys(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	status := api.Service.DeleteCategory(r.Context(), &entity.Category{ID: id})
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Category", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
