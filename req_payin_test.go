package rivo

import (
	"os"
	"testing"
)

type mockLogger struct{}

func (m *mockLogger) Debugf(string, ...interface{}) {}
func (m *mockLogger) Infof(string, ...interface{})  {}
func (m *mockLogger) Warnf(string, ...interface{})  {}
func (m *mockLogger) Errorf(string, ...interface{}) {}

func newTestClient() *Client {
	cfg := &RivoConfig{
		MchId:             "8822871771",
		SecretKey:         "50c7da45c3414e788f684e8e085b30f2",
		PayinCallbackUrl:  "https://api-test.logtec.dev/fapi/payment/psp/public/rivo/deposit/back",
		PayoutCallbackUrl: "https://api-test.logtec.dev/fapi/payment/psp/public/rivo/withdraw/back",
		ReturnUrl:         "https://cpt.supermarkets.com",
	}
	return NewClient(&mockLogger{}, cfg)
}

func TestCreatePayInOrder_Integration(t *testing.T) {
	if os.Getenv("RIVO_RUN_INTEGRATION") != "1" {
		t.Skip("set RIVO_RUN_INTEGRATION=1 to run real gateway integration tests")
	}

	client := newTestClient()
	client.SetDebugModel(true)

	_, err := client.CreatePayInOrder(PayInRequest{
		TradeNo:       "PAY202605130001",
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
	})
	if err != nil {
		t.Fatalf("CreatePayInOrder failed: %v", err)
	}
}
