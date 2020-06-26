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

func (m *MySQL) Select(ctx context.Context, app int64, all bool) (categorys entity.Categorys, err error) {
	var (
		c Categorys
	)
	query := `
	SELECT
		id,
		category_name,
		app_id,
		is_active 
	FROM
		category 
	WHERE
		app_id = ?
		`
	if !all {
		query += " AND is_active = 1"
	}
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.SelectContext(ctx, &c, query, app)
	if err != nil {
		return nil, err
	}
	for _, i := range c {
		categorys = append(categorys, entity.Category{
			ID:           i.ID,
			CategoryName: i.CategoryName,
			AppID:        i.AppID,
			IsActive:     i.IsActive,
		})
	}
	return categorys, nil
}
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool) (category *entity.Category, err error) {
	var (
		p Category
	)
	query := `
	SELECT
		id,
		category_name,
		app_id,
		is_active 
	FROM
		category 
	WHERE
		app_id = ? AND
		id = ?
		`
	if !all {
		query += " AND is_active = 1"
	}
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.GetContext(ctx, &p, query, app, id)
	if err != nil {
		return nil, err
	}
	category = &entity.Category{
		ID:           p.ID,
		CategoryName: p.CategoryName,
		AppID:        p.AppID,
		IsActive:     p.IsActive,
	}
	return category, nil
}
func (m *MySQL) Insert(ctx context.Context, c *entity.Category) (err error) {
	query := `
	INSERT INTO category
		(
			category_name,
			app_id,
			is_active
		) 
		VALUES 
		(
			:category_name,
			:app_id,
			:is_active
		);
	`
	_, err = m.db.NamedExecContext(ctx, query, &Category{
		CategoryName: c.CategoryName,
		AppID:        c.AppID,
		IsActive:     c.IsActive,
	})
	if err != nil {
		return err
	}
	// c.ID, err = res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return err
}
func (m *MySQL) Update(ctx context.Context, c *entity.Category) (err error) {
	query := `
	UPDATE 
		category
	SET
		category_name =:category_name
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	res, err := m.db.NamedExecContext(ctx, query, &Category{
		ID:           c.ID,
		CategoryName: c.CategoryName,
		AppID:        c.AppID,
		IsActive:     c.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, c *entity.Category) (err error) {
	query := `
	UPDATE 
		category
	SET
		is_active =0
	WHERE 
		id = :id AND
		app_id = :app_id AND
		is_active = 1
	`
	res, err := m.db.NamedExecContext(ctx, query, &Category{
		ID:       c.ID,
		AppID:    c.AppID,
		IsActive: c.IsActive,
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
