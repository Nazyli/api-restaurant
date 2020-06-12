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

func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (order entity.Orders, err error) {
	var (
		u    Orders
		args []interface{}
	)
	query := `
	SELECT
		id,
		invoice_num,
		sale_date,
		sale_time,
		sub_total,
		tax,
		total,
		cash,
		change_money,
		other,
		payment_status,
		customer_id,
		employee_id,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		orders
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
		order = append(order, entity.Order{
			ID:            i.ID,
			InvoiceNum:    i.InvoiceNum,
			SaleDate:      i.SaleDate,
			SaleTime:      i.SaleTime,
			SubTotal:      i.SubTotal,
			Tax:           i.Tax,
			Total:         i.Total,
			Cash:          i.Cash,
			ChangeMoney:   i.ChangeMoney,
			Other:         i.Other,
			PaymentStatus: i.PaymentStatus,
			CustomerID:    i.CustomerID,
			EmployeeID:    i.EmployeeID,
			AppID:         i.AppID,
			CreatedAt:     i.CreatedAt,
			CreatedBy:     i.CreatedBy,
			UpdatedAt:     i.UpdatedAt,
			LastUpdateBy:  i.LastUpdateBy,
			DeletedAt:     i.DeletedAt,
			IsActive:      i.IsActive,
		})
	}
	return order, nil
}

// GetByID . . .
func (m *MySQL) GetByInv(ctx context.Context, app int64, inv string, all bool, isAdmin bool, uid string) (order *entity.Order, err error) {
	var (
		i    Order
		args []interface{}
	)
	query := `
	SELECT
		id,
		invoice_num,
		sale_date,
		sale_time,
		sub_total,
		tax,
		total,
		cash,
		change_money,
		other,
		payment_status,
		customer_id,
		employee_id,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		orders
	WHERE
		invoice_num = ? AND
		app_id = ?
		`
	args = append(args, inv, app)
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
	order = &entity.Order{
		ID:            i.ID,
		InvoiceNum:    i.InvoiceNum,
		SaleDate:      i.SaleDate,
		SaleTime:      i.SaleTime,
		SubTotal:      i.SubTotal,
		Tax:           i.Tax,
		Total:         i.Total,
		Cash:          i.Cash,
		ChangeMoney:   i.ChangeMoney,
		Other:         i.Other,
		PaymentStatus: i.PaymentStatus,
		CustomerID:    i.CustomerID,
		EmployeeID:    i.EmployeeID,
		AppID:         i.AppID,
		CreatedAt:     i.CreatedAt,
		CreatedBy:     i.CreatedBy,
		UpdatedAt:     i.UpdatedAt,
		LastUpdateBy:  i.LastUpdateBy,
		DeletedAt:     i.DeletedAt,
		IsActive:      i.IsActive,
	}
	return order, nil
}

func (m *MySQL) Insert(ctx context.Context, order *entity.Order) (err error) {
	query := `
	INSERT INTO orders
		(
			invoice_num,
			employee_id,
			app_id,
			created_at,
			created_by,
			is_active 
		) 
		VALUES 
		(
			:invoice_num,
			:employee_id,
			:app_id,
			:created_at,
			:created_by,
			:is_active 
		);
	`
	res, err := m.db.NamedExecContext(ctx, query, &Order{
		InvoiceNum: order.InvoiceNum,
		EmployeeID: order.EmployeeID,
		AppID:      order.AppID,
		CreatedAt:  order.CreatedAt,
		CreatedBy:  order.CreatedBy,
		IsActive:   order.IsActive,
	})
	if err != nil {
		return err
	}
	order.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (m *MySQL) Update(ctx context.Context, isAdmin bool, order *entity.Order) (err error) {
	query := `
	UPDATE 
		orders
	SET
		sale_date = :sale_date,
		sale_time = :sale_time,
		sub_total = :sub_total,
		tax = :tax,
		total = :total,
		cash = :cash,
		change_money = :change_money,
		other = :other,
		payment_status = :payment_status,
		customer_id = :customer_id,
		employee_id = :employee_id,
		updated_at = :updated_at,
		last_update_by = :last_update_by
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	if !isAdmin {
		query += " AND created_by = :created_by"
	}
	res, err := m.db.NamedExecContext(ctx, query, &Order{
		ID:            order.ID,
		InvoiceNum:    order.InvoiceNum,
		SaleDate:      order.SaleDate,
		SaleTime:      order.SaleTime,
		SubTotal:      order.SubTotal,
		Tax:           order.Tax,
		Total:         order.Total,
		Cash:          order.Cash,
		ChangeMoney:   order.ChangeMoney,
		Other:         order.Other,
		PaymentStatus: order.PaymentStatus,
		CustomerID:    order.CustomerID,
		EmployeeID:    order.EmployeeID,
		AppID:         order.AppID,
		CreatedAt:     order.CreatedAt,
		CreatedBy:     order.CreatedBy,
		UpdatedAt:     order.UpdatedAt,
		LastUpdateBy:  order.LastUpdateBy,
		DeletedAt:     order.DeletedAt,
		IsActive:      order.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, order *entity.Order) (err error) {
	query := `
	UPDATE 
		orders
	SET
		is_active = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE 
		invoice_num = :invoice_num AND
		is_active = 1 AND
		app_id = :app_id AND
		(payment_status IS NULL OR payment_status != 1)
	`
	if !isAdmin {
		query += " AND created_by = :created_by"
	}
	res, err := m.db.NamedExecContext(ctx, query, &Order{
		InvoiceNum:   order.InvoiceNum,
		AppID:        order.AppID,
		LastUpdateBy: order.LastUpdateBy,
		DeletedAt:    order.DeletedAt,
		IsActive:     order.IsActive,
		CreatedBy:    order.CreatedBy,
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
