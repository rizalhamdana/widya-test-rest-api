package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/share"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	log "github.com/sirupsen/logrus"
)

type AuthMiddlewareCustom struct {
	*share.Repository
}

func NewAuthMiddlewareCustom(repo *share.Repository) *AuthMiddlewareCustom {
	return &AuthMiddlewareCustom{repo}
}

func (m *AuthMiddlewareCustom) AuthorizationCustom(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		ctx := "authorization.middleware"

		token := e.Request().Header.Get("Authorization")
		if token == "" {
			err := "need request header 'Authorization'"
			helper.Log(log.ErrorLevel, err, ctx, "auth_login")
			return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err, nil))
		}
		arrToken := strings.Split(token, " ")
		token = arrToken[1]
		ID, err := m.getDataLogin(token)
		if err != nil {
			helper.Log(log.ErrorLevel, err.Error(), ctx, "auth_login")
			return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
		}
		e.Request().Header.Set("x-id-user", *ID)
		return next(e)
	}
}

func (m *AuthMiddlewareCustom) getDataLogin(keyToken string) (ID *string, err error) {
	ctx := "authorization.middleware.get_data_login"

	storedToken := model.StoredToken{}
	err = m.DBRead.Model(&storedToken).Where("token = ?", keyToken).Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			err = errors.New("Invalid Token")
			helper.Log(log.ErrorLevel, err.Error(), ctx, "auth_login")
			return nil, err
		}
		helper.Log(log.ErrorLevel, err.Error(), ctx, "auth_login")
		return nil, err
	}
	return &storedToken.IDUser, nil
}
