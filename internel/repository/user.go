package repository

import (
	"errors"
	"gorm.io/gorm"
	"timertracker/internel/service"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}

type UserRep struct {
	ServiceContext service.ServiceContext
}

type User struct {
	gorm.Model
	ID             uint   `gorm:"primary_key;type:bigint"`
	Name           string `gorm:"type:varchar(100)"`
	PassportNumber string `gorm:"unique;type:varchar(20)"`
	Surname        string `gorm:"type:varchar(100)"`
	Patronymic     string `gorm:"type:varchar(100)"`
	Address        string `gorm:"type:varchar(255)"`
	Tasks          []Task `gorm:"foreignKey:UserID"`
}

func NewUserRepository(serviceContext service.ServiceContext) *UserRep {
	return &UserRep{
		ServiceContext: serviceContext,
	}
}

func (userRepository *UserRep) FindAll() ([]User, error) {
	var users []User
	result := userRepository.ServiceContext.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userRepository *UserRep) FindByID(id int) (User, error) {
	var user User
	result := userRepository.ServiceContext.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, ErrorMsg
		}
		return User{}, result.Error
	}
	return user, nil
}

func (userRepository *UserRep) Create(user User) (User, error) {
	result := userRepository.ServiceContext.DB.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (userRepository *UserRep) Update(user User) (User, error) {
	if user.ID == 0 {
		return User{}, errors.New("пользователь не имеет установленного ID")
	}

	result := userRepository.ServiceContext.DB.Save(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (userRepository *UserRep) Delete(id int) error {
	result := userRepository.ServiceContext.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrorMsg
	}
	return nil
}
