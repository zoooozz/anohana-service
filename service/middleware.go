package service

import (
	"anohana-service/model"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Adminmiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				tokenStr string
				phone    string
				u        string
			)

			tokenStr = c.Request().Header.Get("token")
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(model.JwtSalt), nil
			})
			if err != nil {
				return c.JSON(model.OutputRet(model.RetLoginErr))
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				phone = claims["phone"].(string)
			} else {
				return c.JSON(model.OutputRet(model.RetLoginErr))
			}

			if u, err = svr.dao.GetByUserRedis(phone); err != nil {
				return c.JSON(model.OutputRet(model.RetLoginErr))
			}

			if u == "" {
				return c.JSON(model.OutputRet(model.RetLoginErr))
			}
			c.Set("phone", phone)
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			return next(c)
		}
	}
}
