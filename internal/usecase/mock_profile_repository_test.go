
package usecase

import (
	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockProfileRepository is a mock implementation of ProfileRepository
type MockProfileRepository struct {
	mock.Mock
}

func (m *MockProfileRepository) Store(profile *domain.Profile) error {
	args := m.Called(profile)
	return args.Error(0)
}

func (m *MockProfileRepository) FindByID(id string) (*domain.Profile, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) FindAll() ([]*domain.Profile, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) Update(profile *domain.Profile) error {
	args := m.Called(profile)
	return args.Error(0)
}

func (m *MockProfileRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
