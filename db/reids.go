package db

import (
	"github.com/garyburd/redigo/redis"
)

type RedisInfo struct {
	Address  string
	PassWord string
	Zone     int
	Client   redis.Conn
}

func (redisInfo *RedisInfo) Init() (err error) {
	setDb := redis.DialDatabase(redisInfo.Zone)
	setPassword := redis.DialPassword(redisInfo.PassWord)
	rdb, err := redis.Dial("tcp", redisInfo.Address, setDb, setPassword)
	if err != nil {
		return
	}
	redisInfo.Client = rdb
	return
}
