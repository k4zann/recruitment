package resume

import (
	"context"
	"recruitment/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResumeStore struct {
	MongoCollection *mongo.Collection
}

func NewStore(mongoCollection *mongo.Collection) *ResumeStore {
	return &ResumeStore{MongoCollection: mongoCollection}
}

func (s *ResumeStore) CreateResume(resume types.Resume) error {
	_, err := s.MongoCollection.InsertOne(
		context.Background(),
		resume,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ResumeStore) GetResume(id primitive.ObjectID) (*types.Resume, error) {
	var resume types.Resume

	err := s.MongoCollection.FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&resume)
	if err != nil {
		return nil, err
	}

	return &resume, nil
}

func (s *ResumeStore) GetResumes() ([]types.Resume, error) {
	var resumes []types.Resume

	cursor, err := s.MongoCollection.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var resume types.Resume

		err := cursor.Decode(&resume)
		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}

func (s *ResumeStore) DeleteResume(id primitive.ObjectID) error {
	_, err := s.MongoCollection.DeleteOne(
		context.Background(),
		bson.M{"_id": id},
	)
	if err != nil {
		return err
	}

	return nil
}
