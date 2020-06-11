package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/nazyli/api-restaurant/entity"
)

func (s *svc) SelectCategory(ctx context.Context, all bool) (categorys entity.Categorys, status Status) {
	categorys, err := s.category.Select(ctx, s.AppID, all)
	if err != nil {
		log.Println(err)
		return categorys, Status{http.StatusInternalServerError, ""}
	}
	return categorys, Status{http.StatusOK, ""}
}

func (s *svc) GetCategoryByID(ctx context.Context, id int64, all bool) (category *entity.Category, status Status) {
	category, err := s.category.GetByID(ctx, s.AppID, id, all)
	if err == sql.ErrNoRows {
		log.Println(err)
		return category, Status{http.StatusNotFound, "Category"}
	}
	if err != nil {
		log.Println(err)
		return category, Status{http.StatusInternalServerError, "Category"}
	}
	return category, Status{http.StatusOK, ""}
}

func (s *svc) InsertCategory(ctx context.Context, category *entity.Category) (categoryData *entity.Category, status Status) {
	// Add
	category.IsActive = 1
	category.AppID = s.AppID
	err := s.category.Insert(ctx, category)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Category"}
	}

	return category, Status{http.StatusOK, ""}
}

func (s *svc) UpdateCategory(ctx context.Context, category *entity.Category) (categoryData *entity.Category, status Status) {
	_, status = s.GetCategoryByID(ctx, category.ID, false)
	if status.Code != http.StatusOK {
		return category, status
	}
	category.AppID = s.AppID
	err := s.category.Update(ctx, category)
	if err != nil {
		log.Println(err)
		return category, Status{http.StatusInternalServerError, "Category"}
	}

	// kirim response
	categoryData, status = s.GetCategoryByID(ctx, category.ID, false)
	if status.Code != http.StatusOK {
		return category, status
	}
	return categoryData, Status{http.StatusOK, ""}
}

func (s *svc) DeleteCategory(ctx context.Context, category *entity.Category) (status Status) {
	category.AppID = s.AppID
	_, status = s.GetCategoryByID(ctx, category.ID, false)
	if status.Code != http.StatusOK {
		return status
	}
	err := s.category.Delete(ctx, category)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "Category"}
	}
	return Status{http.StatusOK, ""}
}
