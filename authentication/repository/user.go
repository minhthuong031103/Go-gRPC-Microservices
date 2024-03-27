package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository interface {
	Save()
}

type userRepository struct {
	mongoDb *mongo.Client
}

func NewUserRepository(mongoDb *mongo.Client) UserRepository {
	return &userRepository{mongoDb: mongoDb}
}

func (u *userRepository) Save() {
	//save user to database
}
