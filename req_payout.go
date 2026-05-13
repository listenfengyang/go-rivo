package rivo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CreatePayOutOrder creates a pay-out order.
func (cli *Client) CreatePayOutOrder(req PayOutRequest) (*PayOutResponseData, error) {
	if strings.TrimSpace(req.TradeNo) == "" {
		return nil, fmt.Errorf("tradeNo must not be blank")
	}
	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}
	if strings.TrimSpace(req.Currency) == "" {
		return nil, fmt.Errorf("currency must not be blank")
	}
	if strings.TrimSpace(req.AccType) == "" {
		req.AccType = "01"
	}
	if strings.TrimSpace(req.PaymentMethod) == "" {
		req.PaymentMethod = "BankTransfer"
	}
	if strings.TrimSpace(req.AccountName) == "" {
		return nil, fmt.Errorf("accountName must not be blank")
	}
	if strings.TrimSpace(req.AccountNoUnified) == "" {
		return nil, fmt.Errorf("accountNoUnified must not be blank")
	}
	if strings.TrimSpace(req.BankCode) == "" {
		return nil, fmt.Errorf("bankCode must not be blank")
	}
	if strings.TrimSpace(req.BankName) == "" {
		return nil, fmt.Errorf("bankName must not be blank")
	}

	req.MchId = cli.Params.MchId
	req.SignType = SignTypeSHA512
	if req.Version == "" {
		req.Version = Version10
	}
	req.NotifyUrl = cli.Params.PayoutCallbackUrl

	if err := normalizeAntiReplay(req.Version, &req.Timestamp, &req.Nonce); err != nil {
		return nil, err
	}

	sign, err := signPayload(req, cli.Params.SecretKey)
	if err != nil {
		return nil, err
	}
	req.Sign = sign

	var resp BaseResponse
	_, err = cli.ryClient.R().
		SetHeaders(getHeaders()).
		SetBody(req).
		SetResult(&resp).
		Post(cli.Config.PayoutUrl)
	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("rivo payout failed: code=%s msg=%s", resp.Code, resp.Msg)
	}

	var data PayOutResponseData
	if len(resp.Data) > 0 && string(resp.Data) != "null" {
		if err = json.Unmarshal(resp.Data, &data); err != nil {
			return nil, err
		}
	}
	if data.Sign != "" {
		ok, verifyErr := verifyPayloadSign(data, cli.Params.SecretKey, data.Sign)
		if verifyErr != nil {
			return nil, verifyErr
		}
		if !ok {
			return nil, fmt.Errorf("rivo payout response signature verification failed")
		}
	}

	return &data, nil
}

// QueryPayOutOrder queries a pay-out order.
func (cli *Client) QueryPayOutOrder(req PayOutQueryRequest) (*PayOutQueryResponseData, error) {
	if err := normalizePayOutQueryRequest(&req); err != nil {
		return nil, err
	}

	req.MchId = cli.Params.MchId
	if req.Version == "" {
		req.Version = Version10
	}

	if err := normalizeAntiReplay(req.Version, &req.Timestamp, &req.Nonce); err != nil {
		return nil, err
	}

	sign, err := signPayload(req, cli.Params.SecretKey)
	if err != nil {
		return nil, err
	}
	req.Sign = sign

	queryParams, err := structToSignData(req)
	if err != nil {
		return nil, err
	}
	queryParams["sign"] = req.Sign

	var resp BaseResponse
	_, err = cli.ryClient.R().
		SetQueryParams(queryParams).
		SetResult(&resp).
		Get(cli.Params.payoutQueryURL())
	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("rivo payout query failed: code=%s msg=%s", resp.Code, resp.Msg)
	}

	var data PayOutQueryResponseData
	if len(resp.Data) > 0 && string(resp.Data) != "null" {
		if err = json.Unmarshal(resp.Data, &data); err != nil {
			return nil, err
		}
	}
	if data.Sign != "" {
		ok, verifyErr := verifyPayloadSign(data, cli.Params.SecretKey, data.Sign)
		if verifyErr != nil {
			return nil, verifyErr
		}
		if !ok {
			return nil, fmt.Errorf("rivo payout query response signature verification failed")
		}
	}

	return &data, nil
}

func normalizePayOutQueryRequest(req *PayOutQueryRequest) error {
	if req.MchOrderId == "" && req.TransNo == "" {
		if req.TradeNo != "" {
			req.MchOrderId = req.TradeNo
		}
		if req.OutTradeNo != "" {
			req.TransNo = req.OutTradeNo
		}
	}
	// Deprecated alias fields are not part of official query params/sign payload.
	req.TradeNo = ""
	req.OutTradeNo = ""
	if req.MchOrderId == "" && req.TransNo == "" {
		return fmt.Errorf("mchOrderId or transNo must not be blank")
	}
	return nil
}

// ParsePayOutCallback parses and verifies a pay-out callback.
func (cli *Client) ParsePayOutCallback(callback PayOutCallback) (*PayOutCallback, error) {
	ok, err := verifyPayloadSign(callback, cli.Params.SecretKey, callback.Sign)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("rivo payout callback signature verification failed")
	}
	return &callback, nil
}
