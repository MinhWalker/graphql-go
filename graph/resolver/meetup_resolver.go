package resolver

import (
	"context"
	"graphQL-go/dataloader"
	"graphQL-go/graph/generated"
	"graphQL-go/models"
	"strconv"
)

type meetupResolver struct{ *Resolver }

func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	id, _ := strconv.Atoi(obj.UserID)
	return dataloader.GetUserLoader(ctx).Load(id)
}
