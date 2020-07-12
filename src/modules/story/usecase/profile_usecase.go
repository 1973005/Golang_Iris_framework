package usecase

import (
	"socmed/src/modules/story/model"
)

//StoryUsecase
type StoryUsecase interface {
	GetAll() (model.Stories, error)
}
