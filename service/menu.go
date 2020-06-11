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

func (s *svc) SelectMenues(ctx context.Context, all bool, isAdmin bool, uid string) (menues entity.Menues, status Status) {
	menu, err := s.menu.Select(ctx, s.AppID, all, isAdmin, uid)
	if err != nil {
		log.Println(err)
		return menu, Status{http.StatusInternalServerError, ""}
	}
	return menu, Status{http.StatusOK, ""}
}

func (s *svc) GetMenuByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (menu *entity.Menu, status Status) {
	menu, err := s.menu.GetByID(ctx, s.AppID, id, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return menu, Status{http.StatusNotFound, "Menu"}
	}
	if err != nil {
		log.Println(err)
		return menu, Status{http.StatusInternalServerError, "Menu"}
	}
	return menu, Status{http.StatusOK, ""}
}
func (s *svc) InsertMenu(ctx context.Context, menu *entity.Menu) (menuData *entity.Menu, status Status) {
	// Add
	menu.CreatedAt = null.TimeFrom(time.Now())
	menu.IsActive = 1
	menu.AppID = s.AppID
	menu.Discount = nil
	menu.ShowMenu = 1
	err := s.menu.Insert(ctx, menu)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Menu"}
	}

	return menu, Status{http.StatusOK, ""}
}

func (s *svc) UpdateMenu(ctx context.Context, isAdmin bool, uid string, menu *entity.Menu) (menuData *entity.Menu, status Status) {
	getMenu, status := s.GetMenuByID(ctx, menu.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return menu, status
	}
	menu.UpdatedAt = null.TimeFrom(time.Now())
	menu.AppID = s.AppID
	menu.CreatedBy = getMenu.CreatedBy
	menu.ShowMenu = getMenu.ShowMenu
	menu.Discount = nil
	err := s.menu.Update(ctx, isAdmin, menu)
	if err != nil {
		log.Println(err)
		return menu, Status{http.StatusInternalServerError, "Menu"}
	}

	// kirim response
	menuData, status = s.GetMenuByID(ctx, menu.ID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return menu, status
	}
	return menuData, Status{http.StatusOK, ""}
}
func (s *svc) DeleteMenu(ctx context.Context, id int64, isAdmin bool, uid string) (status Status) {
	getMenu, status := s.GetMenuByID(ctx, s.AppID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return status
	}

	menu := &entity.Menu{
		ID:           id,
		LastUpdateBy: &uid,
		AppID:        s.AppID,
		DeletedAt:    null.TimeFrom(time.Now()),
		CreatedBy:    getMenu.CreatedBy,
	}
	err := s.menu.Delete(ctx, isAdmin, menu)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "Menu"}
	}
	return Status{http.StatusOK, ""}
}
