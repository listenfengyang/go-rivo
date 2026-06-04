package rivo

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-rivo/utils"
)

func (cli *Client) logResty(action string, resp *resty.Response) {
	if cli == nil || cli.logger == nil || resp == nil {
		return
	}

	prettyLog, err := json.MarshalIndent(utils.GetRestyLog(resp), "", "  ")
	if err != nil {
		cli.logger.Infof("PSPResty#rivo#%s marshal log failed: %v", action, err)
		return
	}
	cli.logger.Infof("PSPResty#rivo#%s:\n%s", action, prettyLog)
}
