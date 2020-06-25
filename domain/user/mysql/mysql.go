package mysql

import (
	"context"
	"database/sql"

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
func (m *MySQL) GetByEmail(ctx context.Context, app int64, email string) (user *entity.User, err error) {
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

// GetByID . . .
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (user *entity.User, err error) {
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
	if !isAdmin {
		query += " AND (created_by = ? OR user_hash = ?)"
		args = append(args, uid, uid)

	}
	err = m.db.GetContext(ctx, &u, query, args...)
	if err != nil {
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

// GetByID . . .
func (m *MySQL) GetByHash(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (user *entity.User, err error) {
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
		user_hash = ? AND
		app_id = ?
		`
	args = append(args, uid, app)
	if !all {
		query += " AND is_active = 1"
	}
	if !isAdmin {
		query += " AND (created_by = ? OR user_hash = ?)"
		args = append(args, uid, uid)

	}
	err = m.db.GetContext(ctx, &u, query, args...)
	if err != nil {
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

// Select . . .
func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (users entity.Users, err error) {
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
	if !isAdmin {
		query += "  AND (created_by = ? OR user_hash = ?)"
		args = append(args, uid, uid)

	}
	err = m.db.SelectContext(ctx, &u, query, args...)
	if err != nil {
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

func (m *MySQL) Insert(ctx context.Context, user *entity.User) (err error) {
	query := `
	INSERT INTO user
		(
			username,
			email,
			password,
			user_hash,
			employee_id,
			scope,
			app_id,
			created_at,
			created_by,
			is_active
		) 
		VALUES 
		(
			:username,
			:email,
			:password,
			:user_hash,
			:employee_id,
			:scope,
			:app_id,
			:created_at,
			:created_by,
			:is_active
		);
	`
	res, err := m.db.NamedExecContext(ctx, query, &User{
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		UserHash:   user.UserHash,
		EmployeeID: user.EmployeeID,
		Scope:      user.Scope,
		AppID:      user.AppID,
		CreatedAt:  user.CreatedAt,
		CreatedBy:  user.CreatedBy,
		IsActive:   user.IsActive,
	})
	if err != nil {
		return err
	}
	user.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (m *MySQL) Update(ctx context.Context, isAdmin bool, user *entity.User) (err error) {
	query := `
	UPDATE 
		user
	SET
		username =:username,
		email =:email,
		employee_id =:employee_id,
		scope =:scope,
		app_id =:app_id,
		updated_at =:updated_at,
		last_update_by =:last_update_by
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	if !isAdmin {
		query += " AND (created_by = :created_by OR user_hash = :user_hash)"

	}
	res, err := m.db.NamedExecContext(ctx, query, &User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		EmployeeID:   user.EmployeeID,
		Scope:        user.Scope,
		AppID:        user.AppID,
		CreatedBy:    user.CreatedBy,
		UserHash:     user.UserHash,
		UpdatedAt:    user.UpdatedAt,
		LastUpdateBy: user.LastUpdateBy,
		IsActive:     user.IsActive,
	})
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, user *entity.User) (err error) {
	query := `
	UPDATE
		user
	SET
		is_active = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE
		id = :id AND
		is_active = 1 AND
		app_id = :app_id`

	if !isAdmin {
		query += " AND created_by = :created_by"
	}
	res, err := m.db.NamedExecContext(ctx, query, &User{
		ID:           user.ID,
		AppID:        user.AppID,
		DeletedAt:    user.DeletedAt,
		LastUpdateBy: user.LastUpdateBy,
		CreatedBy:    user.CreatedBy,
	})
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return sql.ErrNoRows
	}
	return err
}
