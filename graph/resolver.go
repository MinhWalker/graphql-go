//go:generate go run github.com/99designs/gqlgen

package graph

import "graphQL-go/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A second description",
		UserID:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Peter",
		Email:    "peter@gmail.com",
	},
}

type Resolver struct{}
