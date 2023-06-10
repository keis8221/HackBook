package repo

import (
	"fmt"

	"github.com/keis8221/surprise/api/model"
	"github.com/keis8221/surprise/api/user"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

func (ur *UserRepoImpl) Create(user *model.User) (*model.User, error) {
	fmt.Print(&user)
	if err := ur.DB.Save(&user).Error; err != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return user, nil
}




