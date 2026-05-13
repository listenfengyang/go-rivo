package rivo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// CreatePayInOrder creates a pay-in order.
func (cli *Client) CreatePayInOrder(req PayInRequest) (*PayInResponseData, error) {
	if strings.TrimSpace(req.TradeNo) == "" {
		return nil, fmt.Errorf("tradeNo must not be blank")
	}
	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}
	if strings.TrimSpace(req.Currency) == "" {
		return nil, fmt.Errorf("currency must not be blank")
	}
	if strings.TrimSpace(req.PaymentMethod) == "" {
		req.PaymentMethod = "01"
	}
	if req.OrderDate == 0 {
		req.OrderDate = time.Now().UnixMilli()
	}

	req.MchId = cli.Params.MchId
	req.SignType = SignTypeSHA512
	if req.Version == "" {
		req.Version = Version10
	}
	req.NotifyUrl = cli.Params.PayinCallbackUrl
	req.ReturnUrl = cli.Params.ReturnUrl

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
		Post(cli.Config.PayinUrl)
	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("rivo payin failed: code=%s msg=%s", resp.Code, resp.Msg)
	}

	var data PayInResponseData
	if len(resp.Data) > 0 && string(resp.Data) != "null" {
		if err = json.Unmarshal(resp.Data, &data); err != nil {
			return nil, err
		}
	}

	return &data, nil
}

// QueryPayInOrder queries a pay-in order.
func (cli *Client) QueryPayInOrder(req PayInQueryRequest) (*PayInQueryResponseData, error) {
	if req.TradeNo == "" && req.OutTradeNo == "" {
		return nil, fmt.Errorf("tradeNo or outTradeNo must not be blank")
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
		Get(cli.Params.payinQueryURL())
	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("rivo payin query failed: code=%s msg=%s", resp.Code, resp.Msg)
	}

	var data PayInQueryResponseData
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
			return nil, fmt.Errorf("rivo payin query response signature verification failed")
		}
	}

	return &data, nil
}

// ParsePayInCallback parses and verifies a pay-in callback.
func (cli *Client) ParsePayInCallback(callback PayInCallback) (PayInCallback, error) {
	ok, err := verifyPayloadSign(callback, cli.Params.SecretKey, callback.Sign)
	if err != nil {
		return callback, err
	}
	if !ok {
		return callback, fmt.Errorf("rivo payin callback signature verification failed")
	}
	return callback, nil
}
