package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

// Pay makes a call to the payment service and returns the response
func (r *MutationResolver) Pay(ctx context.Context, details *models.PaymentDetails) (bool, error) {
	r.Log.Info("Received Pay Mutation", "details", details)

	ok, err := r.PaymentService.Pay(details)
	if err != nil {
		r.Log.Error("Unable to make payment", "error", err)
	}

	return ok, err
}
