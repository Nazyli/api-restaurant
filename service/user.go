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

func (s *svc) SignIn(ctx context.Context, email, password string) (token *auth.Token, errMsg string, status int) {
	user, err := s.user.GetByEmail(ctx, email)
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
func (s *svc) GetByID(ctx context.Context, id int64) (user *entity.User, err int) {
	user, errs := s.user.GetByID(ctx, id)
	if errs == sql.ErrNoRows {
		log.Println(err)
		return user, http.StatusNotFound
	}
	if errs != nil {
		log.Println(err)
		return user, http.StatusInternalServerError
	}
	return user, http.StatusOK
}

// func
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
