package rivo

import "strings"

const (
	DefaultPayInURL       = "https://api.toprivo.com/gateway/pay/first/order/create"
	DefaultPayInQueryURL  = "https://api.toprivo.com/gateway/pay/first/order/query"
	DefaultPayOutURL      = "https://api.toprivo.com/gateway/payout/first/order/create"
	DefaultPayOutQueryURL = "https://api.toprivo.com/gateway/payout/first/order/query"

	// Example-only values for quick bootstrap (not production credentials).
	MERCHANT_ID         = "8822871771"
	SECRET_KEY          = "50c7da45c3414e788f684e8e085b30f2"
	PAYIN_URL           = DefaultPayInURL
	PAYIN_QUERY_URL     = DefaultPayInQueryURL
	PAYOUT_URL          = DefaultPayOutURL
	PAYOUT_QUERY_URL    = DefaultPayOutQueryURL
	PAYIN_CALLBACK_URL  = "https://api-test.logtec.dev/fapi/cpti/payment/psp/public/rivo/deposit/back"
	PAYOUT_CALLBACK_URL = "https://api-test.logtec.dev/fapi/cpti/payment/psp/public/rivo/withdraw/back"
	RETURN_URL          = "https://cpt.supermarkets.com"
)

type RivoConfig struct {
	MchId             string `json:"mchId" mapstructure:"mchId" yaml:"mchId"`
	SecretKey         string `json:"secretKey" mapstructure:"secretKey" yaml:"secretKey"`
	PayinUrl          string `json:"payinUrl" mapstructure:"payinUrl" yaml:"payinUrl"`
	PayinQueryUrl     string `json:"payinQueryUrl" mapstructure:"payinQueryUrl" yaml:"payinQueryUrl"`
	PayoutUrl         string `json:"payoutUrl" mapstructure:"payoutUrl" yaml:"payoutUrl"`
	PayoutQueryUrl    string `json:"payoutQueryUrl" mapstructure:"payoutQueryUrl" yaml:"payoutQueryUrl"`
	PayinCallbackUrl  string `json:"payinCallbackUrl" mapstructure:"payinCallbackUrl" yaml:"payinCallbackUrl"`
	PayoutCallbackUrl string `json:"payoutCallbackUrl" mapstructure:"payoutCallbackUrl" yaml:"payoutCallbackUrl"`
	ReturnUrl         string `json:"returnUrl" mapstructure:"returnUrl" yaml:"returnUrl"`
}

// RivoInitParams is kept for compatibility with existing payment-service-v2 config structs.
type RivoInitParams = RivoConfig

func (c *RivoConfig) normalize() {
	if strings.TrimSpace(c.PayinUrl) == "" {
		c.PayinUrl = DefaultPayInURL
	}
	if strings.TrimSpace(c.PayinQueryUrl) == "" {
		c.PayinQueryUrl = DefaultPayInQueryURL
	}
	if strings.TrimSpace(c.PayoutUrl) == "" {
		c.PayoutUrl = DefaultPayOutURL
	}
	if strings.TrimSpace(c.PayoutQueryUrl) == "" {
		c.PayoutQueryUrl = DefaultPayOutQueryURL
	}
}

func (c *RivoConfig) payinQueryURL() string {
	if strings.TrimSpace(c.PayinQueryUrl) != "" {
		return c.PayinQueryUrl
	}
	return c.PayinUrl
}

func (c *RivoConfig) payoutQueryURL() string {
	if strings.TrimSpace(c.PayoutQueryUrl) != "" {
		return c.PayoutQueryUrl
	}
	return c.PayoutUrl
}
