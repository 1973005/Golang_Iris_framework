package repository

import (
	"socmed/src/modules/story/model"
)

//StoryRepository
type StoryRepository interface {
	Save(*model.Story) error
	FindByID(string) (*model.Story, error)
	FindByProfileID(string) (model.Stories, error)
	FindAll() (model.Stories, error)
}
