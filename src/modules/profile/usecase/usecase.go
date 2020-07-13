package usecase

import (
	"socmed/src/modules/profile/model"

	storyModel "socmed/src/modules/story/model"
)

//profileUseCase
type ProfileUsecase interface {
	SaveProfile(*model.Profile) (*model.Profile, error)
	UpdateProfile(string, *model.Profile) (*model.Profile, error)
	GetByID(string) (*model.Profile, error)
	GetByEmail(string) (*model.Profile, error)

	//story
	CreateaStory(*storyModel.Story) (*storyModel.Story, error)
}
