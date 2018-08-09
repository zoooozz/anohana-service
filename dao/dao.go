package dao

import (
	"anohana-service/config"
	"github.com/donnie4w/go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic"
	"golang/redigo/redis"
	"time"
)

type Dao struct {
	conf       *config.Config
	db         *sqlx.DB
	redis      *redis.Pool
	es         *elastic.Client
	expireTime int
}

func NewDao(c *config.Config) (d *Dao, err error) {
	d = &Dao{
		conf: c,
	}
	db, err := connectDB(c.Database.Master.Addr)
	d.db = db

	redis, err := connectRedis(c.Redis.Master.Addr, c.Redis.Master.Auth, c.Redis.Master.Db)
	d.redis = redis
	d.expireTime = 86400
	return

}

func connectDB(addr string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", addr)
	if err != nil {
		logger.Error(err)
	}
	if err := db.Ping(); err != nil {
		logger.Error(err)
	}
	return
}

func connectRedis(addr string, auth string, db int) (pool *redis.Pool, err error) {

	pool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   30,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", addr,
				redis.DialPassword(auth),
			)
			if err != nil {
				return nil, err
			}
			conn.Do("SELECT", db)
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return
}
