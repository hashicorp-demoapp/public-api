package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/auth"
	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *MutationResolver) SignUp(ctx context.Context, auth models.UserAuth) (*models.AuthResponse, error) {
	authResp, err := r.ProductService.SignUp(auth)
	if err != nil {
		return nil, err
	}
	return &authResp, nil
}

func (r *MutationResolver) Login(ctx context.Context, auth models.UserAuth) (*models.AuthResponse, error) {
	authResp, err := r.ProductService.SignIn(auth)
	if err != nil {
		return nil, err
	}
	return &authResp, nil
}

func (r *MutationResolver) SignOut(ctx context.Context) (bool, error) {
	authToken := auth.GetAuthHeader(ctx)

	err := r.ProductService.SignOut(&authToken)
	if err != nil {
		return false, err
	}
	return true, nil
}
