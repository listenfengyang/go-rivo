package go_zpay

import (
	"fmt"
	"time"

	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
)

// CreatePayInOrder 创建收款订单 (Pay-In)
func (cli *Client) CreatePayInOrder(req *PayInRequest) (*PayInResponseData, error) {
	req.MchId = cli.Config.MchId
	req.SignType = SIGN_TYPE_SHA512
	req.Version = VERSION_1_0
	req.NotifyUrl = cli.Config.PayinCallbackUrl
	req.ReturnUrl = cli.Config.ReturnUrl

	// 生成签名
	params := make(map[string]interface{})
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, err
	}
	fmt.Printf("params: %v\n req: %v\n", params, req)

	req.Sign = utils.GenerateSign(params, cli.Config.SecretKey)

	var resp BaseResponse
	_, err := cli.ryClient.R().
		SetHeaders(getHeaders()).
		SetBody(req).
		SetResult(&resp).
		Post(cli.Config.PayinUrl)

	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("api error: %s - %s", resp.Code, resp.Msg)
	}

	var data PayInResponseData
	if err := mapstructure.Decode(resp.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// QueryPayInOrder 查询收款订单 (Pay-In Query)
func (cli *Client) QueryPayInOrder(req *PayInQueryRequest) (*PayInQueryResponseData, error) {
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
		Get(cli.Config.PayinUrl)

	if err != nil {
		return nil, err
	}

	if resp.Code != "200" {
		return nil, fmt.Errorf("api error: %s - %s", resp.Code, resp.Msg)
	}

	var data PayInQueryResponseData
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

// ParsePayInCallback 解析收款回调
func (cli *Client) ParsePayInCallback(callback PayInCallback) (PayInCallback, error) {
	// 校验签名
	params := make(map[string]interface{})
	if err := mapstructure.Decode(callback, &params); err != nil {
		return callback, err
	}

	if !utils.VerifySign(params, cli.Config.SecretKey, callback.Sign) {
		return callback, fmt.Errorf("dp callback signature verification failed")
	}

	return callback, nil
}
