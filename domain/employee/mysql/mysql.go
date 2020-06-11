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
func (m *MySQL) Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (employee entity.Employees, err error) {
	var (
		u    Employees
		args []interface{}
	)
	query := `
	SELECT
		id,
		position_id,
		name,
		date_of_birth,
		address,
		gender,
		email,
		salary,
		bonus,
		from_date,
		finish_date,
		show_employee,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		image_url,
		image_id,
		is_active 
	FROM
		employee 
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
		employee = append(employee, entity.Employee{
			ID:           i.ID,
			PositionID:   i.PositionID,
			Name:         i.Name,
			DateOfBirth:  i.DateOfBirth,
			Address:      i.Address,
			Gender:       i.Gender,
			Email:        i.Email,
			Salary:       i.Salary,
			Bonus:        i.Bonus,
			FromDate:     i.FromDate,
			FinishDate:   i.FinishDate,
			ShowEmployee: i.ShowEmployee,
			AppID:        i.AppID,
			CreatedAt:    i.CreatedAt,
			CreatedBy:    i.CreatedBy,
			UpdatedAt:    i.UpdatedAt,
			LastUpdateBy: i.LastUpdateBy,
			DeletedAt:    i.DeletedAt,
			ImageUrl:     i.ImageUrl,
			ImageID:      i.ImageID,
			IsActive:     i.IsActive,
		})
	}
	return employee, nil
}

// GetByID . . .
func (m *MySQL) GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (employee *entity.Employee, err error) {
	var (
		i    Employee
		args []interface{}
	)
	query := `
	SELECT
		id,
		position_id,
		name,
		date_of_birth,
		address,
		gender,
		email,
		salary,
		bonus,
		from_date,
		finish_date,
		show_employee,
		app_id,
		created_at,
		created_by,
		updated_at,
		last_update_by,
		deleted_at,
		image_url,
		image_id,
		is_active 
	FROM
		employee
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
	employee = &entity.Employee{
		ID:           i.ID,
		PositionID:   i.PositionID,
		Name:         i.Name,
		DateOfBirth:  i.DateOfBirth,
		Address:      i.Address,
		Gender:       i.Gender,
		Email:        i.Email,
		Salary:       i.Salary,
		Bonus:        i.Bonus,
		FromDate:     i.FromDate,
		FinishDate:   i.FinishDate,
		ShowEmployee: i.ShowEmployee,
		AppID:        i.AppID,
		CreatedAt:    i.CreatedAt,
		CreatedBy:    i.CreatedBy,
		UpdatedAt:    i.UpdatedAt,
		LastUpdateBy: i.LastUpdateBy,
		DeletedAt:    i.DeletedAt,
		ImageUrl:     i.ImageUrl,
		ImageID:      i.ImageID,
		IsActive:     i.IsActive,
	}
	return employee, nil
}

func (m *MySQL) Insert(ctx context.Context, employee *entity.Employee) (err error) {
	query := `
	INSERT INTO employee
		(
			id,
			position_id,
			name,
			date_of_birth,
			address,
			gender,
			email,
			salary,
			bonus,
			from_date,
			finish_date,
			show_employee,
			app_id,
			created_at,
			created_by,
			image_url,
			image_id,
			is_active 
		) 
		VALUES 
		(
			:position_id,
			:name,
			:id,
			:date_of_birth,
			:address,
			:gender,
			:email,
			:salary,
			:bonus,
			:from_date,
			:finish_date,
			:show_employee,
			:app_id,
			:created_at,
			:created_by,
			:image_url,
			:image_id,
			:is_active 
		);
	`
	res, err := m.db.NamedExecContext(ctx, query, &Employee{
		ID:           employee.ID,
		PositionID:   employee.PositionID,
		Name:         employee.Name,
		DateOfBirth:  employee.DateOfBirth,
		Address:      employee.Address,
		Gender:       employee.Gender,
		Email:        employee.Email,
		Salary:       employee.Salary,
		Bonus:        employee.Bonus,
		FromDate:     employee.FromDate,
		FinishDate:   employee.FinishDate,
		ShowEmployee: employee.ShowEmployee,
		AppID:        employee.AppID,
		CreatedAt:    employee.CreatedAt,
		CreatedBy:    employee.CreatedBy,
		UpdatedAt:    employee.UpdatedAt,
		LastUpdateBy: employee.LastUpdateBy,
		DeletedAt:    employee.DeletedAt,
		ImageUrl:     employee.ImageUrl,
		ImageID:      employee.ImageID,
		IsActive:     employee.IsActive,
	})
	if err != nil {
		return err
	}
	employee.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (m *MySQL) Update(ctx context.Context, isAdmin bool, employee *entity.Employee) (err error) {
	query := `
	UPDATE 
		employee
	SET
		position_id = :position_id,
		name = :name,
		date_of_birth = :date_of_birth,
		address = :address,
		gender = :gender,
		email = :email,
		salary = :salary,
		bonus = :bonus,
		from_date = :from_date,
		finish_date = :finish_date,
		show_employee = :show_employee,
		updated_at =:updated_at,
		last_update_by =:last_update_by,
		image_url = :image_url,
		image_id = :image_id	
	WHERE 
		id = :id AND
		is_active = 1 AND
		app_id = :app_id
	`
	if !isAdmin {
		query += " AND created_by = :created_by"

	}
	res, err := m.db.NamedExecContext(ctx, query, &Employee{
		ID:           employee.ID,
		PositionID:   employee.PositionID,
		Name:         employee.Name,
		DateOfBirth:  employee.DateOfBirth,
		Address:      employee.Address,
		Gender:       employee.Gender,
		Email:        employee.Email,
		Salary:       employee.Salary,
		Bonus:        employee.Bonus,
		FromDate:     employee.FromDate,
		FinishDate:   employee.FinishDate,
		ShowEmployee: employee.ShowEmployee,
		AppID:        employee.AppID,
		CreatedAt:    employee.CreatedAt,
		CreatedBy:    employee.CreatedBy,
		UpdatedAt:    employee.UpdatedAt,
		LastUpdateBy: employee.LastUpdateBy,
		DeletedAt:    employee.DeletedAt,
		ImageUrl:     employee.ImageUrl,
		ImageID:      employee.ImageID,
		IsActive:     employee.IsActive,
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

func (m *MySQL) Delete(ctx context.Context, isAdmin bool, employee *entity.Employee) (err error) {
	query := `
	UPDATE 
		employee
	SET
		is_active = 0,
		deleted_at = :deleted_at,
		last_update_by = :last_update_by
	WHERE
		id = :id AND
		app_id = :app_id AND
		is_active = 1
		`
	if !isAdmin {
		query += " AND created_by = :created_by"

	}
	res, err := m.db.NamedExecContext(ctx, query, &Employee{
		ID:           employee.ID,
		AppID:        employee.AppID,
		LastUpdateBy: employee.LastUpdateBy,
		DeletedAt:    employee.DeletedAt,
		IsActive:     employee.IsActive,
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
