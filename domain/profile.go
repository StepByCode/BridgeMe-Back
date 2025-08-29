package domain

import "time"

type Profile struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Affiliation string    `json:"affiliation" validate:"required"`
	Bio         string    `json:"bio" validate:"required"`
	InstagramID string    `json:"instagram_id"`
	TwitterID   string    `json:"twitter_id"`
	CreatedAt   time.Time `json:"created_at"`
}
