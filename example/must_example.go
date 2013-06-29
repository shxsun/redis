package main

import (
	"github.com/shxsun/beelog"
	"github.com/shxsun/redis"
	"strconv"
)

var client redis.Client

func main() {
	var err error
	defer func() {
		if e := recover(); e != nil {
			beelog.Critical(e)
		}
	}()

	err = client.Set("what", []byte(strconv.Itoa(123)))
	if err != nil {
		beelog.Error(err)
		return
	}

	n := redis.MustInt(client.Get("what"))
	beelog.Info("get int from redis:", n)

	client.Del("world")

	client.Lpush("world", []byte("blue"))
	client.Lpush("world", []byte("red"))
	beelog.Info("get list from redis:", redis.MustStrings(client.Lrange("wh", 0, -1)))
}
