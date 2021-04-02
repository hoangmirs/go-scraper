package oauth

import (
	"context"
	"fmt"
	"time"

	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

var (
	oauthServer *server.Server
	clientStore *pg.ClientStore
	tokenStore  *pg.TokenStore
)

func SetUpOAuthServer() error {
	manager := manage.NewDefaultManager()

	pgxConn, err := pgx.Connect(context.TODO(), conf.GetString("dbUrl"))
	if err != nil {
		return err
	}

	adapter := pgx4adapter.NewConn(pgxConn)
	tStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		return err
	}
	defer tStore.Close()

	cStore, err := pg.NewClientStore(adapter)
	if err != nil {
		return err
	}

	manager.MapTokenStorage(tStore)
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

	oServer.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)

	oauthServer = oServer
	clientStore = cStore
	tokenStore = tStore

	return nil
}

func GetOAuthServer() *server.Server {
	return oauthServer
}

func GetClientStore() *pg.ClientStore {
	return clientStore
}

func GetTokenStore() *pg.TokenStore {
	return tokenStore
}

func passwordAuthorizationHandler(email string, password string) (string, error) {
	sessionForm := forms.SessionForm{
		Email:    email,
		Password: password,
	}

	user, err := sessionForm.Authenticate()
	if err != nil {
		return "", errors.ErrInvalidClient
	}

	return fmt.Sprint(user.Id), nil
}
