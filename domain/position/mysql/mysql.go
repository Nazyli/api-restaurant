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

func (m *MySQL) Select(ctx context.Context, app int64, all bool) (positions entity.Positions, err error) {
	var (
		p Positions
	)
	query := `
	SELECT
		id,
		position_name,
		app_id,
		is_active 
	FROM
		position 
	WHERE
		app_id = ?
		`
	if !all {
		query += " AND is_active = 1"
	}
	query = dbdialect.New(m.db).SetQuery(query)
	err = m.db.SelectContext(ctx, &p, query, app)
	if err != nil {
		return nil, err
	}
	for _, i := range p {
		positions = append(positions, entity.Position{
			ID:           i.ID,
			PositionName: i.PositionName,
			AppID:        i.AppID,
			IsActive:     i.IsActive,
		})
	}
	return positions, nil
}
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool) (position *entity.Position, err error) {
	var (
		p Position
	)
	query := `
	SELECT
		id,
		position_name,
		app_id,
		is_active 
	FROM
		position 
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
	position = &entity.Position{
		ID:           p.ID,
		PositionName: p.PositionName,
		AppID:        p.AppID,
		IsActive:     p.IsActive,
	}
	return position, nil
}
func (m *MySQL) Insert(ctx context.Context, position *entity.Position) (err error) {
	query := `
	INSERT INTO position
		(
			position_name,
			app_id,
			is_active
		) 
		VALUES 
		(
			:position_name,
			:app_id,
			:is_active
		);
	`
	_, err = m.db.NamedExecContext(ctx, query, &Position{
		PositionName: position.PositionName,
		AppID:        position.AppID,
		IsActive:     position.IsActive,
	})
	if err != nil {
		return err
	}
	// position.ID, err = res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return err
}
func (m *MySQL) Update(ctx context.Context, position *entity.Position) (err error) {
	query := `
	UPDATE 
		position
	SET
		position_name =:position_name
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	res, err := m.db.NamedExecContext(ctx, query, &Position{
		ID:           position.ID,
		PositionName: position.PositionName,
		AppID:        position.AppID,
		IsActive:     position.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, position *entity.Position) (err error) {
	query := `
	UPDATE 
		position
	SET
		is_active =0
	WHERE 
		id = :id AND
		app_id = :app_id AND
		is_active = 1
	`
	res, err := m.db.NamedExecContext(ctx, query, &Position{
		ID:       position.ID,
		AppID:    position.AppID,
		IsActive: position.IsActive,
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
