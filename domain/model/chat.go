package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionChat = "chats"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id" json:"_"`
	UserID    primitive.ObjectID `bson:"userID" json:"user_id"`
	Message   string             `bson: message json:"message"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type ChatRepository interface {
	Create(c context.Context, chat *Chat) error
	FetchByUserID(c context.Context, userID string) ([]Chat, error)
}

type ChatUsecase interface {
	Create(c context.Context, chat *Chat) error
	FetchByUserID(c context.Context, userID string) ([]Chat, error)
	GetResponseFromOpenAI(c context.Context, prompt string, userID string, token string) (string, error)
}
