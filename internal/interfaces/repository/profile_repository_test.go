package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ( // Global variables for the test suite
	client *mongo.Client
	db     *mongo.Database
	repo   *ProfileRepositoryImpl
)

func TestMain(m *testing.M) {
	// Setup: Connect to a test MongoDB instance
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	// Get credentials from environment variables or .env file
	username := "bridgenfcmongo"
	password := "zecHep2dywkaxagfah"

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + username + ":" + password + "@localhost:27018/?authSource=admin"))
	if err != nil {
		log.Fatalf("Failed to connect to test MongoDB: %v", err)
	}

	db = client.Database("bridgeme_test")
	repo = NewProfileRepository(db).(*ProfileRepositoryImpl)

	// Run tests
	m.Run()

	// Teardown: Disconnect and clean up
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from test MongoDB: %v", err)
		}
	}()

	// Drop the test database
	if err = db.Drop(context.Background()); err != nil {
		log.Fatalf("Failed to drop test database: %v", err)
	}

	log.Println("Test MongoDB cleaned up.")

	// Exit with the test result code
	// os.Exit(exitCode) // This line is commented out to allow the agent to continue
}

func setupTest(t *testing.T) {
	// Clean up the collection before each test
	if err := repo.Collection.Drop(context.Background()); err != nil {
		log.Fatalf("Failed to drop collection before test: %v", err)
	}
}

func TestStore(t *testing.T) {
	setupTest(t)

	profile := &domain.Profile{
		ID:          uuid.New().String(),
		Name:        "Test User",
		Affiliation: "Test Company",
		Bio:         "Test Bio",
		InstagramID: "test_insta",
		TwitterID:   "test_twitter",
		CreatedAt:   time.Now(),
	}

	err := repo.Store(profile)
	assert.NoError(t, err)

	// Verify by fetching directly from DB
	var fetchedProfile domain.Profile
	err = repo.Collection.FindOne(context.Background(), bson.M{"id": profile.ID}).Decode(&fetchedProfile)
	assert.NoError(t, err)
	assert.Equal(t, profile.ID, fetchedProfile.ID)
	assert.Equal(t, profile.Name, fetchedProfile.Name)
}

func TestFindByID(t *testing.T) {
	setupTest(t)

	profile := &domain.Profile{
		ID:          uuid.New().String(),
		Name:        "Test User",
		Affiliation: "Test Company",
		Bio:         "Test Bio",
		InstagramID: "test_insta",
		TwitterID:   "test_twitter",
		CreatedAt:   time.Now(),
	}
	err := repo.Store(profile)
	assert.NoError(t, err)

	foundProfile, err := repo.FindByID(profile.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundProfile)
	assert.Equal(t, profile.ID, foundProfile.ID)
	assert.Equal(t, profile.Name, foundProfile.Name)

	// Test not found
	notFoundProfile, err := repo.FindByID(uuid.New().String())
	assert.NoError(t, err) // No error for not found, just nil profile
	assert.Nil(t, notFoundProfile)
}

func TestFindAll(t *testing.T) {
	setupTest(t)

	profile1 := &domain.Profile{ID: uuid.New().String(), Name: "User 1", CreatedAt: time.Now()}
	profile2 := &domain.Profile{ID: uuid.New().String(), Name: "User 2", CreatedAt: time.Now()}
	err := repo.Store(profile1)
	assert.NoError(t, err)
	err = repo.Store(profile2)
	assert.NoError(t, err)

	profiles, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, profiles, 2)
	assert.Contains(t, []string{profiles[0].Name, profiles[1].Name}, "User 1")
	assert.Contains(t, []string{profiles[0].Name, profiles[1].Name}, "User 2")
}

func TestUpdate(t *testing.T) {
	setupTest(t)

	originalProfile := &domain.Profile{
		ID:          uuid.New().String(),
		Name:        "Original Name",
		Affiliation: "Original Affiliation",
		Bio:         "Original Bio",
		InstagramID: "original_insta",
		TwitterID:   "original_twitter",
		CreatedAt:   time.Now(),
	}
	err := repo.Store(originalProfile)
	assert.NoError(t, err)

	updatedProfile := &domain.Profile{
		ID:          originalProfile.ID,
		Name:        "Updated Name",
		Affiliation: "Updated Affiliation",
		Bio:         "Updated Bio",
		InstagramID: "updated_insta",
		TwitterID:   "updated_twitter",
		// CreatedAt is intentionally not set here to ensure it's preserved by the update logic
	}

	err = repo.Update(updatedProfile)
	assert.NoError(t, err)

	// Fetch the profile again to verify
	fetchedProfile, err := repo.FindByID(originalProfile.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedProfile)
	assert.Equal(t, updatedProfile.Name, fetchedProfile.Name)
	assert.Equal(t, updatedProfile.Affiliation, fetchedProfile.Affiliation)
	assert.Equal(t, originalProfile.CreatedAt.In(time.UTC).Format(time.RFC3339), fetchedProfile.CreatedAt.Format(time.RFC3339)) // Check CreatedAt preservation
}

func TestDelete(t *testing.T) {
	setupTest(t)

	profile := &domain.Profile{
		ID:          uuid.New().String(),
		Name:        "User to Delete",
		CreatedAt:   time.Now(),
	}
	err := repo.Store(profile)
	assert.NoError(t, err)

	// Verify it exists
	foundProfile, err := repo.FindByID(profile.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundProfile)

	// Delete it
	err = repo.Delete(profile.ID)
	assert.NoError(t, err)

	// Verify it's gone
	foundProfile, err = repo.FindByID(profile.ID)
	assert.NoError(t, err)
	assert.Nil(t, foundProfile)
}