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
func (m *MySQL) GetByEmail(ctx context.Context, email string, app int64) (user *entity.User, err error) {
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
		email = ? AND
		app_id = ?
		`
	err = m.db.GetContext(ctx, &u, query, email, app)
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
		AppID:      u.AppID,
	}
	return user, nil
}

// GetByEmail . . .
func (m *MySQL) GetByID(ctx context.Context, all bool, uid string, id int64, app int64) (user *entity.User, err error) {
	var (
		u    User
		args []interface{}
	)
	query := `
	SELECT
		id,
		username,
		email,
		user_hash,
		employee_id,
		scope,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		user
	WHERE
		id = ? AND
		app_id = ?
		`
	args = append(args, id, app)
	if !all {
		query += " AND is_active = 1"
	}
	if uid != "" {
		query += " AND created_by = ?"
		args = append(args, uid)

	}
	err = m.db.GetContext(ctx, &u, query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user = &entity.User{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		UserHash:     u.UserHash,
		EmployeeID:   u.EmployeeID,
		Scope:        u.Scope,
		CreatedAt:    u.CreatedAt,
		CreatedBy:    u.CreatedBy,
		UpdatedAt:    u.UpdatedAt,
		LastUpdateBy: u.LastUpdateBy,
		DeletedAt:    u.DeletedAt,
		IsActive:     u.IsActive,
		AppID:        u.AppID,
	}
	return user, nil
}
func (m *MySQL) Select(ctx context.Context, all bool, uid string, app int64) (users entity.Users, err error) {
	var (
		u    Users
		args []interface{}
	)
	query := `
	SELECT
		id,
		username,
		email,
		user_hash,
		employee_id,
		scope,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		user
	WHERE
		app_id = ?
		`
	args = append(args, app)
	if !all {
		query += " AND is_active = 1"
	}
	if uid != "" {
		query += " AND created_by = ?"
		args = append(args, uid)

	}
	err = m.db.SelectContext(ctx, &u, query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, i := range u {
		users = append(users, entity.User{
			ID:           i.ID,
			Username:     i.Username,
			Email:        i.Email,
			UserHash:     i.UserHash,
			EmployeeID:   i.EmployeeID,
			Scope:        i.Scope,
			CreatedAt:    i.CreatedAt,
			CreatedBy:    i.CreatedBy,
			UpdatedAt:    i.UpdatedAt,
			LastUpdateBy: i.LastUpdateBy,
			DeletedAt:    i.DeletedAt,
			IsActive:     i.IsActive,
			AppID:        i.AppID,
		})
	}
	return users, nil
}
