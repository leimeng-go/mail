package svc

import (
	"fmt"
	"mail/internal/config"
	"mail/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dns:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",c.Database.Host,c.Database.Port,c.Database.User,c.Database.Password,c.Database.Database,c.Database.Charset,c.Database.ParseTime,c.Database.Loc)
	conn:=sqlx.NewSqlConn(c.Database.Driver,dns)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUsersModel(conn,c.Cache),
	}
}
