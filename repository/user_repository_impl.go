package repository

import (
	"go-fiber-clean-arch/config"
	"go-fiber-clean-arch/entity"
	"go-fiber-clean-arch/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		Collection: database.Collection("User"),
	}
}

func (repository *userRepositoryImpl) Create(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":      user.Id,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	})
	exception.PanicIfErr(err)
}

func (repository *userRepositoryImpl) FindAll() (users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfErr(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfErr(err)

	for _, document := range documents {
		users = append(users, entity.User{
			Id:       document["_id"].(string),
			Username: document["username"].(string),
			Email:    document["email"].(string),
			Password: document["password"].(string),
		})
	}
	return users
}

func (repository *userRepositoryImpl) FindById(id string) (user *entity.User) {
	filter := bson.D{{"_id", id}}
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var document bson.M
	err := repository.Collection.FindOne(ctx, filter).Decode(&document)
	exception.PanicIfErr(err)

	return &entity.User{
		Id:       document["_id"].(string),
		Username: document["username"].(string),
		Email:    document["email"].(string),
		Password: document["password"].(string),
	}
}

func (repository *userRepositoryImpl) FindByEmail(email string) (user *entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err := repository.Collection.FindOne(ctx, email).Decode(&user)
	exception.PanicIfErr(err)
	return user
}
