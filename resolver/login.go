package resolver

import "context"

func (r *MutationResolver) Login(ctx context.Context) (bool, error) {
	return true, nil
}
