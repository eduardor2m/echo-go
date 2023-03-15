package repositories

import (
	"errors"

	"github.com/eduardor2m/echo-go/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Get(email string) (*domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(email string) error
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	err = user.Validate()

	if err != nil {
		return nil, err
	}

	if err := repo.Db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDB) Get(email string) (*domain.User, error) {
	var user domain.User

	err := repo.Db.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepositoryDB) GetAll() ([]domain.User, error) {
	var users []domain.User

	err := repo.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo UserRepositoryDB) Update(user *domain.User) (*domain.User, error) {
	err := repo.Db.Model(&user).Where("email = ?", user.Email).Updates(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDB) Delete(email string) error {
	err := repo.Db.Delete(&domain.User{}, "email = ?", email).Error

	if err != nil {
		return err
	}

	return nil
}
