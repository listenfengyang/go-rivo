package go_zpay

const (
	MERCHANT_ID = "8822871771"
	SECRET_KEY  = "50c7da45c3414e788f684e8e085b30f2"
	// "https://api.rivoglobal.io"
	PAYIN_URL           = "https://api.toprivo.com/gateway/pay/first/order/create"
	PAYOUT_URL          = "https://api.toprivo.com/gateway/payout/first/order/create"
	PAYIN_CALLBACK_URL  = "https://api-test.logtec.dev/fapi/payment/psp/public/zpay/deposit/back"
	PAYOUT_CALLBACK_URL = "https://api-test.logtec.dev/fapi/payment/psp/public/zpay/withdraw/back"
	RETURN_URL          = "https://cpt.supermarkets.com"
)

type RivoConfig struct {
	MchId             string `json:"mchId" yaml:"mchId"`
	SecretKey         string `json:"secretKey" yaml:"secretKey"`
	PayinUrl          string `json:"payinUrl" yaml:"payinUrl"`
	PayoutUrl         string `json:"payoutUrl" yaml:"payoutUrl"`
	PayinCallbackUrl  string `json:"payinCallbackUrl" yaml:"payinCallbackUrl"`
	PayoutCallbackUrl string `json:"payoutCallbackUrl" yaml:"payoutCallbackUrl"`
	ReturnUrl         string `json:"returnUrl" yaml:"returnUrl"`
}
