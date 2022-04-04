package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello/1" tags:"Hello" method:"get" summary:"You first hello api"`
	Name   string `p:"username"  v:"length:4,30#请输入账号|账号长度为:min到:max位"`
}
type HelloRes struct {
	ResponseFormat string
	Reply          string `dc:"Reply content"`
	Id             string `dc:"your user id"`
}
