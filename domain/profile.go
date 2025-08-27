package domain

import "time"

type Profile struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Affiliation string    `json:"affiliation"`
	Bio         string    `json:"bio"`
	InstagramID string    `json:"instagram_id"`
	TwitterID   string    `json:"twitter_id"`
	CreatedAt   time.Time `json:"created_at"`
}
