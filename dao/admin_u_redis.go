package dao

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"golang/redigo/redis"
)

const (
	_tokenKey = "anohana_login_token:%s"
)

func (d *Dao) SetByUserRedis(phone, r string) (err error) {
	conn := d.redis.Get()
	defer conn.Close()
	key := fmt.Sprintf(_tokenKey, phone)
	if err = conn.Send("SET", key, r); err != nil {
		return
	}
	if err = conn.Send("EXPIRE", key, d.expireTime); err != nil {
		return
	}
	if err = conn.Flush(); err != nil {
		return
	}

	for i := 0; i < 2; i++ {
		if _, err = conn.Receive(); err != nil {
			logger.Error(err)
			return
		}
	}
	return nil
}

func (d *Dao) GetByUserRedis(phone string) (r string, err error) {
	conn := d.redis.Get()
	defer conn.Close()
	key := fmt.Sprintf(_tokenKey, phone)
	if r, err = redis.String(conn.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		logger.Info(err)
	}
	return
}
