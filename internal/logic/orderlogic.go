package logic

import (
	"context"

	"mail/internal/svc"
	"mail/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.OrderRequest) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *OrderLogic)ordersByUser(uid int64) ([]*types.OrderReply,error) {
	if uid==123{
		return []*types.OrderReply{
			{
				ID: "32fj2938jrf93",
				State: 1,
				CreateAt: "2022-08-01",
		},{
			ID: "vjiejfoef323",
			State: 1,
			CreateAt: "2022-08-02",
		}},nil 
    }
	return nil,nil 
}