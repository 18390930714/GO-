package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// https://github.com/gomodule/redigo 包也是常用redis包
// go 语言连接redis: 普通连接
func main() {
	//doConnect()
	doDemo()
}

func doConnect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) //根据根context创建父context：ctx; withTimeout函数，设置超时时间
	defer cancel()                                                                 //延迟调用,直到 return 前才被执行,因此，可以用来做资源清理

	//执行命令获取结果
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val, err)
	fmt.Println("**************************")
	//先获取到命令对象
	cmder := rdb.Get(ctx, "key") //rdb.Get() -> redis的get key 命令
	fmt.Println(cmder.Val())     //获取值
	fmt.Println("**************************")
	fmt.Println(cmder.Err()) //获取错误
	fmt.Println("**************************")
	//直接执行命令获取错误
	err = rdb.Set(ctx, "key", 10, time.Hour).Err() //rdb.Set() -> redis的set命令，插入数据

	//直接执行命令获取值
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value)
	fmt.Println("**************************")
}

// Do 方法 可以执行任意命令或自定义命令的方法
func doDemo() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//直接执行命令获取错误
	err := rdb.Do(ctx, "set", "key1", 10, "EX", 3600).Err()
	fmt.Println(err)

	//执行命令获取结果
	val, err := rdb.Do(ctx, "get", "EX").Result()
	fmt.Println(val, err)
}

// go-redis 库提供了一个 redis.Nil 错误来表示 Key 不存在的错误。因此在使用 go-redis 时需要注意对返回错误的判断。在某些场景下我们应该区别处理 redis.Nil 和其他不为 nil 的错误。
func getValueFromRedis(key, defaultValue string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		//
		return "", err
	}
	return val, nil
}

func zsetDome() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	zsetKey := "language_rank"

	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "0Python"},
		{Score: 97.0, Member: "Java"},
		{Score: 99.0, Member: "C/C++"},
	}
	//ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("zincrby failed, err:%v\n", newScore)

	//取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//取95-100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}

	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}

	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// scan 遍历
func scanKeyDemo() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 {
			break
		}
	}
}

func scanKeysDemo2() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

// 匹配指定模式的key
func deleteByMatch(match string, timeout time.Duration) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.94:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter := rdb.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

/*
iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()
*/

//Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
