package main

import (
    "context"
    "github.com/umi0410/how-to-backend-in-go/testcode/ent"
)

type UserRepository interface{
    Create(input *UserCreateInput) (*ent.User, error)
}


func NewUserRepository(userClient *ent.UserClient) UserRepository{
    return &UserRepositoryImpl{
        Client: userClient,
    }
}

type UserRepositoryImpl struct{
   Client *ent.UserClient
}

func (u *UserRepositoryImpl) Create(input *UserCreateInput) (*ent.User, error) {
    user, err := u.Client.Create().
        SetID(input.ID).
        SetPassword(input.Password).
        SetName(input.Name).
        Save(context.TODO())
    return user, err
}
