package usecases

import (
	"github.com/eduardor2m/echo-go/application/repositories"
	"github.com/eduardor2m/echo-go/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) Get(email string) (*domain.User, error) {
	user, err := u.UserRepository.Get(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) GetAll() ([]domain.User, error) {
	users, err := u.UserRepository.GetAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserUseCase) Update(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) Delete(id string) error {
	err := u.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
