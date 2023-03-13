package repositories

import (
	"github.com/eduardor2m/echo-go/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
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

	err = repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
