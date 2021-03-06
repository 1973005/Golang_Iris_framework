package model

import (
	"time"
)

//profile
type Profile struct {
	ID           string    `bson:"id"`
	FirstName    string    `bson:"first_name"`
	LastName     string    `bson:"last_name"`
	Email        string    `bson:"email"`
	Password     string    `bson:"password"`
	ImageProfile string    `bson:"image_profile"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"update_at"`
}

//Profiles
type Profiles []Profile

//isvalidpassword
func (p *Profile) IsValidPassword(password string) bool {
	return p.Password == password
}
