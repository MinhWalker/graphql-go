package resolver

import (
	"context"
	"graphQL-go/graph/generated"
	"graphQL-go/models"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}
