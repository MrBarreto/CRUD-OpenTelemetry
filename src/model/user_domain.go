package model

import (
	customError "github.com/MrBarreto/CRUD-OpenTelemetry/src/configuration/custom_error"
)

type userDomain struct {
	email    string
	password string
	name     string
	age      int
}

func NewUserDomain(email, password, name string, age int) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

type UserDomainInterface interface {
	CreateUser() *customError.CustomError
	UpdateUser(string) *customError.CustomError
	FindUser(string) (UserDomainInterface, *customError.CustomError)
	DeleteUser(string) *customError.CustomError
}

func (ud *userDomain) CreateUser() *customError.CustomError {
	return nil
}

func (ud *userDomain) UpdateUser(string) *customError.CustomError {
	return nil
}

func (ud *userDomain) FindUser(string) (UserDomainInterface, *customError.CustomError) {
	return nil, nil
}

func (ud *userDomain) DeleteUser(string) *customError.CustomError {
	return nil
}
