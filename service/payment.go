package service

import (
	"github.com/hashicorp-demoapp/public-api/models"
	"github.com/hashicorp-demoapp/public-api/payments"
)

// PaymentService handles interactions with the payments api
type PaymentService struct {
	client *payments.HTTPClient
}

// NewPaymentService returns an instance of PaymentService
func NewPaymentService(c *payments.HTTPClient) *PaymentService {
	return &PaymentService{c}
}

func (p *PaymentService) Pay(pd *models.PaymentDetails) (bool, error) {
	return p.client.MakePayment(pd)
}
