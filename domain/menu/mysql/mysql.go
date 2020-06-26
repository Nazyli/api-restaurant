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

// GetByID . . .
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (menu *entity.Menu, err error) {
	var (
		u    Menu
		args []interface{}
	)
	query := `
	SELECT
		id,
		category_id,
		name,
		price,
		disc,
		show_menu,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		menu
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
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.GetContext(ctx, &u, query, args...)
	if err != nil {
		return nil, err
	}
	menu = &entity.Menu{
		ID:           u.ID,
		CategoryID:   u.CategoryID,
		Name:         u.Name,
		Price:        u.Price,
		Discount:     u.Discount,
		ShowMenu:     u.ShowMenu,
		AppID:        u.AppID,
		CreatedAt:    u.CreatedAt,
		CreatedBy:    u.CreatedBy,
		UpdatedAt:    u.UpdatedAt,
		LastUpdateBy: u.LastUpdateBy,
		DeletedAt:    u.DeletedAt,
		IsActive:     u.IsActive,
	}
	return menu, nil
}

// Select . . .
func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (menues entity.Menues, err error) {
	var (
		i    Menues
		args []interface{}
	)
	query := `
	SELECT
		id,
		category_id,
		name,
		price,
		disc,
		show_menu,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		is_active
	FROM
		menu
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
	err = m.db.SelectContext(ctx, &i, query, args...)
	if err != nil {
		return nil, err
	}
	for _, u := range i {
		menues = append(menues, entity.Menu{
			ID:           u.ID,
			CategoryID:   u.CategoryID,
			Name:         u.Name,
			Price:        u.Price,
			Discount:     u.Discount,
			ShowMenu:     u.ShowMenu,
			AppID:        u.AppID,
			CreatedAt:    u.CreatedAt,
			CreatedBy:    u.CreatedBy,
			UpdatedAt:    u.UpdatedAt,
			LastUpdateBy: u.LastUpdateBy,
			DeletedAt:    u.DeletedAt,
			IsActive:     u.IsActive,
		})
	}
	return menues, nil
}

func (m *MySQL) Insert(ctx context.Context, menu *entity.Menu) (err error) {
	query := `
	INSERT INTO menu
		(
			category_id,
			name,
			price,
			disc,
			show_menu,
			app_id,
			created_at,
			created_by,
			is_active
		) 
		VALUES 
		(
			:category_id,
			:name,
			:price,
			:disc,
			:show_menu,
			:app_id,
			:created_at,
			:created_by,
			:is_active
		);
	`
	_, err = m.db.NamedExecContext(ctx, query, &Menu{
		ID:           menu.ID,
		CategoryID:   menu.CategoryID,
		Name:         menu.Name,
		Price:        menu.Price,
		Discount:     menu.Discount,
		ShowMenu:     menu.ShowMenu,
		AppID:        menu.AppID,
		CreatedAt:    menu.CreatedAt,
		CreatedBy:    menu.CreatedBy,
		UpdatedAt:    menu.UpdatedAt,
		LastUpdateBy: menu.LastUpdateBy,
		DeletedAt:    menu.DeletedAt,
		IsActive:     menu.IsActive,
	})
	if err != nil {
		return err
	}
	// menu.ID, err = res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return err
}

func (m *MySQL) Update(ctx context.Context, isAdmin bool, menu *entity.Menu) (err error) {
	query := `
	UPDATE 
		menu
	SET
		category_id = :category_id,
		name = :name,
		price = :price,
		disc = :disc,
		show_menu = :show_menu,
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
	res, err := m.db.NamedExecContext(ctx, query, &Menu{
		ID:           menu.ID,
		CategoryID:   menu.CategoryID,
		Name:         menu.Name,
		Price:        menu.Price,
		Discount:     menu.Discount,
		ShowMenu:     menu.ShowMenu,
		AppID:        menu.AppID,
		CreatedAt:    menu.CreatedAt,
		CreatedBy:    menu.CreatedBy,
		UpdatedAt:    menu.UpdatedAt,
		LastUpdateBy: menu.LastUpdateBy,
		DeletedAt:    menu.DeletedAt,
		IsActive:     menu.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, menu *entity.Menu) (err error) {
	query := `
	UPDATE
		menu
	SET
		is_active = 0,
		show_menu = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE
		id = :id AND
		app_id = :app_id AND is_active = 1`

	if !isAdmin {
		query += " AND created_by = :created_by"
	}
	res, err := m.db.NamedExecContext(ctx, query, &Menu{
		ID:           menu.ID,
		AppID:        menu.AppID,
		DeletedAt:    menu.DeletedAt,
		LastUpdateBy: menu.LastUpdateBy,
		CreatedBy:    menu.CreatedBy,
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
