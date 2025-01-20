package services

import (
	"context"
	"errors"

	"example.com/m/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type UserServiceImpl struct {
	UserCollection *mongo.Collection
	Ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		UserCollection: userCollection,
		Ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.UserCollection.InsertOne(u.Ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(username *string) (*models.User, error) {
	var user models.User
	query := bson.D{{Key: "user_name", Value: *username}}
	err := u.UserCollection.FindOne(u.Ctx, query).Decode(&user)
	return &user, err
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.UserCollection.Find(u.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.Ctx)

	for cursor.Next(u.Ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{{Key: "user_name", Value: user.Name}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "user_name", Value: user.Name},
		{Key: "user_age", Value: user.Age},
		{Key: "user_address", Value: user.Address},
	}}}
	result, err := u.UserCollection.UpdateOne(u.Ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no user found to update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(username *string) error {
	filter := bson.D{{Key: "user_name", Value: *username}}
	result, err := u.UserCollection.DeleteOne(u.Ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found to delete")
	}
	return nil
}
