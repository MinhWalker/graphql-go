package resolver

import (
	"context"
	"graphQL-go/graph/generated"
	"graphQL-go/models"
)

type userResolver struct{ *Resolver }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {

	return r.MeetupsRepo.GetMeetupsByUserID(obj.ID)
}
