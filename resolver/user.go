package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *QueryResolver) User(ctx context.Context, userID string) (*models.User, error) {
	return nil, nil
}

func (r *QueryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return []*models.User{}, nil
}
