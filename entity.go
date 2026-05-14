package rivo

import (
	"encoding/json"
	"strings"
)

type BaseResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

// StringValue supports both JSON string and number fields.
type StringValue string

func (s *StringValue) UnmarshalJSON(data []byte) error {
	txt := strings.TrimSpace(string(data))
	if txt == "" || txt == "null" {
		*s = ""
		return nil
	}
	if strings.HasPrefix(txt, "\"") {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*s = StringValue(v)
		return nil
	}
	*s = StringValue(txt)
	return nil
}

func (s StringValue) String() string {
	return string(s)
}

// Collection (Pay-In)
type PayInRequest struct {
	MchId         string  `json:"mchId,omitempty"`
	TradeNo       string  `json:"tradeNo"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"paymentMethod,omitempty"`
	OrderDate     int64   `json:"orderDate,omitempty"`
	ClientIp      string  `json:"clientIp,omitempty"`
	Name          string  `json:"name,omitempty"`
	Phone         string  `json:"phone,omitempty"`
	Email         string  `json:"email,omitempty"`
	City          string  `json:"city,omitempty"`
	Address       string  `json:"address,omitempty"`
	NotifyUrl     string  `json:"notifyUrl,omitempty"`
	ReturnUrl     string  `json:"returnUrl,omitempty"`
	Subject       string  `json:"subject,omitempty"`
	ExtraData     string  `json:"extraData,omitempty"`
	SignType      string  `json:"signType,omitempty"`
	Version       string  `json:"version,omitempty"`
	Timestamp     int64   `json:"timestamp,omitempty"`
	Nonce         string  `json:"nonce,omitempty"`
	Sign          string  `json:"sign,omitempty"`
}

type PayInResponseData struct {
	MchId      string      `json:"mchId"`
	TradeNo    string      `json:"tradeNo"`
	OutTradeNo string      `json:"outTradeNo"`
	Amount     StringValue `json:"amount"`
	OrderDate  int64       `json:"orderDate"`
	CashierUrl string      `json:"cashierUrl"`
	Qrcode     string      `json:"qrcode"`
	PayStatus  int         `json:"payStatus"`
	Sign       string      `json:"sign"`
}

type PayInQueryRequest struct {
	MchId      string `json:"mchId,omitempty"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"outTradeNo,omitempty"`
	Version    string `json:"version,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	Nonce      string `json:"nonce,omitempty"`
	Sign       string `json:"sign,omitempty"`
}

type PayInQueryResponseData struct {
	MchId      string      `json:"mchId"`
	TradeNo    string      `json:"tradeNo"`
	OutTradeNo string      `json:"outTradeNo"`
	Amount     StringValue `json:"amount"`
	OrderDate  int64       `json:"orderDate"`
	CashierUrl string      `json:"cashierUrl"`
	Qrcode     string      `json:"qrcode"`
	PayTime    int64       `json:"payTime"`
	PayStatus  int         `json:"payStatus"`
	Sign       string      `json:"sign"`
}

type PayInCallback struct {
	MchId         string      `json:"mchId"`
	TradeNo       string      `json:"tradeNo"`
	OutTradeNo    string      `json:"outTradeNo"`
	Amount        StringValue `json:"amount"`
	PidAmount     StringValue `json:"pidAmount"`
	Currency      string      `json:"currency"`
	OrderDate     int64       `json:"orderDate"`
	PayTime       int64       `json:"payTime"`
	PayStatus     int         `json:"payStatus"`
	PayerName     string      `json:"payername"`
	PayerAccount  string      `json:"payeraccount"`
	PayerBankCode string      `json:"payerbankcode"`
	ExtraData     string      `json:"extraData,omitempty"`
	Sign          string      `json:"sign"`
}

// Disbursement (Pay-Out)
type PayOutRequest struct {
	MchId            string  `json:"mchId,omitempty"`
	TradeNo          string  `json:"tradeNo"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	AccType          string  `json:"accType,omitempty"`
	PaymentMethod    string  `json:"paymentMethod,omitempty"`
	AccountName      string  `json:"accountName,omitempty"`
	AccountNoUnified string  `json:"accountNoUnified,omitempty"`
	BankCode         string  `json:"bankCode,omitempty"`
	BankName         string  `json:"bankName,omitempty"`
	AccTel           string  `json:"accTel,omitempty"`
	AccIdCard        string  `json:"accIdCard,omitempty"`
	Purpose          string  `json:"purpose,omitempty"`
	NotifyUrl        string  `json:"notifyUrl,omitempty"`
	ExtraData        string  `json:"extraData,omitempty"`
	SignType         string  `json:"signType,omitempty"`
	Version          string  `json:"version,omitempty"`
	Timestamp        int64   `json:"timestamp,omitempty"`
	Nonce            string  `json:"nonce,omitempty"`
	Sign             string  `json:"sign,omitempty"`
}

type PayOutResponseData struct {
	MchId      string      `json:"mchId"`
	TradeNo    string      `json:"tradeNo"`
	OutTradeNo string      `json:"outTradeNo"`
	Amount     StringValue `json:"amount"`
	OrderDate  int64       `json:"orderDate"`
	Status     int         `json:"status"`
	Sign       string      `json:"sign"`
}

type PayOutQueryRequest struct {
	MchId       string `json:"mchId,omitempty"`
	MchOrderId  string `json:"mchOrderId,omitempty"`
	TransNo     string `json:"transNo,omitempty"`
	ChannelCode string `json:"channelCode,omitempty"`
	Version     string `json:"version,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	Nonce       string `json:"nonce,omitempty"`
	Sign        string `json:"sign,omitempty"`

	// Deprecated aliases for backward compatibility.
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"outTradeNo,omitempty"`
}

type PayOutQueryResponseData struct {
	MchId      string      `json:"mchId"`
	TradeNo    string      `json:"tradeNo"`
	OutTradeNo string      `json:"outTradeNo"`
	Amount     StringValue `json:"amount"`
	OrderDate  int64       `json:"orderDate"`
	Status     int         `json:"status"`
	Sign       string      `json:"sign"`
}

type PayOutCallback struct {
	MchId      string      `json:"mchId"`
	TradeNo    string      `json:"tradeNo"`
	OutTradeNo string      `json:"outTradeNo"`
	Amount     StringValue `json:"amount"`
	Currency   string      `json:"currency"`
	Status     int         `json:"status"`
	OrderDate  int64       `json:"orderDate"`
	PayTime    int64       `json:"payTime,omitempty"`
	ExtraData  StringValue `json:"extraData,omitempty"`
	Sign       string      `json:"sign"`
}
