package usecase

import (
	"github.com/dokkiichan/BridgeMe-Back/domain"
	"github.com/google/uuid"
	"time"
)

type ProfileInteractorInterface interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
	GetProfile(id string) (*domain.Profile, error)
	GetAllProfiles() ([]*domain.Profile, error)
}

type ProfileInteractor struct {
	ProfileRepository ProfileRepository
}

func NewProfileInteractor(repo ProfileRepository) *ProfileInteractor {
	return &ProfileInteractor{ProfileRepository: repo}
}

func (interactor *ProfileInteractor) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	profile.ID = uuid.New().String()
	profile.CreatedAt = time.Now()
	if err := interactor.ProfileRepository.Store(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func (interactor *ProfileInteractor) GetProfile(id string) (*domain.Profile, error) {
	return interactor.ProfileRepository.FindByID(id)
}

func (interactor *ProfileInteractor) GetAllProfiles() ([]*domain.Profile, error) {
	return interactor.ProfileRepository.FindAll()
}
