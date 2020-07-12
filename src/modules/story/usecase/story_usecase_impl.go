package usecase

import (
	"socmed/src/modules/story/model"
	"socmed/src/modules/story/repository"
)

//story UsecaseImpl
type storyUsecaseImpl struct {
	storyRepository repository.StoryRepository
}

func NewStoryUsecase(storyRepository repository.StoryRepository) *storyUsecaseImpl {
	return &storyUsecaseImpl{storyRepository}
}

//GetAll
func (storyUsecase *storyUsecaseImpl) GetAll() (model.Stories, error) {
	stories, err := storyUsecase.storyRepository.FindAll()

	if err != nil {
		return nil, err
	}
	return stories, nil
}
