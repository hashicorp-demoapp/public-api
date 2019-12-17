package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *queryResolver) User(ctx context.Context, userID string) (*models.User, error) {
	return nil, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return []*models.User{}, nil
}