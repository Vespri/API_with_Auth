package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	UpdateUser(ID int, inputData UpdateUserInput) (User, error)
	GetUserByID(ID int) (User, error)
	DeleteUser(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Age = input.Age
	user.Email = input.Email
	user.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	NewUser, err := s.repository.Save(user)
	if err != nil {
		return NewUser, err
	}
	return NewUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUser(ID int, inputData UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID != inputData.User.ID {
		return user, errors.New("Not an owner of the user")
	}

	user.Email = inputData.Email
	user.Username = inputData.Username

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) DeleteUser(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	deleteUser, errfound := s.repository.Delete(user)
	if errfound != nil {
		return deleteUser, errfound
	}
	return deleteUser, nil
}
