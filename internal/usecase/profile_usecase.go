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
