package main
import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// go get -u github.com/go-redis/redis

var redisdb *redis.Client

func init() {
	//redis-cli -p 7000 -h 10.20.52.203 -c
	/* redis return *Client */
	/* redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     "10.20.52.203:9001",
		Password: "",
		DB:       0,
	}) */

	//cluster redis return cluster client
	redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	cmd := redisdb.Ping()
	if cmd.Err() != nil {
		fmt.Println("init error", cmd.Err())
		return
	}
	fmt.Println("init success")
}

func main() {
	var key string = "asd"
	//時間設為0 = 永久
	cmd2 := redisdb.Set(key, "123", time.Second*1000)
	fmt.Println(cmd2)

	cmd := redisdb.Get(key)
	fmt.Println(cmd)

	if cmd.Err() != nil {
		fmt.Println("get error!!", cmd.Err())
	}
	fmt.Println(cmd.Val())
	redisdb.Del(key)
}
