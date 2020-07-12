package model

import (
	profileModel "socmed/src/modules/profile/model"
	"time"
)

type Story struct {
	ID        string                `bson:"id"`
	Profile   *profileModel.Profile `bson:"profile"`
	Title     string                `bson:"title"`
	Content   string                `bson:"content"`
	CreatedAt time.Time             `bson:"created_at"`
	UpdatedAt time.Time             `bson:"Updated_at"`
}

//Stories, List dari Story
type Stories []Story
