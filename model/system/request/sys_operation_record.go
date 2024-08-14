package request

import (
	"gva/model/common/request"
	"gva/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
