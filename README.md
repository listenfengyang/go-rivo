# go-rivo
Rivo Global Payment SDK for Go.

## Features
- Pay-In (Collection) integration.
- Pay-Out (Disbursement) integration.
- Order Query for both Pay-In and Pay-Out.
- SHA-512 Signature generation and verification.
- Async Callback parsing and validation.

## Installation
```bash
go get github.com/listenfengyang/go-zpay
```

## Usage

### Initialize Client
```go
config := &go_rivo.RivoConfig{
    MchId:     "your_mch_id",
    SecretKey: "your_secret_key",
    BaseUrl:   go_rivo.BASE_URL_TEST,
}
client := go_rivo.NewClient(logger, config)
```

### Create Pay-In Order
```go
req := &go_rivo.PayInRequest{
    TradeNo:       "order_123",
    Amount:        decimal.NewFromFloat(100.0),
    Currency:      "VND",
    PaymentMethod: "01",
    OrderDate:     time.Now().UnixNano() / 1e6,
    // ... other fields
}
resp, err := client.CreatePayInOrder(req)
```
