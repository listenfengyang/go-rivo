package rivo

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	// Params is kept for compatibility with payment-service-v2 internal usages.
	Params *RivoInitParams
	// Config is an alias-style field for external callers.
	Config *RivoConfig

	ryClient  *resty.Client
	debugMode bool
	logger    Logger
}

func NewClient(logger Logger, config *RivoInitParams) *Client {
	if config == nil {
		config = &RivoInitParams{}
	}
	config.normalize()

	return &Client{
		Params: config,
		Config: config,

		ryClient:  resty.New(),
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugMode(debugMode bool) {
	cli.debugMode = debugMode
	cli.ryClient.SetDebug(debugMode)
}

// SetDebugModel keeps compatibility with existing SDK call sites.
func (cli *Client) SetDebugModel(debugMode bool) {
	cli.SetDebugMode(debugMode)
}
