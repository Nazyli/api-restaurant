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

func (api *API) handleSelectPositions(w http.ResponseWriter, r *http.Request) {
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

	positions, status := api.service.SelectPosition(r.Context(), all)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Positions")
		return
	}
	// display array scope
	res := make([]DataResponse, 0, len(positions))
	for _, p := range positions {
		res = append(res, DataResponse{
			ID:   p.ID,
			Type: "Position",
			Attributes: entity.Position{
				ID:           p.ID,
				PositionName: p.PositionName,
				AppID:        p.AppID,
				IsActive:     p.IsActive,
			},
		})
	}

	responses.OK(w, res)
}

func (api *API) handleGetPositionById(w http.ResponseWriter, r *http.Request) {
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
	p, status := api.service.GetPositionByID(r.Context(), id, all)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Get Position", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Position",
		Attributes: entity.Position{
			ID:           p.ID,
			PositionName: p.PositionName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) handlePostPositions(w http.ResponseWriter, r *http.Request) {
	var (
		params reqPosition
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
	position := &entity.Position{
		PositionName: params.PositionName,
	}
	p, status := api.service.InsertPosition(r.Context(), position)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Insert Position", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Position",
		Attributes: entity.Position{
			ID:           p.ID,
			PositionName: p.PositionName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) handlePatchPositions(w http.ResponseWriter, r *http.Request) {
	var (
		params   reqPosition
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
	position := &entity.Position{
		ID:           id,
		PositionName: params.PositionName,
	}
	p, status := api.service.UpdatePosition(r.Context(), position)
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Position", status.ErrMsg)
		return
	}
	res := DataResponse{
		ID:   p.ID,
		Type: "Position",
		Attributes: entity.Position{
			ID:           p.ID,
			PositionName: p.PositionName,
			AppID:        p.AppID,
			IsActive:     p.IsActive,
		},
	}
	responses.OK(w, res)
}
func (api *API) handleDeletePositions(w http.ResponseWriter, r *http.Request) {
	paramsID := mux.Vars(r)
	id, err := strconv.ParseInt(paramsID["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, "ID must Integer")
		return
	}
	status := api.service.DeletePosition(r.Context(), &entity.Position{ID: id})
	if status.Code != http.StatusOK {
		responses.ERROR(w, status.Code, "Failed Update Position", status.ErrMsg)
		return
	}
	responses.OK(w, "OK")
}
