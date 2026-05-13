package rivo

import (
	"os"
	"testing"
)

func TestCreatePayOutOrder_Integration(t *testing.T) {
	if os.Getenv("RIVO_RUN_INTEGRATION") != "1" {
		t.Skip("set RIVO_RUN_INTEGRATION=1 to run real gateway integration tests")
	}

	client := newTestClient()
	client.SetDebugModel(true)

	_, err := client.CreatePayOutOrder(PayOutRequest{
		TradeNo:          "PO202605130001",
		Amount:           1000.00,
		Currency:         "VND",
		AccType:          "01",
		PaymentMethod:    "BankTransfer",
		AccountName:      "BENEFICIARY NAME",
		AccountNoUnified: "1234567890",
		BankCode:         "VCB",
		BankName:         "Vietcombank",
	})
	if err != nil {
		t.Fatalf("CreatePayOutOrder failed: %v", err)
	}
}
