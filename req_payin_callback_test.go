package go_zpay

import (
	"testing"
)

func TestParsePayInCallback(t *testing.T) {
	config := &RivoConfig{
		MchId:             MERCHANT_ID,
		SecretKey:         SECRET_KEY,
		PayinUrl:          PAYIN_URL,
		PayoutUrl:         PAYOUT_URL,
		PayinCallbackUrl:  PAYIN_CALLBACK_URL,
		PayoutCallbackUrl: PAYOUT_CALLBACK_URL,
		ReturnUrl:         RETURN_URL,
	}
	client := NewClient(&MockLogger{}, config)

	callback, err := client.ParsePayInCallback(MakeDemoRequest())
	if err != nil {
		t.Fatalf("ParsePayInCallback failed: %v", err)
	}

	t.Logf("callback: %v\n", callback)
}

// {\"mchId\":\"8822871771\",\"tradeNo\":\"PAY202604150006\",\"outTradeNo\":\"PAY2605092052970940983377920\",\"amount\":\"1001.00\",\"pidAmount\":\"1001.00\",\"currency\":\"VND\",\"orderDate\":1778301389000,\"payTime\":1778276314000,\"payStatus\":1,\"payername\":\"jane\",\"sign\":\"b75c98ffcb405a1f8c766f4a6e22e3d87e8ebaf5a2f20e04f72b24f83ab73563ea3f1edc72a9040f84117cf1819607be120b355b1f7b003210134742a77fc7b4\"}
func MakeDemoRequest() PayInCallback {
	return PayInCallback{
		MchId:      "8822871771",
		TradeNo:    "PAY202604150006",
		OutTradeNo: "PAY2605092052970940983377920",
		Amount:     "1001.00",
		PidAmount:  "1001.00",
		Currency:   "VND",
		OrderDate:  1778301389000,
		PayTime:    1778276314000,
		PayStatus:  1,
		PayerName:  "jane",
		Sign:       "b75c98ffcb405a1f8c766f4a6e22e3d87e8ebaf5a2f20e04f72b24f83ab73563ea3f1edc72a9040f84117cf1819607be120b355b1f7b003210134742a77fc7b4",
	}
}
