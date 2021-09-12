package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphQL-go/graph/generated"
	"graphQL-go/graph/model"
	"graphQL-go/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}
	if user == nil {
		return nil, errors.New("User with id not exit")
	}

	return user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
