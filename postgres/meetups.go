package postgres

import (
	"github.com/go-pg/pg/v10"
	"graphQL-go/models"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetups(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}

func (m *MeetupsRepo) GetMeetupsByUserID(userID string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}
