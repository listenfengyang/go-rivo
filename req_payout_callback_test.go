package rivo

import "testing"

func TestParsePayOutCallback(t *testing.T) {
	client := newTestClient()

	callback := PayOutCallback{
		MchId:      "8822871771",
		TradeNo:    "PO202604150001",
		OutTradeNo: "PO2605092052970940983377920",
		Amount:     StringValue("1000.00"),
		Currency:   "VND",
		Status:     1,
		OrderDate:  1778301389000,
		PayTime:    1778276314000,
	}

	sign, err := signPayload(callback, client.Config.SecretKey)
	if err != nil {
		t.Fatalf("signPayload failed: %v", err)
	}
	callback.Sign = sign

	parsed, err := client.ParsePayOutCallback(callback)
	if err != nil {
		t.Fatalf("ParsePayOutCallback failed: %v", err)
	}

	if parsed.TradeNo != callback.TradeNo {
		t.Fatalf("unexpected tradeNo: got %s want %s", parsed.TradeNo, callback.TradeNo)
	}
}
