package oauth

import (
	"context"
	"time"

	"github.com/hoangmirs/go-scraper/conf"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore

func SetUpOAuthServer() error {
	manager := manage.NewDefaultManager()

	pgxConn, err := pgx.Connect(context.TODO(), conf.GetString("dbUrl"))
	if err != nil {
		return err
	}

	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		return err
	}
	defer tokenStore.Close()

	cStore, err := pg.NewClientStore(adapter)
	if err != nil {
		return err
	}

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(cStore)

	oServer := server.NewDefaultServer(manager)
	oServer.SetAllowGetAccessRequest(true)
	oServer.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	oServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logs.Error("Internal error: %v", err.Error())
		return
	})

	oServer.SetResponseErrorHandler(func(re *errors.Response) {
		logs.Error("Response error: %v", re.Error.Error())
	})

	oauthServer = oServer
	clientStore = cStore

	return nil
}

func GetOAuthServer() *server.Server {
	return oauthServer
}

func GetClientStore() *pg.ClientStore {
	return clientStore
}
