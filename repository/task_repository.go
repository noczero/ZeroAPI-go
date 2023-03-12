package repository

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"

	"github.com/noczero/ZeroAPI-go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) model.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *model.Task) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]model.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var tasks []model.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []model.Task{}, err
	}

	return tasks, err
}
