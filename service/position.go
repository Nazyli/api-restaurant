package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/nazyli/api-restaurant/entity"
)

func (s *svc) SelectPosition(ctx context.Context, all bool) (positions entity.Positions, status Status) {
	positions, err := s.position.Select(ctx, s.AppID, all)
	if err != nil {
		log.Println(err)
		return positions, Status{http.StatusInternalServerError, ""}
	}
	return positions, Status{http.StatusOK, ""}
}

func (s *svc) GetPositionByID(ctx context.Context, id int64, all bool) (position *entity.Position, status Status) {
	position, err := s.position.GetByID(ctx, s.AppID, id, all)
	if err == sql.ErrNoRows {
		log.Println(err)
		return position, Status{http.StatusNotFound, "Position"}
	}
	if err != nil {
		log.Println(err)
		return position, Status{http.StatusInternalServerError, "Position"}
	}
	return position, Status{http.StatusOK, ""}
}

func (s *svc) InsertPosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status) {
	// Add
	position.IsActive = 1
	position.AppID = s.AppID
	err := s.position.Insert(ctx, position)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "Position"}
	}

	return position, Status{http.StatusOK, ""}
}

func (s *svc) UpdatePosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status) {
	_, status = s.GetPositionByID(ctx, position.ID, false)
	if status.Code != http.StatusOK {
		return position, status
	}
	position.AppID = s.AppID
	err := s.position.Update(ctx, position)
	if err != nil {
		log.Println(err)
		return position, Status{http.StatusInternalServerError, "Position"}
	}

	// kirim response
	positionData, status = s.GetPositionByID(ctx, position.ID, false)
	if status.Code != http.StatusOK {
		return position, status
	}
	return positionData, Status{http.StatusOK, ""}
}

func (s *svc) DeletePosition(ctx context.Context, position *entity.Position) (status Status) {
	position.AppID = s.AppID
	_, status = s.GetPositionByID(ctx, position.ID, false)
	if status.Code != http.StatusOK {
		return status
	}
	err := s.position.Delete(ctx, position)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "Position"}
	}
	return Status{http.StatusOK, ""}
}
