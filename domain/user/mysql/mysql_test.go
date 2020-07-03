package mysql_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_userMysql "github.com/nazyli/api-restaurant/domain/user/mysql"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestGetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "user_hash", "employee_id", "scope"}).
		AddRow(1, "username 1", "email 1", "password 1", "user_hash 1", 1, "scope")

	query := "SELECT id, username, email, password, user_hash, employee_id, scope FROM users WHERE is_active = 1 AND email = \\? AND app_id = \\?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	a := _userMysql.New(sqlxDB)

	num := int64(1)
	email := "test@gmail.com"
	anArticle, err := a.GetByEmail(context.TODO(), num, email)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "username", "email", "user_hash", "employee_id", "scope", "app_id", "created_at", "created_by", "updated_at", "last_update_by", "deleted_at", "is_active"}).
		AddRow(1, "username 1", "email 1", "user_hash 1", 1, "scope", 1, time.Now(), "created_by 1", time.Now(), "last_update_by 1", time.Now(), 1)

	query := "SELECT id, username, email, user_hash, employee_id, scope, app_id, created_at, created_by, updated_at, last_update_by, deleted_at, is_active FROM users WHERE id = \\? AND app_id = \\? AND is_active = 1"
	mock.ExpectQuery(query).WillReturnRows(rows)
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	a := _userMysql.New(sqlxDB)

	anArticle, err := a.GetByID(context.TODO(), 1, 1, false, true, "uid 1")
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestGetByHash(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "username", "email", "user_hash", "employee_id", "scope", "app_id", "created_at", "created_by", "updated_at", "last_update_by", "deleted_at", "is_active"}).
		AddRow(1, "username 1", "email 1", "user_hash 1", 1, "scope", 1, time.Now(), "created_by 1", time.Now(), "last_update_by 1", time.Now(), 1)

	query := "SELECT id, username, email, user_hash, employee_id, scope, app_id, created_at, created_by, updated_at, last_update_by, deleted_at, is_active FROM users WHERE user_hash = \\? AND app_id = \\? AND is_active = 1"
	mock.ExpectQuery(query).WillReturnRows(rows)
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	a := _userMysql.New(sqlxDB)

	anArticle, err := a.GetByHash(context.TODO(), 1, false, false, "uid 1")
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
func TestSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockUser := entity.Users{
		entity.User{
			ID:           1,
			Username:     "Username 1",
			Email:        "Email 1",
			UserHash:     "UserHash 1",
			EmployeeID:   nil,
			Scope:        "Scope 1",
			CreatedAt:    null.TimeFrom(time.Now()),
			CreatedBy:    "CreatedBy 1",
			UpdatedAt:    null.TimeFrom(time.Now()),
			LastUpdateBy: nil,
			DeletedAt:    null.TimeFrom(time.Now()),
			AppID:        1,
		},
		entity.User{
			ID:           2,
			Username:     "Username 2",
			Email:        "Email 2",
			UserHash:     "UserHash 2",
			EmployeeID:   nil,
			Scope:        "Scope 2",
			CreatedAt:    null.TimeFrom(time.Now()),
			CreatedBy:    "CreatedBy 2",
			UpdatedAt:    null.TimeFrom(time.Now()),
			LastUpdateBy: nil,
			DeletedAt:    null.TimeFrom(time.Now()),
			IsActive:     1,
			AppID:        1,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "user_hash", "employee_id", "scope", "app_id", "created_at", "created_by", "updated_at", "last_update_by", "deleted_at", "is_active"}).
		AddRow(mockUser[0].ID, mockUser[0].Username, mockUser[0].Email, mockUser[0].UserHash, mockUser[0].EmployeeID, mockUser[0].Scope, mockUser[0].AppID, mockUser[0].CreatedAt, mockUser[0].CreatedBy, mockUser[0].UpdatedAt, mockUser[0].LastUpdateBy, mockUser[0].DeletedAt, mockUser[0].IsActive).
		AddRow(mockUser[1].ID, mockUser[1].Username, mockUser[1].Email, mockUser[1].UserHash, mockUser[1].EmployeeID, mockUser[1].Scope, mockUser[1].AppID, mockUser[1].CreatedAt, mockUser[0].CreatedBy, mockUser[1].UpdatedAt, mockUser[1].LastUpdateBy, mockUser[1].DeletedAt, mockUser[1].IsActive)

	query := "SELECT id, username, email, user_hash, employee_id, scope, app_id, created_at, created_by, updated_at, last_update_by, deleted_at, is_active FROM users WHERE app_id = \\? AND is_active = 1"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectQuery(query).WillReturnRows(rows)
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	a := _userMysql.New(sqlxDB)
	list, err := a.Select(context.TODO(), 1, false, true, "uid1")
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestInsert(t *testing.T) {
	var eid int64
	eid = 1
	us := &entity.User{
		Username:   "Username 1",
		Email:      "Email 1",
		Password:   "Password 1",
		UserHash:   "eb55808b848359c7566d41a69d712cc7d421dca3",
		EmployeeID: &eid,
		Scope:      "Scope 1",
		AppID:      1,
		CreatedAt:  null.TimeFrom(time.Now()),
		CreatedBy:  "CreatedBy 1",
		IsActive:   1,
	}
	t.Run("Must return the newly created user id, if given user input is valid", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		query := "INSERT INTO users ( username, email, password, user_hash, employee_id, scope, app_id, created_at, created_by, is_active ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
		// prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(us.Username, us.Email, us.Password, us.UserHash, us.EmployeeID, us.Scope, us.AppID, us.CreatedAt, us.CreatedBy, us.IsActive).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		// mock.ExpectRollback()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		a := _userMysql.New(sqlxDB)
		err = a.Insert(context.TODO(), us)
		assert.NoError(t, err)
		// assert.Nil(t, err)
		// assert.NotNil(t, us.ID)
		assert.Equal(t, int64(1), us.ID)
	})

}

