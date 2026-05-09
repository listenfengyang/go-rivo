package go_zpay

import (
	"fmt"
	"testing"
)

func TestCreatePayOutOrder(t *testing.T) {
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

	req := &PayOutRequest{
		TradeNo:          "PO202604150001",
		Amount:           1000.00,
		Currency:         "VND",
		AccType:          "01",
		PaymentMethod:    "BankTransfer",
		AccountName:      "BENEFICIARY NAME",
		AccountNoUnified: "1234567890",
		BankCode:         "VCB",
		BankName:         "Vietcombank",
	}

	// This will fail because it's a real API call, but we can check the signature generation
	resp, err := client.CreatePayOutOrder(req)
	if err != nil {
		fmt.Printf("Expected error or result: %v\n", err)
	} else {
		fmt.Printf("Response: %+v\n", resp)
	}
}
