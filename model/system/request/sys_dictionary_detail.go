package request

import (
	"gva/model/common/request"
	"gva/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
