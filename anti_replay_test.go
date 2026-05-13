package rivo

import "testing"

func TestNormalizeAntiReplay_Version10PairRule(t *testing.T) {
	ts := int64(1776211200000)
	nonce := ""
	if err := normalizeAntiReplay(Version10, &ts, &nonce); err == nil {
		t.Fatalf("expected error when only timestamp is provided")
	}

	ts = 0
	nonce = "abc"
	if err := normalizeAntiReplay(Version10, &ts, &nonce); err == nil {
		t.Fatalf("expected error when only nonce is provided")
	}

	ts = 0
	nonce = ""
	if err := normalizeAntiReplay(Version10, &ts, &nonce); err != nil {
		t.Fatalf("unexpected error for both empty: %v", err)
	}
}

func TestNormalizeAntiReplay_Version11AutoFill(t *testing.T) {
	ts := int64(0)
	nonce := ""
	if err := normalizeAntiReplay(Version11, &ts, &nonce); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts <= 0 {
		t.Fatalf("timestamp should be auto-filled for version 1.1")
	}
	if nonce == "" {
		t.Fatalf("nonce should be auto-filled for version 1.1")
	}
}

func TestNormalizePayOutQueryRequest_DeprecatedAlias(t *testing.T) {
	req := PayOutQueryRequest{
		TradeNo:    "PO202604150001",
		OutTradeNo: "WD202604150000001",
	}
	if err := normalizePayOutQueryRequest(&req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.MchOrderId != "PO202604150001" || req.TransNo != "WD202604150000001" {
		t.Fatalf("alias mapping failed: %+v", req)
	}
	if req.TradeNo != "" || req.OutTradeNo != "" {
		t.Fatalf("deprecated alias fields should be cleared before signing/query: %+v", req)
	}
}
