package api_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/delivery/api"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/mocks"
	"github.com/nazyli/api-restaurant/service"
	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleGetUserById(t *testing.T) {
	var mockUser entity.User
	err := faker.FakeData(&mockUser)
	mockUser.UserHash = mockUser.CreatedBy
	assert.NoError(t, err)
	mockUserRepo := new(mocks.UserMock)
	service := service.New(1, 1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	id := int(mockUser.ID)
	mockUserRepo.On("GetByID", mock.Anything, int64(id)).Return(mockUser, nil)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/"+strconv.Itoa(id)+"?app_id=", strings.NewReader(""))

	scope := "read:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)

	router.ServeHTTP(rec, req)
	api.HandleGetUserById(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}

func TestHandleSelectUsers(t *testing.T) {
	var mockUser entity.User
	err := faker.FakeData(&mockUser)
	mockListUsers := make(entity.Users, 0)
	mockListUsers = append(mockListUsers, mockUser)

	mockUser.UserHash = mockUser.CreatedBy
	assert.NoError(t, err)
	mockUserRepo := new(mocks.UserMock)
	service := service.New(1, 1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	mockUserRepo.On("Select", mock.Anything).Return(mockListUsers, nil)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user?app_id=", strings.NewReader(""))

	scope := "read:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)

	router.ServeHTTP(rec, req)
	api.HandleSelectUsers(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}
func TestHandlePostUsers(t *testing.T) {
	mockUser := entity.User{
		Username:   "Username 1",
		Email:      "email@gmail.com",
		Password:   "Password 1",
		EmployeeID: nil,
		Scope:      "Scope 1",
	}
	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUserRepo := new(mocks.UserMock)
	service := service.New(1, 1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	mockUserRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(entity.User{}, errors.New("Not Found")).Once()
	mockUserRepo.On("Insert", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/user?app_id=", strings.NewReader(string(j)))

	scope := "create:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)
	router.ServeHTTP(rec, req)
	api.HandlePostUsers(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}
func TestHandlePatchUsers(t *testing.T) {
	mockUser := entity.User{
		ID:         1,
		Username:   "Username 1",
		Email:      "email@gmail.com",
		EmployeeID: nil,
		Scope:      "Scope 1",
	}
	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUserRepo := new(mocks.UserMock)
	service := service.New(1, 1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	id := int(mockUser.ID)
	mockUserRepo.On("GetByID", mock.Anything, int64(id)).Return(mockUser, nil)
	mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, int64(id)).Return(mockUser, nil)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", "/user/"+strconv.Itoa(id)+"?app_id=", strings.NewReader(string(j)))

	scope := "update:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)
	// params := req.URL.Query()
	// params.Set(":id", strconv.Itoa(id))
	router.ServeHTTP(rec, req)
	api.HandlePatchUsers(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}
func TestHandleDeleteUsers(t *testing.T) {
	mockUser := entity.User{
		ID: 1,
	}
	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUserRepo := new(mocks.UserMock)
	service := service.New(1, 1, mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	id := int(mockUser.ID)
	mockUserRepo.On("GetByID", mock.Anything, int64(id)).Return(mockUser, nil)
	mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil).Once()

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/user/"+strconv.Itoa(id)+"?app_id=", strings.NewReader(string(j)))

	scope := "delete:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)
	// params := req.URL.Query()
	// params.Set(":id", strconv.Itoa(id))
	router.ServeHTTP(rec, req)
	api.HandleDeleteUsers(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}
func TestHandleGetHello(t *testing.T) {
	service := service.New(1, 1, nil, nil, nil, nil, nil, nil, nil, nil)
	api := api.New(api.CloudinaryConfig{}, service)
	router := mux.NewRouter()
	api.Register(router)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", strings.NewReader(""))
	router.ServeHTTP(rec, req)

	api.HandleGetHello(rec, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
