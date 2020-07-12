package repository

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"socmed/src/modules/story/model"
)

type storyRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

//NewStoryRepositoryMongo
func NewStoryRepositoryMongo(db *mgo.Database, collection string) *storyRepositoryMongo {
	return &storyRepositoryMongo{db: db, collection: collection}
}

//Save
func (r *storyRepositoryMongo) Save(story *model.Story) error {
	err := r.db.C(r.collection).Insert(story)
	return err
}

//FindByID
func (r *storyRepositoryMongo) FindByID(id string) (*model.Story, error) {
	var story model.Story

	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&story)

	if err != nil {
		return nil, err
	}
	return &story, nil
}

//FindByProfileID
func (r *storyRepositoryMongo) FindByProfileID(profileID string) (model.Stories, error) {
	var stories model.Stories

	err := r.db.C(r.collection).Find(bson.M{"profile.id": profileID}).All(&stories)

	if err != nil {
		return nil, err
	}
	return stories, nil
}

//FindAll
func (r *storyRepositoryMongo) FindAll() (model.Stories, error) {
	var stories model.Stories

	err := r.db.C(r.collection).Find(bson.M{}).All(&stories)

	if err != nil {
		return nil, err
	}
	return stories, nil
}
