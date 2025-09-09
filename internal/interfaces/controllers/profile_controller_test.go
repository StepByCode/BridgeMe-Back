package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// MockProfileInteractor is a mock for ProfileInteractorInterface
type MockProfileInteractor struct {
	mock.Mock
}

func (m *MockProfileInteractor) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	args := m.Called(profile)
	// Ensure we return a non-nil pointer for the first return value if the test expects one.
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) GetProfile(id string) (*domain.Profile, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) GetAllProfiles() ([]*domain.Profile, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) UpdateProfile(id string, profile *domain.Profile) (*domain.Profile, error) {
	args := m.Called(id, profile)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) DeleteProfile(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateProfile(t *testing.T) {
	e := echo.New()
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	profileJSON := `{"name":"Test User","affiliation":"Test Inc.","bio":"A test bio","instagram_id":"test_insta","twitter_id":"test_x"}`
	req := httptest.NewRequest(http.MethodPost, "/profiles", strings.NewReader(profileJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	createdProfile := &domain.Profile{ID: uuid.New().String(), Name: "Test User"}
	mockInteractor.On("CreateProfile", mock.AnythingOfType("*domain.Profile")).Return(createdProfile, nil)

	err := controller.CreateProfile(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, createdProfile.ID, response["id"])
	mockInteractor.AssertExpectations(t)
}

func TestGetProfileById(t *testing.T) {
	e := echo.New()
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	testUUID := uuid.New()
	profile := &domain.Profile{ID: testUUID.String(), Name: "Test User", Affiliation: "Test Inc.", Bio: "A test bio"}

	mockInteractor.On("GetProfile", testUUID.String()).Return(profile, nil)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUUID.String())

	// The controller method expects openapi_types.UUID
	var idParam openapi_types.UUID
	idParam, err := uuid.Parse(testUUID.String())
	assert.NoError(t, err)

	err = controller.GetProfileById(ctx, idParam)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseProfile generated.ProfileInput
	json.Unmarshal(rec.Body.Bytes(), &responseProfile)
	assert.Equal(t, profile.Name, *responseProfile.Name)
	mockInteractor.AssertExpectations(t)
}

func TestGetProfiles(t *testing.T) {
	e := echo.New()
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	profiles := []*domain.Profile{
		{ID: uuid.New().String(), Name: "Test User 1"},
		{ID: uuid.New().String(), Name: "Test User 2"},
	}
	mockInteractor.On("GetAllProfiles").Return(profiles, nil)

	req := httptest.NewRequest(http.MethodGet, "/profiles", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := controller.GetProfiles(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseProfiles []generated.ProfileInput
	json.Unmarshal(rec.Body.Bytes(), &responseProfiles)
	assert.Len(t, responseProfiles, 2)
	assert.Equal(t, profiles[0].Name, *responseProfiles[0].Name)
	mockInteractor.AssertExpectations(t)
}

func TestUpdateProfile(t *testing.T) {
	e := echo.New()
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	testUUID := uuid.New()
	profileJSON := `{"name":"Updated User","affiliation":"Updated Inc.","bio":"An updated bio","instagram_id":"updated_insta","twitter_id":"updated_x"}`
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(profileJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUUID.String())

	updatedProfile := &domain.Profile{ID: testUUID.String(), Name: "Updated User"}
	mockInteractor.On("UpdateProfile", testUUID.String(), mock.AnythingOfType("*domain.Profile")).Return(updatedProfile, nil)

	var idParam openapi_types.UUID
	idParam, err := uuid.Parse(testUUID.String())
	assert.NoError(t, err)

	err = controller.UpdateProfile(ctx, idParam)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, updatedProfile.ID, response["id"])
	mockInteractor.AssertExpectations(t)
}

func TestDeleteProfile(t *testing.T) {
	e := echo.New()
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	testUUID := uuid.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUUID.String())

	mockInteractor.On("DeleteProfile", testUUID.String()).Return(nil)

	var idParam openapi_types.UUID
	idParam, err := uuid.Parse(testUUID.String())
	assert.NoError(t, err)

	err = controller.DeleteProfile(ctx, idParam)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockInteractor.AssertExpectations(t)
}