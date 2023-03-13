package domain

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	Email    string `json:"email" valid:"notnull,email" gorm:"type:varchar(255),unique_index"`
	Password string `json:"-" valid:"notnull" gorm:"type:varchar(255)"`
	// Token    string `json:"token" gorm:"type:uuid,unique_index"`
}

func NewUser(name, email, password string) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	err = user.Validate()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Prepare() error {

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(password)

	// token := uuid.NewV4().String()
	// u.Token = token

	return nil

}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
