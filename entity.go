package go_zpay

import "github.com/shopspring/decimal"

// Common Response Structure
type BaseResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Collection (Pay-In)
type PayInRequest struct {
	MchId         string  `json:"mchId" mapstructure:"mchId"`                                            // Merchant ID
	TradeNo       string  `json:"tradeNo" mapstructure:"tradeNo"`                                        // Merchant order number
	Amount        float64 `json:"amount" mapstructure:"amount,omitempty"`                                // Order amount
	Currency      string  `json:"currency" mapstructure:"currency"`                                      // Currency code
	PaymentMethod string  `json:"paymentMethod" mapstructure:"paymentMethod,omitempty"`                  // Payment method code
	OrderDate     int64   `json:"orderDate" mapstructure:"orderDate,omitempty"`                          // Merchant order time, ms timestamp
	ClientIp      string  `json:"clientIp" mapstructure:"clientIp,omitempty"`                            // Payer IP address
	Name          string  `json:"name" mapstructure:"name,omitempty"`                                    // Payer full name
	Phone         string  `json:"phone" mapstructure:"phone,omitempty"`                                  // Payer phone number
	Email         string  `json:"email" mapstructure:"email,omitempty"`                                  // Payer email
	City          string  `json:"city" mapstructure:"city,omitempty"`                                    // Payer city
	Address       string  `json:"address" mapstructure:"address,omitempty"`                              // Payer address
	NotifyUrl     string  `json:"notifyUrl" mapstructure:"notifyUrl,omitempty"`                          // Async notification URL
	ReturnUrl     string  `json:"returnUrl" mapstructure:"returnUrl" mapstructure:"returnUrl,omitempty"` // Return URL
	Subject       string  `json:"subject" mapstructure:"subject,omitempty"`                              // Order subject
	ExtraData     string  `json:"extraData,omitempty" mapstructure:"extraData,omitempty"`
	SignType      string  `json:"signType" mapstructure:"signType,omitempty"` // Fixed value SHA512
	Version       string  `json:"version" mapstructure:"version,omitempty"`   // Fixed value 1.0
	// Timestamp     int64   `json:"timestamp,omitempty" mapstructure:"timestamp" mapstructure:"timestamp,omitempty,omitempty"`
	// Nonce         string  `json:"nonce,omitempty" mapstructure:"nonce" mapstructure:"nonce,omitempty,omitempty"`
	Sign string `json:"sign,omitempty" mapstructure:"sign,omitempty"`
}

type PayInResponseData struct {
	MchId      string `json:"mchId" mapstructure:"mchId"`
	TradeNo    string `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo string `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Amount     string `json:"amount" mapstructure:"amount,omitempty"`
	OrderDate  int64  `json:"orderDate" mapstructure:"orderDate,omitempty"`
	CashierUrl string `json:"cashierUrl" mapstructure:"cashierUrl,omitempty"`
	Qrcode     string `json:"qrcode" mapstructure:"qrcode,omitempty"`
	PayStatus  int    `json:"payStatus" mapstructure:"payStatus,omitempty"`
	Sign       string `json:"sign,omitempty" mapstructure:"sign,omitempty"`
}

type PayInQueryRequest struct {
	MchId      string `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo    string `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo string `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Version    string `json:"version" mapstructure:"version" mapstructure:"version,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty" mapstructure:"timestamp" mapstructure:"timestamp,omitempty" mapstructure:"timestamp,omitempty,omitempty"`
	Nonce      string `json:"nonce,omitempty" mapstructure:"nonce" mapstructure:"nonce,omitempty" mapstructure:"nonce,omitempty,omitempty,omitempty"`
	Sign       string `json:"sign,omitempty"`
}

type PayInQueryResponseData struct {
	MchId      string          `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo    string          `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo string          `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Amount     decimal.Decimal `json:"amount" mapstructure:"amount,omitempty"`
	OrderDate  int64           `json:"orderDate" mapstructure:"orderDate,omitempty"`
	CashierUrl string          `json:"cashierUrl" mapstructure:"cashierUrl,omitempty"`
	Qrcode     string          `json:"qrcode" mapstructure:"qrcode,omitempty"`
	PayTime    int64           `json:"payTime" mapstructure:"payTime,omitempty"`
	PayStatus  int             `json:"payStatus" mapstructure:"payStatus,omitempty"`
	Sign       string          `json:"sign,omitempty"`
}

