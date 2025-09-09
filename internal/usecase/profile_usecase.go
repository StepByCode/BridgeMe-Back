package usecase

import (
	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/google/uuid"
	"time"
)

type ProfileUseCaseInterface interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
	GetProfile(id string) (*domain.Profile, error)
	GetAllProfiles() ([]*domain.Profile, error)
	UpdateProfile(profile *domain.Profile) (*domain.Profile, error)
	DeleteProfile(id string) error
}

type ProfileUseCase struct {
	ProfileRepository ProfileRepository
}

func NewProfileUseCase(repo ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{ProfileRepository: repo}
}

func (uc *ProfileUseCase) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	profile.ID = uuid.New().String()
	profile.CreatedAt = time.Now()
	if err := uc.ProfileRepository.Store(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func (uc *ProfileUseCase) GetProfile(id string) (*domain.Profile, error) {
	return uc.ProfileRepository.FindByID(id)
}

func (uc *ProfileUseCase) GetAllProfiles() ([]*domain.Profile, error) {
	return uc.ProfileRepository.FindAll()
}

func (uc *ProfileUseCase) UpdateProfile(profile *domain.Profile) (*domain.Profile, error) {
	if err := uc.ProfileRepository.Update(profile); err != nil {
		return nil, err
	}
	// Fetch the updated profile to ensure all fields, including CreatedAt, are correct
	updatedProfile, err := uc.ProfileRepository.FindByID(profile.ID)
	if err != nil {
		return nil, err
	}
	return updatedProfile, nil
}

func (uc *ProfileUseCase) DeleteProfile(id string) error {
	return uc.ProfileRepository.Delete(id)
}
