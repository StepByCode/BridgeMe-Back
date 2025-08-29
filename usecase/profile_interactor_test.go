
package usecase

import (
	"testing"

	"github.com/dokkiichan/BridgeMe-Back/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	interactor := NewProfileInteractor(mockRepo)

	profile := &domain.Profile{Name: "Test User"}

	mockRepo.On("Store", mock.AnythingOfType("*domain.Profile")).Return(nil)

	createdProfile, err := interactor.CreateProfile(profile)

	assert.NoError(t, err)
	assert.NotNil(t, createdProfile)
	assert.NotEmpty(t, createdProfile.ID)
	assert.Equal(t, "Test User", createdProfile.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	interactor := NewProfileInteractor(mockRepo)

	profile := &domain.Profile{ID: "test-id", Name: "Test User"}

	mockRepo.On("FindByID", "test-id").Return(profile, nil)

	foundProfile, err := interactor.GetProfile("test-id")

	assert.NoError(t, err)
	assert.NotNil(t, foundProfile)
	assert.Equal(t, "test-id", foundProfile.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetAllProfiles(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	interactor := NewProfileInteractor(mockRepo)

	profiles := []*domain.Profile{
		{ID: "test-id-1", Name: "Test User 1"},
		{ID: "test-id-2", Name: "Test User 2"},
	}

	mockRepo.On("FindAll").Return(profiles, nil)

	foundProfiles, err := interactor.GetAllProfiles()

	assert.NoError(t, err)
	assert.NotNil(t, foundProfiles)
	assert.Len(t, foundProfiles, 2)
	mockRepo.AssertExpectations(t)
}
