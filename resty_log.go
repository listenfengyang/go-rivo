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

	restLog, err := json.Marshal(utils.GetRestyLog(resp))
	if err != nil {
		cli.logger.Errorf("PSPResty#rivo#%s log marshal err:%s", action, err.Error())
		return
	}
	cli.logger.Infof("PSPResty#rivo#%s->%s", action, string(restLog))
}
