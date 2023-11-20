package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Urls 具体url信息。
type Urls struct {

	// urlid
	Id *int64 `json:"id,omitempty"`

	// url具体值。
	Url *string `json:"url,omitempty"`

	// url状态。
	Status *string `json:"status,omitempty"`

	// 任务类型。
	Type *string `json:"type,omitempty"`

	// 任务id。
	TaskId *int64 `json:"task_id,omitempty"`

	// 修改时间戳（毫秒）。
	ModifyTime *int64 `json:"modify_time,omitempty"`

	// 创建时间戳（毫秒）。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 文件类型，目录还是文件。
	FileType *string `json:"file_type,omitempty"`

	// 目录刷新方式，all：刷新目录下全部资源；detect_modify_refresh：刷新目录下已变更的资源，默认值为all。
	Mode *string `json:"mode,omitempty"`
}

func (o Urls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Urls struct{}"
	}

	return strings.Join([]string{"Urls", string(data)}, " ")
}
