package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
