package presenter

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/api/profile/usecase"
	"github.com/sirupsen/logrus"
)

type HttpHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(usecase usecase.ProfileUsecase) *HttpHandler {
	return &HttpHandler{usecase: usecase}
}

func (h *HttpHandler) Mount(group *echo.Group) {
	group.GET("/profile", h.GetProfile)
}

func (h *HttpHandler) GetProfile(e echo.Context) (err error) {
	ctx := "handler_profile.get_profile"

	IDUser := e.Request().Header.Get("x-id-user")
	if len(IDUser) <= 0 {
		err = errors.New("Please specify Token")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Get Profile")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	profileResult, err := h.usecase.GetProfile(IDUser)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Get Profile")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Berhasil mendapatkan data", profileResult))

}
