package handlers

import (
	"net/http"

	"tansan/user/models"
	"tansan/user/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type userHttpHandler struct {
 userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(userUsecase usecases.UserUsecase) UserHandler {
 return &userHttpHandler{
  userUsecase: userUsecase,
 }
}

func (h *userHttpHandler) DetectUser(c echo.Context) error {
 reqBody := new(models.AddUserData)

 if err := c.Bind(reqBody); err != nil {
  log.Errorf("Error binding request body: %v", err)
  return response(c, http.StatusBadRequest, "Bad request")
 }

 if err := h.userUsecase.UserDataProcessing(reqBody); err != nil {
  return response(c, http.StatusInternalServerError, "Processing data failed")
 }

 return response(c, http.StatusOK, "Success !!!")
}
