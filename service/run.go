package service

import (
	"anohana-service/config"
	"anohana-service/dao"
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	svr *service
)

type service struct {
	dao *dao.Dao
}

func Run() (err error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	if svr, err = initService(); err != nil {
		panic(err)
	}
	initRouter(e)
	initLog()
	http := e.Start(":" + fmt.Sprintf("%d", config.Conf.Http.Port))
	e.Logger.Fatal(http)
	return
}

func initLog() {
	logger.SetConsole(true)
	logger.SetRollingDaily(config.Conf.Log.Addr, config.Conf.Log.Dir)
	logger.SetLevel(logger.INFO)
}

func initService() (s *service, err error) {
	s = &service{}
	if s.dao, err = dao.NewDao(config.Conf); err != nil {
		return
	}
	return
}