func TestUpdate(t *testing.T) {
	var eid int64
	eid = 1
	us := &entity.User{
		ID:           12,
		Username:     "Username 1",
		Email:        "Email 1",
		EmployeeID:   &eid,
		Scope:        "Scope 1",
		AppID:        1,
		UpdatedAt:    null.TimeFrom(time.Now()),
		LastUpdateBy: nil,
		IsActive:     1,
	}
	t.Run("Must return the newly updated user id, if given user input is valid", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		query := "UPDATE users SET username =?, email =?, employee_id =?, scope =?, app_id =?, updated_at =?, last_update_by =? WHERE id = ? AND is_active = 1 AND app_id = ?"
		// prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(us.Username, us.Email, us.EmployeeID, us.Scope, us.AppID, us.UpdatedAt, us.LastUpdateBy, us.ID, us.AppID).WillReturnResult(sqlmock.NewResult(12, 1))
		mock.ExpectCommit()
		// mock.ExpectRollback()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		a := _userMysql.New(sqlxDB)
		err = a.Update(context.TODO(), true, us)
		assert.NoError(t, err)
		// assert.Nil(t, err)
		// assert.NotNil(t, us.ID)
		assert.Equal(t, int64(12), us.ID)
	})
}

func TestDelete(t *testing.T) {
	us := &entity.User{
		ID:           12,
		AppID:        1,
		DeletedAt:    null.TimeFrom(time.Now()),
		LastUpdateBy: nil,
	}
	t.Run("Must return the null err deleted, if given user input is valid", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		query := `UPDATE
					users
				SET
					is_active = 0,
					deleted_at = ?,
					last_update_by = ?
				WHERE
					id = ? AND
					is_active = 1 AND
					app_id = ?`
		// prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
		// mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(us.DeletedAt, us.LastUpdateBy, us.ID, us.AppID).WillReturnResult(sqlmock.NewResult(12, 1))
		// mock.ExpectCommit()
		// mock.ExpectRollback()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		a := _userMysql.New(sqlxDB)
		err = a.Delete(context.TODO(), true, us)
		assert.NoError(t, err)
		// assert.Nil(t, err)
		// assert.NotNil(t, us.ID)
		assert.Equal(t, int64(12), us.ID)
	})
}
