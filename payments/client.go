package payments

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp-demoapp/public-api/models"
)

// HTTPClient is a concrete implmentation of an HTTP client which can communicate with the payments service
type HTTPClient struct {
	baseURL string
}

func NewHTTP(baseURL string) *HTTPClient {
	return &HTTPClient{baseURL: baseURL}
}

// MakePayment calls the payments api
func (h *HTTPClient) MakePayment(details *models.PaymentDetails) (bool, error) {
	pr := &PaymentRequest{}
	pr.FromModel(details)

	resp, err := http.DefaultClient.Post(
		h.baseURL,
		"application/json",
		pr,
	)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	return true, nil
}

// PaymentRequest defines the JSON request for the Payments API
type PaymentRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Number string `json:"number"`
	Expiry string `json:"expiry"`
	CVC    string `json:"cvc"`

	readIndex int64
	buffer    []byte
}

// FromModel converts a graphql model into a payment request
func (pr *PaymentRequest) FromModel(m *models.PaymentDetails) {
	pr.CVC = fmt.Sprintf("%d", m.Cv2)
	pr.Expiry = m.Expiry
	pr.Name = m.Name
	pr.Number = m.Number
	pr.Type = m.Type
}

func (pr *PaymentRequest) Read(p []byte) (n int, err error) {
	// if this is first read marshal the struct
	if pr.readIndex == 0 {
		pr.buffer, err = json.Marshal(pr)
		if err != nil {
			return 0, err
		}
	}

	// we have read all there is to read, reset
	if pr.readIndex >= int64(len(pr.buffer)) {
		err = io.EOF
		pr.readIndex = 0
		return
	}

	n = copy(p, pr.buffer[pr.readIndex:])
	pr.readIndex += int64(n)
	return
}
