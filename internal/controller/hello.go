package controller

import (
	"context"
	"gf-ruoyi/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
