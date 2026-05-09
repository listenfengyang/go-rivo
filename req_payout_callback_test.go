package go_zpay

import (
	"testing"
)

func TestParsePayOutCallback(t *testing.T) {
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

	callback, err := client.ParsePayOutCallback(MakeDemoWdRequest())
	if err != nil {
		t.Fatalf("ParsePayOutCallback failed: %v", err)
	}

	t.Logf("wd callback: %+v", callback)
}

func MakeDemoWdRequest() PayOutCallback {
	return PayOutCallback{
		Sign: "",
	}
}
