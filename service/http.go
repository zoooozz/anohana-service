package service

import (
	"anohana-service/model"
	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	result := map[string]interface{}{
		"list": "非法请求",
	}
	return c.JSON(model.OutputRet(model.Success, result))
}
