package usecase

import "github.com/dokkiichan/BridgeMe-Back/internal/domain"

type ProfileRepository interface {
	Store(profile *domain.Profile) error
	FindByID(id string) (*domain.Profile, error)
	FindAll() ([]*domain.Profile, error)
	Update(profile *domain.Profile) error
	Delete(id string) error
}
