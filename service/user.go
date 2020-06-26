package service

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"golang.org/x/crypto/bcrypt"
	null "gopkg.in/guregu/null.v3"
)

func (s *svc) SignIn(ctx context.Context, email, password string) (token *auth.Token, status Status) {
	user, err := s.user.GetByEmail(ctx, s.AppID, email)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, Status{http.StatusNotFound, "Email"}
		}
		return nil, Status{http.StatusNotFound, "Email"}
	}

	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil, Status{http.StatusUnprocessableEntity, "Incorrect Password"}

	}
	token, err = auth.CreateToken(user.UserHash, &user.Scope)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusUnprocessableEntity, "Generate Token"}
	}
	return token, Status{http.StatusOK, ""}
}
func (s *svc) GetUserByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (user *entity.User, status Status) {
	user, err := s.user.GetByID(ctx, s.AppID, id, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return user, Status{http.StatusNotFound, "User"}
	}
	if err != nil {
		log.Println(err)
		return user, Status{http.StatusInternalServerError, "User"}
	}
	return user, Status{http.StatusOK, ""}
}
func (s *svc) GetUserByHash(ctx context.Context, all bool, isAdmin bool, uid string) (user *entity.User, status Status) {
	user, err := s.user.GetByHash(ctx, s.AppID, all, isAdmin, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		return user, Status{http.StatusNotFound, "User"}
	}
	if err != nil {
		log.Println(err)
		return user, Status{http.StatusInternalServerError, "User"}
	}
	return user, Status{http.StatusOK, ""}
}
func (s *svc) SelectUsers(ctx context.Context, all bool, isAdmin bool, uid string) (users entity.Users, status Status) {
	user, err := s.user.Select(ctx, s.AppID, all, isAdmin, uid)
	if err != nil {
		log.Println(err)
		return user, Status{http.StatusInternalServerError, ""}
	}
	return user, Status{http.StatusOK, ""}
}
func (s *svc) InsertUser(ctx context.Context, uid string, user *entity.User) (userData *entity.User, status Status) {
	var (
		hashUser = HashSHA1(user.Username)
	)
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "User"}
	}
	password := string(hashedPassword)

	// Add
	user.Password = password
	user.CreatedAt = null.TimeFrom(time.Now())
	user.CreatedBy = uid
	user.UserHash = hashUser
	user.IsActive = 1
	user.AppID = s.AppID
	err = s.user.Insert(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, Status{http.StatusInternalServerError, "User"}
	}

	return user, Status{http.StatusOK, ""}
}
func (s *svc) UpdateUser(ctx context.Context, id int64, isAdmin bool, uid string, user *entity.User) (userData *entity.User, status Status) {
	getUser, status := s.GetUserByID(ctx, id, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return user, status
	}

	user.LastUpdateBy = &uid
	user.UpdatedAt = null.TimeFrom(time.Now())
	user.AppID = s.AppID
	user.ID = id
	user.CreatedBy = getUser.CreatedBy
	user.UserHash = getUser.UserHash
	err := s.user.Update(ctx, isAdmin, user)
	if err != nil {
		log.Println(err)
		return user, Status{http.StatusInternalServerError, "User"}
	}

	// kirim response
	userData, status = s.GetUserByID(ctx, id, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return user, status
	}
	return userData, Status{http.StatusOK, ""}
}
func (s *svc) DeleteUser(ctx context.Context, id int64, isAdmin bool, uid string) (status Status) {
	getUser, status := s.GetUserByID(ctx, s.AppID, false, isAdmin, uid)
	if status.Code != http.StatusOK {
		return status
	}

	user := &entity.User{
		ID:           id,
		LastUpdateBy: &uid,
		AppID:        s.AppID,
		DeletedAt:    null.TimeFrom(time.Now()),
		CreatedBy:    getUser.CreatedBy,
	}
	err := s.user.Delete(ctx, isAdmin, user)
	if err != nil {
		log.Println(err)
		return Status{http.StatusInternalServerError, "User"}
	}
	return Status{http.StatusOK, ""}
}

// VerifyPassword . . .
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HashSHA1 . .
func HashSHA1(text string) (hash string) {
	var sha = sha1.New()
	sha.Write([]byte(text))
	encrypted := sha.Sum(nil)
	hash = fmt.Sprintf("%x", encrypted)
	return
}

// Hash . . .
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
