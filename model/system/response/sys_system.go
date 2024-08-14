package response

import "gva/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
