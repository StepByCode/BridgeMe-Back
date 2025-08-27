package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dokkiichan/BridgeMe-Back/domain"
	"github.com/dokkiichan/BridgeMe-Back/usecase"
	"github.com/gorilla/mux"
)

type ProfileController struct {
	Interactor usecase.ProfileInteractor
}

func NewProfileController(interactor usecase.ProfileInteractor) *ProfileController {
	return &ProfileController{Interactor: interactor}
}

// @Summary Create a new profile
// @Description Create a new profile with the input payload
// @Accept  json
// @Produce  json
// @Param   profile  body      domain.Profile  true  "Profile to create"
// @Success 201      {object}  domain.Profile
// @Failure 400      {object}  string
// @Failure 500      {object}  string
// @Router /profiles [post]
func (c *ProfileController) Create(w http.ResponseWriter, r *http.Request) {
	var profile domain.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProfile, err := c.Interactor.CreateProfile(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProfile)
}

// @Summary Get a profile by ID
// @Description Get a single profile by its ID
// @Produce  json
// @Param   id   path      string  true  "Profile ID"
// @Success 200  {object}  domain.Profile
// @Failure 400  {object}  string
// @Failure 404  {object}  string
// @Failure 500  {object}  string
// @Router /profiles/{id} [get]
func (c *ProfileController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	profile, err := c.Interactor.GetProfile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// @Summary Get all profiles
// @Description Get a list of all profiles
// @Produce  json
// @Success 200  {array}   domain.Profile
// @Failure 500  {object}  string
// @Router /profiles [get]
func (c *ProfileController) Index(w http.ResponseWriter, r *http.Request) {
	profiles, err := c.Interactor.GetAllProfiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}
