package controller

import (
	"context"

	v1 "gf2-start-up/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	res = &v1.HelloRes{
		ResponseFormat: "json",
		Reply:          "",
		Id:             "",
	}
	return
}
