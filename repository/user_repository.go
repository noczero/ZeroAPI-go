package repository

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"

	"github.com/noczero/ZeroAPI-go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) model.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *model.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]model.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []model.User

	err = cursor.All(c, &users)
	if users == nil {
		return []model.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (model.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user model.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (model.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user model.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
