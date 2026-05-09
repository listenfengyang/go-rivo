package go_zpay

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-zpay/utils"
)

type Client struct {
	Config *RivoConfig

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, config *RivoConfig) *Client {
	return &Client{
		Config: config,

		ryClient:  resty.New(),
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugMode(debugMode bool) {
	cli.debugMode = debugMode
	if debugMode {
		cli.ryClient.SetDebug(true)
	}
}
