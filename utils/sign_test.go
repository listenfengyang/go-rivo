package utils

import (
	"testing"
)

func TestGenerateSign(t *testing.T) {
	params := map[string]interface{}{
		"address":       "District 1",
		"amount":        "100000.00",
		"city":          "Ho Chi Minh",
		"clientIp":      "203.0.113.10",
		"currency":      "VND",
		"email":         "buyer@example.com",
		"extraData":     `{"goodsId":"SKU1001","note":"vip"}`,
		"mchId":         "MCH10001",
		"name":          "NGUYEN VAN A",
		"nonce":         "PAYNONCE001",
		"notifyUrl":     "https://merchant.example.com/api/rivo/pay/notify",
		"orderDate":     int64(1776211200000),
		"paymentMethod": "01",
		"phone":         "0988888888",
		"returnUrl":     "https://merchant.example.com/pay/result",
		"signType":      "SHA512",
		"timestamp":     int64(1776211200123),
		"tradeNo":       "PAY202604150001",
		"version":       "1.0",
	}
	key := "demo_secret_key_123456"
	expectedSign := "e72941ccf60957f9c0d5b03e100d00d807d7dbddd02ac1aafaa4dab3e930eb0212bd1bb95037105eea77685444a2c370a1f1729da4edf610c4f1a46511adff07"

	actualSign := GenerateSign(params, key)
	if actualSign != expectedSign {
		t.Errorf("GenerateSign failed, expected: %s, actual: %s", expectedSign, actualSign)
	}
}
