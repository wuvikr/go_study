package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func initDB() {
	// 连接 Redis 服务器
	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.20.10.10:6379",
		Password: "",
		DB:       0,
		PoolSize: 20, // 连接池的大小
	})

	fmt.Printf("rdb: %v\n", rdb)
}

func doCommand() {
	ctx, cf := context.WithTimeout(context.Background(), 500*time.Microsecond)

	defer cf()

	// 执行get命令,先获取对象再取值
	sc := rdb.Get(ctx, "key")
	fmt.Printf("Val: %v\n", sc.Val())
	fmt.Printf("Err: %v\n", sc.Err())

	// 执行set命令
	val, err := rdb.Set(ctx, "key", 18, time.Hour).Result()
	fmt.Printf("val: %v, err: %v\n", val, err)

}

// go-redis do 方法
func do() {
	ctx, cnacel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cnacel()

	// 执行命令
	i, err := rdb.Do(ctx, "set", "age", 10, "EX", 3600).Result()
	fmt.Printf("i: %v\n", i)
	fmt.Printf("err: %v\n", err)

	i2, err2 := rdb.Do(ctx, "get", "age").Result()
	fmt.Printf("i2: %v\n", i2)
	fmt.Printf("err2: %v\n", err2)
}

func main() {
	initDB()

	// doCommand()

	do()
}
