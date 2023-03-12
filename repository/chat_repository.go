package repository

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chatRepository struct {
	database   mongo.Database
	collection string
}

func NewChatRepository(db mongo.Database, collection string) model.ChatRepository {
	return &chatRepository{
		database:   db,
		collection: collection,
	}
}

func (cr chatRepository) Create(c context.Context, chat *model.Chat) error {
	collection := cr.database.Collection(cr.collection)

	_, err := collection.InsertOne(c, chat)

	return err
}

func (cr chatRepository) FetchByUserID(c context.Context, userID string) ([]model.Chat, error) {
	collection := cr.database.Collection(cr.collection)

	var chats []model.Chat

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return chats, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &chats)
	if chats == nil {
		return []model.Chat{}, err
	}

	return chats, err

}
