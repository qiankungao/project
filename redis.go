package main

import (
	"fmt"
	"time"

	"github.com/1975210542/project/models"

	"github.com/gomodule/redigo/redis"
)

func main() {

	for i := 0; i < 4; i++ {
		go func() {
			c := models.Redis()

			_, setErr := c.Do("mset", "name", "gaoqiankun", "age", 25)
			models.CheckErr("setErr", setErr)
			if r, mgetErr := redis.Strings(c.Do("mget", "name", "age")); mgetErr == nil {
				for _, v := range r {
					fmt.Println("mget", v)
				}
			}
		}()
	}
	time.Sleep(1 * time.Second)

	//	p1.Description = "my blog"
	//	p1.Url = "http://xxbandy.github.io"
	//	p1.Author = "bgbiao"

	//	_, hmsetErr := c.Do("hmset", redis.Args{}.Add("hao123").AddFlat(&p1)...)
	//	errCheck("hmset", hmsetErr)

	//	m := map[string]string{
	//		"description": "oschina",
	//		"url":         "http://my.oschina.net/myblog",
	//		"author":      "xxbandy",
	//	}

	//	_, hmset1Err := c.Do("hmset", redis.Args{}.Add("hao").AddFlat(m)...)
	//	errCheck("hmset1", hmset1Err)

	//	for _, key := range []string{"hao123", "hao"} {
	//		v, err := redis.Values(c.Do("hgetall", key))
	//		errCheck("hmgetV", err)
	//		//等同于hgetall的输出类型，输出字符串为k/v类型
	//		//hashV,_ := redis.StringMap(c.Do("hgetall",key))
	//		//fmt.Println(hashV)
	//		//等同于hmget 的输出类型，输出字符串到一个字符串列表
	//		hashV2, _ := redis.Strings(c.Do("hmget", key, "description", "url", "author"))
	//		for _, hashv := range hashV2 {
	//			fmt.Println(hashv)
	//		}
	//		if err := redis.ScanStruct(v, &p2); err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		fmt.Printf("%+v\n", p2)

	//	}

}
