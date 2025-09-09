package controllers

import (
	"net/http"
	"time"

	"github.com/dokkiichan/BridgeMe-Back/internal/domain"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/generated"
	"github.com/dokkiichan/BridgeMe-Back/internal/usecase"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ProfileController struct {
	Interactor usecase.ProfileUseCaseInterface
}

func NewProfileController(interactor usecase.ProfileUseCaseInterface) *ProfileController {
	return &ProfileController{Interactor: interactor}
}

// ensure ProfileController implements generated.ServerInterface
var _ generated.ServerInterface = (*ProfileController)(nil)

func (c *ProfileController) CreateProfile(ctx echo.Context) error {
	var req generated.CreateProfileJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	profile := &domain.Profile{
		Name:        *req.Name,
		Affiliation: *req.Affiliation,
		Bio:         *req.Bio,
		InstagramID: *req.InstagramId,
		TwitterID:   *req.TwitterId,
	}

	createdProfile, err := c.Interactor.CreateProfile(profile)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create profile")
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": createdProfile.ID})
}

func (c *ProfileController) GetProfileById(ctx echo.Context, id openapi_types.UUID) error {
	profile, err := c.Interactor.GetProfile(id.String())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get profile")
	}
	if profile == nil {
		return ctx.String(http.StatusNotFound, "Profile not found")
	}

	res := generated.Profile{
		Name:        &profile.Name,
		Affiliation: &profile.Affiliation,
		Bio:         &profile.Bio,
		InstagramId: &profile.InstagramID,
		TwitterId:   &profile.TwitterID,
		Id:          &profile.ID,
		CreatedAt:   func() *string { s := profile.CreatedAt.Format(time.RFC3339); return &s }(),
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *ProfileController) GetProfiles(ctx echo.Context) error {
	profiles, err := c.Interactor.GetAllProfiles()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get all profiles")
	}

	var res []generated.Profile
	for _, p := range profiles {
		res = append(res, generated.Profile{
			Name:        &p.Name,
			Affiliation: &p.Affiliation,
			Bio:         &p.Bio,
			InstagramId: &p.InstagramID,
			TwitterId:   &p.TwitterID,
			Id:          &p.ID,
			CreatedAt:   func() *string { s := p.CreatedAt.Format(time.RFC3339); return &s }(),
		})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *ProfileController) UpdateProfile(ctx echo.Context, id openapi_types.UUID) error {
	var req generated.UpdateProfileJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	profile := &domain.Profile{
		ID:          id.String(),
		Name:        *req.Name,
		Affiliation: *req.Affiliation,
		Bio:         *req.Bio,
		InstagramID: *req.InstagramId,
		TwitterID:   *req.TwitterId,
	}

	updatedProfile, err := c.Interactor.UpdateProfile(profile)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update profile")
	}

	res := generated.Profile{
		Name:        &updatedProfile.Name,
		Affiliation: &updatedProfile.Affiliation,
		Bio:         &updatedProfile.Bio,
		InstagramId: &updatedProfile.InstagramID,
		TwitterId:   &updatedProfile.TwitterID,
		Id:          &updatedProfile.ID,
		CreatedAt:   func() *string { s := updatedProfile.CreatedAt.Format(time.RFC3339); return &s }(),
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *ProfileController) DeleteProfile(ctx echo.Context, id openapi_types.UUID) error {
	if err := c.Interactor.DeleteProfile(id.String()); err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete profile")
	}

	return ctx.NoContent(http.StatusNoContent)
}
