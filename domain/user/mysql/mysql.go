package mysql

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nazyli/api-restaurant/entity"
)

// MySQL struct
type MySQL struct {
	db *sqlx.DB
}

// New init mysql
func New(db *sqlx.DB) *MySQL {
	return &MySQL{db}
}

// GetByEmail . . .
func (m *MySQL) GetByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	var u User
	query := `
	SELECT
		id,
		username,
		email,
		password,
		user_hash,
		employee_id,
		scope 
	FROM
		user
	WHERE
		is_active = 1 AND
		email = ?;
		`
	err = m.db.GetContext(ctx, &u, query, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user = &entity.User{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		Password:   u.Password,
		UserHash:   u.UserHash,
		EmployeeID: u.EmployeeID,
		Scope:      u.Scope,
		IsActive:   u.IsActive,
	}
	return user, nil
}

// GetByEmail . . .
func (m *MySQL) GetByID(ctx context.Context, id int64) (user *entity.User, err error) {
	var u User
	query := `
	SELECT
		*
	FROM
		users
	WHERE
		id = ?;
		`
	err = m.db.GetContext(ctx, &u, query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user = &entity.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Scope:     u.Scope,
	}
	return user, nil
}
