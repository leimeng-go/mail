package logic

import (
	"context"

	"mail/internal/svc"
	"mail/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrdersLogic {
	return &UserOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOrdersLogic) UserOrders(req *types.UserOrderRequest) (resp *types.UserOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
