package services

import (
	"context"
	"errors"

	"example.com/contacts_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection 	*mongo.Collection
	context 		context.Context
}

func NewUserServiceImpl(userCollection *mongo.Collection, context context.Context) *UserServiceImpl {
	return &UserServiceImpl{userCollection: userCollection, context: context}
}

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := us.userCollection.InsertOne(us.context, user)
	return err
}

func (us *UserServiceImpl) GetUser(userName *string) (*models.User, error) {
	var user *models.User
	err := us.userCollection.FindOne(us.context, bson.M{"user_name": userName}).Decode(&user)
	return user, err
}

func (us *UserServiceImpl) GetUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := us.userCollection.Find(us.context, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(us.context)
	for cursor.Next(us.context) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("users not found")
	}

	return users, nil
}

func (us *UserServiceImpl) UpdateUser(user *models.User) error {
	result, err := us.userCollection.UpdateOne(us.context, bson.M{"user_name": user.Name}, bson.M{"$set": user})
	if result.MatchedCount != 1 {
		return errors.New("user not found")
	}
	return err
}

func (us *UserServiceImpl) DeleteUser(userName *string) error {
	result, err := us.userCollection.DeleteOne(us.context, bson.M{"user_name": userName})
	if result.DeletedCount != 1 {
		return errors.New("user not found")
	}
	return err
}