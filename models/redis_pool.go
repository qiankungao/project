package models

import (
	"fmt"

	"os"
	"time"

	"github.com/1975210542/project/common"

	"github.com/gomodule/redigo/redis"
)

//构造连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:         5,
		MaxActive:       18,
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 300 * time.Second,
		Dial:            func() (redis.Conn, error) { return redisConn() },
	}
}

//构造连接
func redisConn() (redis.Conn, error) {

	ip := common.GetConfig("redis::ip")
	port := common.GetConfig("redis::port")
	passwd := common.GetConfig("redis::password")

	c, err := redis.Dial("tcp",
		ip+":"+port,
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
		redis.DialPassword(passwd),
		redis.DialKeepAlive(1*time.Second),
	)
	return c, err
}
func Redis() redis.Conn {
	pool := NewPool()
	defer pool.Close()

	return pool.Get()
}

//构造验错函数
func CheckErr(tp string, err error) {
	if err != nil {
		fmt.Println("sorry,has some error for %s.\r\n", tp, err)
		os.Exit(-1)
	}
}