type PayInCallback struct {
	MchId         string `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo       string `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo    string `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Amount        string `json:"amount" mapstructure:"amount,omitempty"`
	PidAmount     string `json:"pidAmount" mapstructure:"pidAmount,omitempty"`
	Currency      string `json:"currency" mapstructure:"currency,omitempty"`
	OrderDate     int64  `json:"orderDate" mapstructure:"orderDate,omitempty"`
	PayTime       int64  `json:"payTime" mapstructure:"payTime,omitempty"`
	PayStatus     int    `json:"payStatus" mapstructure:"payStatus" mapstructure:"payStatus,omitempty"`
	PayerName     string `json:"payername" mapstructure:"payername,omitempty"`
	PayerAccount  string `json:"payeraccount" mapstructure:"payeraccount,omitempty"`
	PayerBankCode string `json:"payerbankcode" mapstructure:"payerbankcode,omitempty"`
	ExtraData     string `json:"extraData,omitempty" mapstructure:"extraData,omitempty"`
	Sign          string `json:"sign,omitempty" mapstructure:"sign,omitempty"`
}

// Disbursement (Pay-Out)
type PayOutRequest struct {
	MchId            string  `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo          string  `json:"tradeNo" mapstructure:"tradeNo" mapstructure:"tradeNo,omitempty"`
	Amount           float64 `json:"amount" mapstructure:"amount" mapstructure:"amount,omitempty"`
	Currency         string  `json:"currency" mapstructure:"currency" mapstructure:"currency,omitempty"`
	AccType          string  `json:"accType" mapstructure:"accType" mapstructure:"accType,omitempty"`
	PaymentMethod    string  `json:"paymentMethod" mapstructure:"paymentMethod" mapstructure:"paymentMethod,omitempty"`
	AccountName      string  `json:"accountName" mapstructure:"accountName" mapstructure:"accountName,omitempty"`
	AccountNoUnified string  `json:"accountNoUnified" mapstructure:"accountNoUnified" mapstructure:"accountNoUnified,omitempty"`
	BankCode         string  `json:"bankCode" mapstructure:"bankCode" mapstructure:"bankCode,omitempty"`
	BankName         string  `json:"bankName" mapstructure:"bankName" mapstructure:"bankName,omitempty"`
	// AccTel           string  `json:"accTel" mapstructure:"accTel" mapstructure:"accTel,omitempty"`
	// AccIdCard        string  `json:"accIdCard" mapstructure:"accIdCard" mapstructure:"accIdCard,omitempty"`
	// Purpose          string  `json:"purpose" mapstructure:"purpose" mapstructure:"purpose,omitempty"`
	NotifyUrl string `json:"notifyUrl" mapstructure:"notifyUrl" mapstructure:"notifyUrl,omitempty"`
	// ExtraData string `json:"extraData" mapstructure:"extraData" mapstructure:"extraData,omitempty"`
	SignType string `json:"signType" mapstructure:"signType" mapstructure:"signType,omitempty"`
	Version  string `json:"version" mapstructure:"version" mapstructure:"version,omitempty"`
	// Timestamp int64  `json:"timestamp,omitempty" mapstructure:"timestamp" mapstructure:"timestamp,omitempty" mapstructure:"timestamp,omitempty,omitempty,omitempty"`
	// Nonce     string `json:"nonce,omitempty" mapstructure:"nonce" mapstructure:"nonce,omitempty" mapstructure:"nonce,omitempty,omitempty,omitempty"`
	Sign string `json:"sign,omitempty" mapstructure:"sign" mapstructure:"sign,omitempty"`
}

type PayOutResponseData struct {
	MchId      string          `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo    string          `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo string          `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Amount     decimal.Decimal `json:"amount" mapstructure:"amount,omitempty"`
	OrderDate  int64           `json:"orderDate" mapstructure:"orderDate,omitempty"`
	Status     int             `json:"status" mapstructure:"status,omitempty"`
	Sign       string          `json:"sign,omitempty"`
}

type PayOutQueryRequest struct {
	MchId      string `json:"mchId"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"outTradeNo,omitempty"`
	Version    string `json:"version"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	Nonce      string `json:"nonce,omitempty"`
	Sign       string `json:"sign"`
}

type PayOutQueryResponseData struct {
	MchId      string          `json:"mchId"`
	TradeNo    string          `json:"tradeNo"`
	OutTradeNo string          `json:"outTradeNo"`
	Amount     decimal.Decimal `json:"amount"`
	OrderDate  int64           `json:"orderDate"`
	Status     int             `json:"status"`
	Sign       string          `json:"sign"`
}

type PayOutCallback struct {
	MchId      string          `json:"mchId" mapstructure:"mchId,omitempty"`
	TradeNo    string          `json:"tradeNo" mapstructure:"tradeNo,omitempty"`
	OutTradeNo string          `json:"outTradeNo" mapstructure:"outTradeNo,omitempty"`
	Amount     decimal.Decimal `json:"amount" mapstructure:"amount,omitempty"`
	Currency   string          `json:"currency" mapstructure:"currency,omitempty"`
	Status     int             `json:"status" mapstructure:"status,omitempty"`
	OrderDate  int64           `json:"orderDate" mapstructure:"orderDate,omitempty"`
	PayTime    int64           `json:"payTime" mapstructure:"payTime,omitempty"`
	ExtraData  string          `json:"extraData,omitempty" mapstructure:"extraData,omitempty,omitempty"`
	Sign       string          `json:"sign,omitempty"`
}
