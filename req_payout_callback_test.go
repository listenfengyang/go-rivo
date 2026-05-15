package rivo

import (
	"encoding/json"
	"testing"
)

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

func TestParsePayOutCallback_WithoutPayTime(t *testing.T) {
	client := newTestClient()

	callback := PayOutCallback{
		MchId:      "8822871771",
		TradeNo:    "PO202604150002",
		OutTradeNo: "PO2605092052970940983377921",
		Amount:     StringValue("1000.00"),
		Currency:   "VND",
		Status:     1,
		OrderDate:  1778301389000,
	}

	sign, err := signPayload(callback, client.Config.SecretKey)
	if err != nil {
		t.Fatalf("signPayload failed: %v", err)
	}
	callback.Sign = sign

	if _, err := client.ParsePayOutCallback(callback); err != nil {
		t.Fatalf("ParsePayOutCallback failed without payTime: %v", err)
	}
}

func TestPayOutCallback_UnmarshalObjectExtraData(t *testing.T) {
	raw := []byte(`{
		"mchId":"8822871771",
		"tradeNo":"PO202604150003",
		"outTradeNo":"PO2605092052970940983377922",
		"amount":1000.00,
		"currency":"VND",
		"orderDate":1778301389000,
		"status":1,
		"extraData":{"batchNo":"BATCH001"}
	}`)

	var callback PayOutCallback
	if err := json.Unmarshal(raw, &callback); err != nil {
		t.Fatalf("unmarshal payout callback failed: %v", err)
	}
	if callback.ExtraData.String() == "" {
		t.Fatal("extraData should keep object payload as signable string")
	}
}

func TestParsePayOutCallback_InvalidSign(t *testing.T) {
	client := newTestClient()

	callback := PayOutCallback{
		MchId:      "8822871771",
		TradeNo:    "PO202604150004",
		OutTradeNo: "PO2605092052970940983377923",
		Amount:     StringValue("1000.00"),
		Currency:   "VND",
		Status:     1,
		OrderDate:  1778301389000,
		Sign:       "invalid-sign",
	}

	if _, err := client.ParsePayOutCallback(callback); err == nil {
		t.Fatal("expected signature verification error for invalid sign")
	}
}
