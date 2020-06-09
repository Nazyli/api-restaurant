package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *svc) SignIn(ctx context.Context, email, password string, app int64) (token *auth.Token, errMsg string, status int) {
	user, err := s.user.GetByEmail(ctx, email, app)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, "Email Not Found", http.StatusNotFound
		}
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	err = VerifyPassword(*user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil, "Incorrect Password", http.StatusUnprocessableEntity
	}
	token, err = auth.CreateToken(user.ID, user.Scope)
	if err != nil {
		log.Println(err)
		return nil, "Generate Token", http.StatusUnprocessableEntity
	}
	return token, "", http.StatusOK
}
func (s *svc) GetUserByID(ctx context.Context, all bool, uid string, id int64, app int64) (user *entity.User, status int) {
	user, err := s.user.GetByID(ctx, all, uid, id, app)
	if err == sql.ErrNoRows {
		log.Println(err)
		return user, http.StatusNotFound
	}
	if err != nil {
		log.Println(err)
		return user, http.StatusInternalServerError
	}
	return user, http.StatusOK
}
func (s *svc) SelectUsers(ctx context.Context, all bool, uid string, app int64) (users entity.Users, status int) {
	user, err := s.user.Select(ctx, all, uid, app)
	if err != nil {
		log.Println(err)
		return user, http.StatusInternalServerError
	}
	return user, http.StatusOK
}

// func
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
