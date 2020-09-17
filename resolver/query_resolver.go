package resolver

import "context"

type QueryResolver struct{ *Resolver }

func (r *QueryResolver) Version(ctx context.Context) (string, error) {
	return "1", nil
}
