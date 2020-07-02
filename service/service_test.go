package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/mocks"
	"github.com/nazyli/api-restaurant/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	null "gopkg.in/guregu/null.v3"
)

func TestGetUserByID(t *testing.T) {
	mockUser := entity.User{
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
	}
	mockUserRepo := new(mocks.UserMock)
	svc := service.New(1, 0.1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		a, _ := svc.GetUserByID(context.TODO(), mockUser.ID, false, false, "uid")
		// assert.Equal(t, http.StatusOK, err.Code)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(entity.User{}, nil /*errors.New("Unexpected")*/).Once()
		a, _ := svc.GetUserByID(context.TODO(), mockUser.ID, false, false, "uid")
		// assert.Error(t, errors.New(err.ErrMsg))
		assert.Equal(t, &entity.User{}, a)
		// assert.NotEqual(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestSelectUsers(t *testing.T) {
	mockUser := entity.User{
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
	}
	mockListUser := make(entity.Users, 0)
	mockListUser = append(mockListUser, mockUser)
	mockUserRepo := new(mocks.UserMock)
	svc := service.New(1, 0.1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Select", mock.Anything).Return(mockListUser, nil).Once()
		a, _ := svc.SelectUsers(context.TODO(), false, false, "uid")
		assert.Len(t, a, len(mockListUser))
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("Select", mock.Anything).Return(nil, nil /*errors.New("Unexpected")*/).Once()
		a, _ := svc.SelectUsers(context.TODO(), false, false, "uid")
		assert.Len(t, a, 0)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestInsertUser(t *testing.T) {
	mockUser := entity.User{
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
	}
	mockUserRepo := new(mocks.UserMock)
	svc := service.New(1, 0.1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(entity.User{}, errors.New("Not Found")).Once()
		mockUserRepo.On("Insert", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()
		_, err := svc.InsertUser(context.TODO(), "uid", &mockUser)
		assert.Equal(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("existing-email", func(t *testing.T) {
		mockUserRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		_, err := svc.InsertUser(context.TODO(), "uid", &mockUser)
		assert.Equal(t, http.StatusFound, err.Code)
		mockUserRepo.AssertExpectations(t)

	})
}

func TestUpdateUser(t *testing.T) {
	mockUser := entity.User{
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
	}
	mockUserRepo := new(mocks.UserMock)
	svc := service.New(1, 0.1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		_, err := svc.UpdateUser(context.TODO(), mockUser.ID, false, "uid", &mockUser)
		assert.Equal(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("id-not-found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(entity.User{}, errors.New("id-not-found")).Once()
		_, err := svc.UpdateUser(context.TODO(), mockUser.ID, false, "uid", &mockUser)
		assert.NotEqual(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})

}
func TestDeleteUser(t *testing.T) {
	mockUser := entity.User{
		ID:           1,
		CreatedBy:    "CreatedBy 1",
		LastUpdateBy: nil,
		DeletedAt:    null.TimeFrom(time.Now()),
		AppID:        1,
	}
	mockUserRepo := new(mocks.UserMock)
	svc := service.New(1, 0.1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()
		err := svc.DeleteUser(context.TODO(), mockUser.ID, false, "uid")
		assert.Equal(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("id-not-found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(entity.User{}, errors.New("id-not-found")).Once()
		err := svc.DeleteUser(context.TODO(), mockUser.ID, false, "uid")
		assert.NotEqual(t, http.StatusOK, err.Code)
		mockUserRepo.AssertExpectations(t)
	})
}
