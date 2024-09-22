package logic

import (
	"context"

	"mail/internal/svc"
	"mail/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserRequest) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
    mUser,err:=l.svcCtx.UserModel.FindOne(l.ctx,req.ID)
	if err!=nil{
		return  nil,err 
	}
	resp=&types.UserReply{
		ID: mUser.Id,
		Name: mUser.Name,
	}
	return
}
