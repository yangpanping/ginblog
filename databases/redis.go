package databases

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func RedisInit() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()
}
