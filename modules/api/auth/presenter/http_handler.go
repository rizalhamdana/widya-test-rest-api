package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/api/auth/usecase"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	"github.com/sirupsen/logrus"
)

type HttpHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) *HttpHandler {
	return &HttpHandler{usecase: usecase}
}

func (h *HttpHandler) Mount(group *echo.Group) {
	group.POST("/login", h.LoginUser)
	group.POST("/register", h.RegisterUser)
}

func (h *HttpHandler) RegisterUser(e echo.Context) (err error) {
	ctx := "handler_auth.register_user"

	registerUser := model.User{}
	err = e.Bind(&registerUser)
	if err != nil {
		helper.Log(logrus.DebugLevel, err.Error(), ctx, "Binding request")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("failed to bind data request", nil))
	}

	successMessage, err := h.usecase.RegisterUser(&registerUser)
	if err != nil {
		helper.Log(logrus.DebugLevel, err.Error(), ctx, "Register user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(*successMessage, nil))

}

func (h *HttpHandler) LoginUser(e echo.Context) (err error) {
	ctx := "handler_auth.login_user"

	email, pass, ok := e.Request().BasicAuth()
	if !ok {
		err := "basic auth not found"
		helper.Log(logrus.ErrorLevel, err, ctx, "auth_login")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err, nil))
	}

	token, err := h.usecase.LoginUser(email, pass)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "auth_login")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data := make(map[string]string)
	data["token"] = *token
	return e.JSON(http.StatusOK, helper.ResponseDetail{
		Message: "success",
		Data:    data,
	})

}
