package service

import (
	"anohana-service/common"
	"anohana-service/model"
	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	var (
		err      error
		resp     *model.AdminUModel
		token    string
		phone    string
		password string
	)
	phone = c.FormValue("phone")
	if phone == "" || !common.PhoneLegitimate(phone) {
		return c.JSON(model.OutputRet(model.ParamsPhoneErr))
	}

	password = c.FormValue("password")
	if password == "" {
		return c.JSON(model.OutputRet(model.ParamsPassErr))
	}

	if resp, token, err = svr.PassLogin(phone, password); err != nil {
		return c.JSON(model.OutputRet(model.ParamsPassErr))
	}

	if token == "" || resp == nil {
		return c.JSON(model.OutputRet(model.ParamsPassErr))
	}
	if resp.State == model.AdminUstateDisable {
		return c.JSON(model.OutputRet(model.RetLoginInfoErr))
	}

	result := map[string]interface{}{
		"token": token,
	}
	return c.JSON(model.OutputRet(model.Success, result))
}
