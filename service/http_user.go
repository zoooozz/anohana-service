package service

import (
	"anohana-service/model"
	"encoding/json"
	"github.com/labstack/echo"
)

func user(c echo.Context) error {

	var (
		r   string
		err error
	)
	phone := c.Get("phone").(string)
	r, err = svr.dao.GetByUserRedis(phone)
	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	user := &model.AdminUModel{}
	_ = json.Unmarshal([]byte(r), user)

	result := map[string]interface{}{
		"list": user,
	}
	return c.JSON(model.OutputRet(model.Success, result))
}

func userEdit(c echo.Context) error {

	var (
		r   string
		err error
	)
	phone := c.Get("phone").(string)
	r, err = svr.dao.GetByUserRedis(phone)
	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	user := &model.AdminUModel{}
	_ = json.Unmarshal([]byte(r), user)

	result := map[string]interface{}{
		"list": user,
	}
	return c.JSON(model.OutputRet(model.Success, result))
}
