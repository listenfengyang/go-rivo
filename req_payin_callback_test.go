package rivo

import "testing"

func TestParsePayInCallback(t *testing.T) {
	client := newTestClient()

	callback := PayInCallback{
		MchId:      "8822871771",
		TradeNo:    "PAY202604150006",
		OutTradeNo: "PAY2605092052970940983377920",
		Amount:     StringValue("1001.00"),
		PidAmount:  StringValue("1001.00"),
		Currency:   "VND",
		OrderDate:  1778301389000,
		PayTime:    1778276314000,
		PayStatus:  1,
		PayerName:  "jane",
	}

	sign, err := signPayload(callback, client.Config.SecretKey)
	if err != nil {
		t.Fatalf("signPayload failed: %v", err)
	}
	callback.Sign = sign

	parsed, err := client.ParsePayInCallback(callback)
	if err != nil {
		t.Fatalf("ParsePayInCallback failed: %v", err)
	}

	if parsed.TradeNo != callback.TradeNo {
		t.Fatalf("unexpected tradeNo: got %s want %s", parsed.TradeNo, callback.TradeNo)
	}
}

func TestParsePayInCallback_InvalidSign(t *testing.T) {
	client := newTestClient()

	callback := PayInCallback{
		MchId:      "8822871771",
		TradeNo:    "PAY202604150007",
		OutTradeNo: "PAY2605092052970940983377921",
		Amount:     StringValue("1001.00"),
		Currency:   "VND",
		OrderDate:  1778301389000,
		PayTime:    1778276314000,
		PayStatus:  1,
		Sign:       "invalid-sign",
	}

	if _, err := client.ParsePayInCallback(callback); err == nil {
		t.Fatal("expected signature verification error for invalid sign")
	}
}
