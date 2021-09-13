package resolver

import (
	"context"
	"graphQL-go/graph/generated"
	"graphQL-go/models"
)

type meetupResolver struct{ *Resolver }

func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return r.UsersRepo.GetUserByID(obj.UserID)
}
