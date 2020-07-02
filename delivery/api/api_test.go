package api_test

import (
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

	scope := "su,read:user"
	token, err := auth.CreateToken(mockUser.UserHash, &scope)
	req.Header.Add("Authorization", "Bearer "+token.Token)

	router.ServeHTTP(rec, req)
	api.HandleGetHello(rec, req)

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
