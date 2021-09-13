package postgres

import (
	"github.com/go-pg/pg/v10"
	"graphQL-go/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (m *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
