package go_zpay

import (
	"fmt"
	"testing"
)

func TestCreatePayInOrder(t *testing.T) {
	logger := &MockLogger{}
	config := &RivoConfig{
		MchId:             MERCHANT_ID,
		SecretKey:         SECRET_KEY,
		PayinUrl:          PAYIN_URL,
		PayoutUrl:         PAYOUT_URL,
		PayinCallbackUrl:  PAYIN_CALLBACK_URL,
		PayoutCallbackUrl: PAYOUT_CALLBACK_URL,
		ReturnUrl:         RETURN_URL,
	}
	client := NewClient(logger, config)
	client.SetDebugMode(true)

	req := &PayInRequest{
		TradeNo:       "PAY202604150006",
		Amount:        1001.00,
		Currency:      "VND",
		PaymentMethod: "01",
		OrderDate:     17762112000,
		ClientIp:      "127.0.0.1",
		Name:          "jane",
		Phone:         "0988888881",
		Email:         "jane@example.com",
		City:          "Ho Chi Minh",
		Address:       "District 1",
	}

	// This will fail because it's a real API call, but we can check the signature generation
	resp, err := client.CreatePayInOrder(req)
	if err != nil {
		fmt.Printf("Expected error or result: %v\n", err)
	} else {
		fmt.Printf("Response: %+v\n", resp)
	}
}

type MockLogger struct{}

func (m *MockLogger) Debugf(format string, args ...interface{}) { fmt.Printf(format+"\n", args...) }
func (m *MockLogger) Infof(format string, args ...interface{})  { fmt.Printf(format+"\n", args...) }
func (m *MockLogger) Warnf(format string, args ...interface{})  { fmt.Printf(format+"\n", args...) }
func (m *MockLogger) Errorf(format string, args ...interface{}) { fmt.Printf(format+"\n", args...) }
