package go_zpay

import (
	"fmt"
	"time"

	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
)

// CreatePayOutOrder 创建代付订单 (Pay-Out)
func (cli *Client) CreatePayOutOrder(req *PayOutRequest) (*PayOutResponseData, error) {
	req.MchId = cli.Config.MchId
	req.SignType = SIGN_TYPE_SHA512
	req.Version = VERSION_1_0
	req.NotifyUrl = cli.Config.PayoutCallbackUrl

	// 生成签名
	params := make(map[string]interface{})
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, err
	}
	req.Sign = utils.GenerateSign(params, cli.Config.SecretKey)

	var resp BaseResponse
	_, err := cli.ryClient.R().
		SetHeaders(getHeaders()).
		SetBody(req).
		SetResult(&resp).
		Post(cli.Config.PayoutUrl)

	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("api error: %s - %s", resp.Code, resp.Msg)
	}

	var data PayOutResponseData
	if err := mapstructure.Decode(resp.Data, &data); err != nil {
		return nil, err
	}

	// 校验响应签名
	dataParams := make(map[string]interface{})
	if err := mapstructure.Decode(data, &dataParams); err != nil {
		return nil, err
	}
	if !utils.VerifySign(dataParams, cli.Config.SecretKey, data.Sign) {
		return nil, fmt.Errorf("response signature verification failed")
	}

	return &data, nil
}

// QueryPayOutOrder 查询代付订单 (Pay-Out Query)
func (cli *Client) QueryPayOutOrder(req *PayOutQueryRequest) (*PayOutQueryResponseData, error) {
	req.MchId = cli.Config.MchId
	req.Version = VERSION_1_0
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().UnixNano() / 1e6
	}

	// 生成签名
	params := make(map[string]interface{})
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, err
	}
	req.Sign = utils.GenerateSign(params, cli.Config.SecretKey)

	// GET 请求参数
	queryParams := make(map[string]string)
	for k, v := range params {
		if v != nil {
			queryParams[k] = fmt.Sprintf("%v", v)
		}
	}
	queryParams["sign"] = req.Sign

	var resp BaseResponse
	_, err := cli.ryClient.R().
		SetQueryParams(queryParams).
		SetResult(&resp).
		Get(cli.Config.PayoutUrl)

	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("api error: %s - %s", resp.Code, resp.Msg)
	}

	var data PayOutQueryResponseData
	if err := mapstructure.Decode(resp.Data, &data); err != nil {
		return nil, err
	}

	// 校验响应签名
	dataParams := make(map[string]interface{})
	if err := mapstructure.Decode(data, &dataParams); err != nil {
		return nil, err
	}
	if !utils.VerifySign(dataParams, cli.Config.SecretKey, data.Sign) {
		return nil, fmt.Errorf("response signature verification failed")
	}

	return &data, nil
}

// ParsePayOutCallback 解析代付回调
func (cli *Client) ParsePayOutCallback(callback PayOutCallback) (*PayOutCallback, error) {
	// 校验签名
	params := make(map[string]interface{})
	if err := mapstructure.Decode(callback, &params); err != nil {
		return nil, err
	}
	if !utils.VerifySign(params, cli.Config.SecretKey, callback.Sign) {
		return nil, fmt.Errorf("wd callback signature verification failed")
	}

	return &callback, nil
}
