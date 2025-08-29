
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dokkiichan/BridgeMe-Back/domain"
	
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProfileInteractor is a mock for ProfileInteractor
type MockProfileInteractor struct {
	mock.Mock
}

func (m *MockProfileInteractor) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	args := m.Called(profile)
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) GetProfile(id string) (*domain.Profile, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileInteractor) GetAllProfiles() ([]*domain.Profile, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Profile), args.Error(1)
}

func TestCreate(t *testing.T) {
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	profile := &domain.Profile{ID: "new-id", Name: "test", Affiliation: "test", Bio: "test"}
	mockInteractor.On("CreateProfile", mock.AnythingOfType("*domain.Profile")).Return(profile, nil)

	body, _ := json.Marshal(map[string]string{
		"name": "test",
		"affiliation": "test",
		"bio": "test",
	})

	req := httptest.NewRequest("POST", "/profiles", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	controller.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockInteractor.AssertExpectations(t)
}

func TestShow(t *testing.T) {
	mockInteractor := new(MockProfileInteractor)
	controller := NewProfileController(mockInteractor)

	profile := &domain.Profile{ID: "test-id", Name: "Test User"}
	mockInteractor.On("GetProfile", "test-id").Return(profile, nil)

	req := httptest.NewRequest("GET", "/profiles/test-id", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/profiles/{id}", controller.Show)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockInteractor.AssertExpectations(t)
}
