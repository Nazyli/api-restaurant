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

// Select . . .
func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (customers entity.Customers, err error) {
	var (
		u    Customers
		args []interface{}
	)
	query := `
	SELECT
		id,
		name,
		email,
		address,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		customer
	WHERE
		app_id = ?
		`
	args = append(args, app)
	if !all {
		query += " AND is_active = 1"
	}
	if !isAdmin {
		query += "  AND created_by = ?"
		args = append(args, uid)

	}
	err = m.db.SelectContext(ctx, &u, query, args...)
	if err != nil {
		return nil, err
	}
	for _, i := range u {
		customers = append(customers, entity.Customer{
			ID:           i.ID,
			Name:         i.Name,
			Email:        i.Email,
			Addreas:      i.Addreas,
			AppID:        i.AppID,
			CreatedAt:    i.CreatedAt,
			CreatedBy:    i.CreatedBy,
			UpdatedAt:    i.UpdatedAt,
			LastUpdateBy: i.LastUpdateBy,
			DeletedAt:    i.DeletedAt,
			IsActive:     i.IsActive,
		})
	}
	return customers, nil
}

// GetByID . . .
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (customer *entity.Customer, err error) {
	var (
		i    Customer
		args []interface{}
	)
	query := `
	SELECT
		id,
		name,
		email,
		address,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		customer
	WHERE
		id = ? AND
		app_id = ?
		`
	args = append(args, id, app)
	if !all {
		query += " AND is_active = 1"
	}
	if !isAdmin {
		query += " AND created_by = ?"
		args = append(args, uid)

	}
	err = m.db.GetContext(ctx, &i, query, args...)
	if err != nil {
		return nil, err
	}
	customer = &entity.Customer{
		ID:           i.ID,
		Name:         i.Name,
		Email:        i.Email,
		Addreas:      i.Addreas,
		AppID:        i.AppID,
		CreatedAt:    i.CreatedAt,
		CreatedBy:    i.CreatedBy,
		UpdatedAt:    i.UpdatedAt,
		LastUpdateBy: i.LastUpdateBy,
		DeletedAt:    i.DeletedAt,
		IsActive:     i.IsActive,
	}
	return customer, nil
}

func (m *MySQL) Insert(ctx context.Context, customer *entity.Customer) (err error) {
	query := `
	INSERT INTO customer
		(
			id,
			name,
			email,
			address,
			app_id,
			created_at,
			created_by,
			is_active
		) 
		VALUES 
		(
			:id,
			:name,
			:email,
			:address,
			:app_id,
			:created_at,
			:created_by,
			:is_active
		);
	`
	res, err := m.db.NamedExecContext(ctx, query, &Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		Addreas:      customer.Addreas,
		AppID:        customer.AppID,
		CreatedAt:    customer.CreatedAt,
		CreatedBy:    customer.CreatedBy,
		UpdatedAt:    customer.UpdatedAt,
		LastUpdateBy: customer.LastUpdateBy,
		DeletedAt:    customer.DeletedAt,
		IsActive:     customer.IsActive,
	})
	if err != nil {
		return err
	}
	customer.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (m *MySQL) Update(ctx context.Context, isAdmin bool, customer *entity.Customer) (err error) {
	query := `
	UPDATE 
		customer
	SET
		name =:name,
		email =:email,
		address =:address,
		app_id =:app_id,
		updated_at =:updated_at,
		last_update_by =:last_update_by
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	if !isAdmin {
		query += " AND created_by = :created_by"

	}
	res, err := m.db.NamedExecContext(ctx, query, &Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		Addreas:      customer.Addreas,
		AppID:        customer.AppID,
		CreatedAt:    customer.CreatedAt,
		CreatedBy:    customer.CreatedBy,
		UpdatedAt:    customer.UpdatedAt,
		LastUpdateBy: customer.LastUpdateBy,
		DeletedAt:    customer.DeletedAt,
		IsActive:     customer.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, customer *entity.Customer) (err error) {
	query := `
	UPDATE 
		customer
	SET
		is_active = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE
		id = :id AND
		app_id = :app_id
		`
	if !isAdmin {
		query += " AND created_by = :created_by"

	}
	res, err := m.db.NamedExecContext(ctx, query, &Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		Addreas:      customer.Addreas,
		AppID:        customer.AppID,
		CreatedAt:    customer.CreatedAt,
		CreatedBy:    customer.CreatedBy,
		UpdatedAt:    customer.UpdatedAt,
		LastUpdateBy: customer.LastUpdateBy,
		DeletedAt:    customer.DeletedAt,
		IsActive:     customer.IsActive,
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
