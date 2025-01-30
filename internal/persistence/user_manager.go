package persistence

import "github.com/t-shah02/mochi/internal/models"

type UserManager struct {
	userFilePath string
	managedUsers []*models.User
}

func NewUserManager(userFilePath string) *UserManager {
	return &UserManager{
		userFilePath: userFilePath,
		managedUsers: make([]*models.User, 0),
	}
}

func (manager *UserManager) Init() {

}
