package usecase

import (
	"socmed/src/modules/profile/model"
	"socmed/src/modules/profile/repository"
)

// profileusecase
type profileUsecaseImpl struct {
	profileRepository repository.ProfileRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository) *profileUsecaseImpl {
	return &profileUsecaseImpl{profileRepository}
}

//SaveProfile
func (profileUsecase *profileUsecaseImpl) SaveProfile(profile *model.Profile) (*model.Profile, error) {

	err := profileUsecase.profileRepository.Save(profile)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

//UpdateProfile
func (profileUsecase *profileUsecaseImpl) UpdateProfile(id string, profile *model.Profile) (*model.Profile, error) {

	err := profileUsecase.profileRepository.Update(id, profile)

	if err != nil {
		return nil, err
	}
	return profile, nil
}

//GetByID
func (profileUsecase *profileUsecaseImpl) GetByID(id string) (*model.Profile, error) {

	profile, err := profileUsecase.profileRepository.FindByID(id)

	if err != nil {
		return nil, err
	}
	return profile, nil
}

//GetByEmail
func (profileUsecase *profileUsecaseImpl) GetByEmail(email string) (*model.Profile, error) {

	profile, err := profileUsecase.profileRepository.FindByID(email)

	if err != nil {
		return nil, err
	}
	return profile, nil
}
