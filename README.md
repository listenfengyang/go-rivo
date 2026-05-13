# go-rivo

Rivo Global payment SDK for Go.

## Installation

```bash
go get github.com/listenfengyang/go-rivo
```

If you migrate this repo into another organization (for example `github.com/asaka1234/go-rivo`), update `go.mod` `module` path before publishing.

## Features

- Pay-In order create and query
- Pay-Out order create and query
- Callback signature verification
- SHA512 signing utility

## Quick Start

```go
package main

import (
	"log"

	rivo "github.com/listenfengyang/go-rivo"
)

type noopLogger struct{}

func (n *noopLogger) Debugf(string, ...interface{}) {}
func (n *noopLogger) Infof(string, ...interface{})  {}
func (n *noopLogger) Warnf(string, ...interface{})  {}
func (n *noopLogger) Errorf(string, ...interface{}) {}

func main() {
	cfg := &rivo.RivoInitParams{
		MchId:             "your_mch_id",
		SecretKey:         "your_secret_key",
		PayinCallbackUrl:  "https://your-domain.com/payment/psp/public/rivo/deposit/back",
		PayoutCallbackUrl: "https://your-domain.com/payment/psp/public/rivo/withdraw/back",
		ReturnUrl:         "https://your-domain.com/cashier/return",
	}
	client := rivo.NewClient(&noopLogger{}, cfg)

	resp, err := client.CreatePayInOrder(rivo.PayInRequest{
		TradeNo:       "PAY202605130001",
		Amount:        1000,
		Currency:      "VND",
		PaymentMethod: "01",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("cashier url: %s", resp.CashierUrl)
}
```

## Notes

- Default endpoints:
  - `https://api.toprivo.com/gateway/pay/first/order/create`
  - `https://api.toprivo.com/gateway/payout/first/order/create`
- Query endpoints from official docs:
  - `GET /gateway/pay/first/order/query`
  - `GET /gateway/payout/first/order/query` (uses `mchOrderId` / `transNo`)
- Anti-replay rule:
  - `version=1.1`: `timestamp` and `nonce` are required (SDK auto-fills if omitted).
  - `version=1.0`: `timestamp` and `nonce` must be both provided or both omitted.
- Callback handling: merchant side should return plain text `ok` after successful processing.
- Constants `MERCHANT_ID` / `SECRET_KEY` / callback URLs are example placeholders only.
