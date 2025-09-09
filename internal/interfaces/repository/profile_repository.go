package repository

import (
	"context"
	"log"

	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/dokkiichan/BridgeMe-Back/internal/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProfileRepository(db *mongo.Database) usecase.ProfileRepository {
	return &ProfileRepositoryImpl{
		Collection: db.Collection("profiles"),
	}
}

func (r *ProfileRepositoryImpl) Store(profile *domain.Profile) error {
	_, err := r.Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Printf("Error storing profile: %v", err)
		return err
	}
	return nil
}

func (r *ProfileRepositoryImpl) FindByID(id string) (*domain.Profile, error) {
	var profile domain.Profile
	err := r.Collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Error finding profile by ID: %v", err)
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepositoryImpl) FindAll() ([]*domain.Profile, error) {
	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error finding all profiles: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var profiles []*domain.Profile
	if err = cursor.All(context.Background(), &profiles); err != nil {
		log.Printf("Error decoding profiles: %v", err)
		return nil, err
	}

	return profiles, nil
}

func (r *ProfileRepositoryImpl) Update(profile *domain.Profile) error {
	_, err := r.Collection.ReplaceOne(context.Background(), bson.M{"id": profile.ID}, profile)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		return err
	}
	return nil
}

func (r *ProfileRepositoryImpl) Delete(id string) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error deleting profile: %v", err)
		return err
	}
	return nil
}
