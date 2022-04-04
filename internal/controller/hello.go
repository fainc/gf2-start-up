package controller

import (
	"context"

	v1 "gf2-start-up/api/v1"
	"gf2-start-up/internal/middleware"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	err = middleware.Response.NotAuthorized()
	return
}
