package system

import (
	"gva/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
