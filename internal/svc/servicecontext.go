package svc

import (
	"mail/internal/config"
	"mail/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewSqlConn("mysql","root:181205@tcp(IP_ADDRESS:3306)/mail?charset=utf8mb4&parseTime=true")
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUsersModel(conn,c.Cache),
	}
}
