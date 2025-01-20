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
type ProductService interface {
	CreateProduct(*models.Product) error
	GetProduct(*string) (*models.Product, error)
	GetAllProduct() ([]*models.Product, error)
	DeleteProduct(*string) error
}
type UserServiceImpl struct {
	UserCollection *mongo.Collection
	Ctx            context.Context
}
type ProductServiceImpl struct {
	productcollection *mongo.Collection
	Ctx               context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		UserCollection: userCollection,
		Ctx:            ctx,
	}
}
func NewProductService(productcollection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{
		productcollection: productcollection,
		Ctx:               ctx,
	}
}
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.UserCollection.InsertOne(u.Ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(email *string) (*models.User, error) {
	var user models.User
	query := bson.D{{Key: "email", Value: *email}}
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
	filter := bson.D{{Key: "eamil", Value: user.Email}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "user_name", Value: user.Name},
		{Key: "user_age", Value: user.Age},
		
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
	filter := bson.D{{Key: "email", Value: *username}}
	result, err := u.UserCollection.DeleteOne(u.Ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found to delete")
	}
	return nil
}
func (p *ProductServiceImpl) CreateProduct(product *models.Product) error {
	_, err := p.productcollection.InsertOne(p.Ctx, product)
	return err
}

func (p *ProductServiceImpl) GetProduct(pname *string) (*models.Product, error) {
	var product models.Product
	query := bson.D{{Key: "product_name", Value: *pname}}
	err := p.productcollection.FindOne(p.Ctx, query).Decode(&product)
	return &product, err
}

func (p *ProductServiceImpl) GetAllProduct() ([]*models.Product, error) {
	var products []*models.Product
	cursor, err := p.productcollection.Find(p.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.Ctx)

	for cursor.Next(p.Ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("no users found")
	}
	return products, nil
}

func (p *ProductServiceImpl) DeleteProduct(username *string) error {
	filter := bson.D{{Key: "name", Value: *username}}
	result, err := p.productcollection.DeleteOne(p.Ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found to delete")
	}
	return nil
}
