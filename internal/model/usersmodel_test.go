package model

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rs/xid"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestNewUsersModel(t *testing.T) {
	conn:=sqlx.NewSqlConn("mysql","root:181205@tcp(192.168.3.55:3306)/test?charset=utf8mb4&parseTime=true")
	cacheConf:=cache.NodeConf{}
	cacheConf.Host="192.168.3.55:6379"
	cacheConf.Pass="181205"
	userModel:=NewUsersModel(conn,cacheConf)
	user:=Users{}
	user.Name="haha11"
	user.Mobile="18128282885"
	user.Avatar=sql.NullString{
		Valid: true,
		String: "https://www.baidu.com",
	}
	user.Id=xid.New().String()
	userModel.Insert(context.TODO(),&user)
}