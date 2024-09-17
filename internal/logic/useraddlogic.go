package logic

import (
	"context"

	"mail/internal/model"
	"mail/internal/svc"
	"mail/internal/types"

	"github.com/rs/xid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddRequest) (resp *types.UserAddReply, err error) {
	// todo: add your logic here and delete this line
    l.svcCtx.UserModel.Insert(l.ctx,&model.Users{
		Id: xid.ID.String(),
	})
	return
}
