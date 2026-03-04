package models

import "time"

type url struct {
	Id				string
	OriginalURL		string
	ShortCode		string
	CreatedAt		time.Time
	HitCount		int	
}