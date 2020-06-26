package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/dbdialect"
)

// MySQL struct
type MySQL struct {
	db *sqlx.DB
}

// New init mysql
func New(db *sqlx.DB) *MySQL {
	return &MySQL{db}
}

func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, err error) {
	var (
		u    OrderDetails
		args []interface{}
	)
	query := `
	SELECT
		id,
		invoice_num,
		menu_id,
		amount,
		price,
		disc,
		sub_total,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		order_detail
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
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.SelectContext(ctx, &u, query, args...)
	if err != nil {
		return nil, err
	}
	for _, i := range u {
		orderDetails = append(orderDetails, entity.OrderDetail{
			ID:           i.ID,
			InvoiceNum:   i.InvoiceNum,
			MenuID:       i.MenuID,
			Amount:       i.Amount,
			Price:        i.Price,
			Discount:     i.Discount,
			SubTotal:     i.SubTotal,
			AppID:        i.AppID,
			CreatedAt:    i.CreatedAt,
			CreatedBy:    i.CreatedBy,
			UpdatedAt:    i.UpdatedAt,
			LastUpdateBy: i.LastUpdateBy,
			DeletedAt:    i.DeletedAt,
			IsActive:     i.IsActive,
		})
	}
	return orderDetails, nil
}

func (m *MySQL) SelectByInv(ctx context.Context, app int64, inv string, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, err error) {
	var (
		u    OrderDetails
		args []interface{}
	)
	query := `
	SELECT
		id,
		invoice_num,
		menu_id,
		amount,
		price,
		disc,
		sub_total,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		order_detail
	WHERE
		app_id = ? AND
		invoice_num = ?
		`
	args = append(args, app, inv)
	if !all {
		query += " AND is_active = 1"
	}
	if !isAdmin {
		query += "  AND created_by = ?"
		args = append(args, uid)

	}
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.SelectContext(ctx, &u, query, args...)
	if err != nil {
		return nil, err
	}
	for _, i := range u {
		orderDetails = append(orderDetails, entity.OrderDetail{
			ID:           i.ID,
			InvoiceNum:   i.InvoiceNum,
			MenuID:       i.MenuID,
			Amount:       i.Amount,
			Price:        i.Price,
			Discount:     i.Discount,
			SubTotal:     i.SubTotal,
			AppID:        i.AppID,
			CreatedAt:    i.CreatedAt,
			CreatedBy:    i.CreatedBy,
			UpdatedAt:    i.UpdatedAt,
			LastUpdateBy: i.LastUpdateBy,
			DeletedAt:    i.DeletedAt,
			IsActive:     i.IsActive,
		})
	}
	return orderDetails, nil
}

func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (orderDetail *entity.OrderDetail, err error) {
	var (
		i    OrderDetail
		args []interface{}
	)
	query := `
	SELECT
		id,
		invoice_num,
		menu_id,
		amount,
		price,
		disc,
		sub_total,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		order_detail
	WHERE
		app_id = ? AND
		id = ?
		`
	args = append(args, app, id)
	if !all {
		query += " AND is_active = 1"
	}
	if !isAdmin {
		query += "  AND created_by = ?"
		args = append(args, uid)

	}
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.GetContext(ctx, &i, query, args...)
	if err != nil {
		return nil, err
	}
	orderDetail = &entity.OrderDetail{
		ID:           i.ID,
		InvoiceNum:   i.InvoiceNum,
		MenuID:       i.MenuID,
		Amount:       i.Amount,
		Price:        i.Price,
		Discount:     i.Discount,
		SubTotal:     i.SubTotal,
		AppID:        i.AppID,
		CreatedAt:    i.CreatedAt,
		CreatedBy:    i.CreatedBy,
		UpdatedAt:    i.UpdatedAt,
		LastUpdateBy: i.LastUpdateBy,
		DeletedAt:    i.DeletedAt,
		IsActive:     i.IsActive,
	}
	return orderDetail, nil
}

// Insert
func (m *MySQL) Insert(ctx context.Context, orderDetail *entity.OrderDetail) (err error) {
	query := `
	INSERT INTO order_detail
		(
			invoice_num,
			menu_id,
			amount,
			price,
			disc,
			sub_total,
			app_id,
			created_at,
			created_by,
			is_active 
		) 
		VALUES 
		(
			:invoice_num,
			:menu_id,
			:amount,
			:price,
			:disc,
			:sub_total,
			:app_id,
			:created_at,
			:created_by,
			:is_active 
		);
	`
	_, err = m.db.NamedExecContext(ctx, query, &OrderDetail{
		ID:           orderDetail.ID,
		InvoiceNum:   orderDetail.InvoiceNum,
		MenuID:       orderDetail.MenuID,
		Amount:       orderDetail.Amount,
		Price:        orderDetail.Price,
		Discount:     orderDetail.Discount,
		SubTotal:     orderDetail.SubTotal,
		AppID:        orderDetail.AppID,
		CreatedAt:    orderDetail.CreatedAt,
		CreatedBy:    orderDetail.CreatedBy,
		UpdatedAt:    orderDetail.UpdatedAt,
		LastUpdateBy: orderDetail.LastUpdateBy,
		DeletedAt:    orderDetail.DeletedAt,
		IsActive:     orderDetail.IsActive,
	})
	if err != nil {
		return err
	}
	// orderDetail.ID, err = res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return err
}
func (m *MySQL) Update(ctx context.Context, isAdmin bool, orderDetail *entity.OrderDetail) (err error) {
	query := `
	UPDATE 
		order_detail
	SET
		menu_id = :menu_id,
		amount = :amount,
		price = :price,
		disc = :disc,
		sub_total = :sub_total,
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
	res, err := m.db.NamedExecContext(ctx, query, &OrderDetail{
		ID:           orderDetail.ID,
		InvoiceNum:   orderDetail.InvoiceNum,
		MenuID:       orderDetail.MenuID,
		Amount:       orderDetail.Amount,
		Price:        orderDetail.Price,
		Discount:     orderDetail.Discount,
		SubTotal:     orderDetail.SubTotal,
		AppID:        orderDetail.AppID,
		CreatedAt:    orderDetail.CreatedAt,
		CreatedBy:    orderDetail.CreatedBy,
		UpdatedAt:    orderDetail.UpdatedAt,
		LastUpdateBy: orderDetail.LastUpdateBy,
		DeletedAt:    orderDetail.DeletedAt,
		IsActive:     orderDetail.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, orderDetail *entity.OrderDetail) (err error) {
	query := `
	UPDATE 
		order_detail
	SET
		is_active = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	if !isAdmin {
		query += " AND created_by = :created_by"
	}
	res, err := m.db.NamedExecContext(ctx, query, &OrderDetail{
		ID:           orderDetail.ID,
		InvoiceNum:   orderDetail.InvoiceNum,
		MenuID:       orderDetail.MenuID,
		Amount:       orderDetail.Amount,
		Price:        orderDetail.Price,
		Discount:     orderDetail.Discount,
		SubTotal:     orderDetail.SubTotal,
		AppID:        orderDetail.AppID,
		CreatedAt:    orderDetail.CreatedAt,
		CreatedBy:    orderDetail.CreatedBy,
		UpdatedAt:    orderDetail.UpdatedAt,
		LastUpdateBy: orderDetail.LastUpdateBy,
		DeletedAt:    orderDetail.DeletedAt,
		IsActive:     orderDetail.IsActive,
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
