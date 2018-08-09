package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Name     string
	Database *Database
	Http     *Http
	Log      *Log
	Redis    *Redis
	Qiniu    *Qiniu
}

type Http struct {
	Port int
}

type Mysql struct {
	Addr   string
	Active int
	Idle   int
}

type Database struct {
	Master *Mysql
}

type RedisConn struct {
	Addr string
	Auth string
	Db   int
}

type Redis struct {
	Master *RedisConn
}

type Log struct {
	Addr string
	Dir  string
}
type Qiniu struct {
	Access string
	Secret string
	Bucket string
}

var (
	Conf     = &Config{}
	ConfPath string
)

func init() {
	flag.StringVar(&ConfPath, "conf", "dev.toml", "config path")
}

func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}
