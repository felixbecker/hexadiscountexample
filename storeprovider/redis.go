package storeprovider

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisProvider struct {
	redisPool *redis.Pool
}

func NewRedisProvider(addr string) *RedisProvider {
	rp := RedisProvider{}
	rp.redisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	conn := rp.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatal("error connecting redis")
	}

	_, err = redis.Float64(conn.Do("GET", "rate"))

	if errors.Is(err, redis.ErrNil) {
		fmt.Println("error is true -> Can not get a rate")
		conn.Do("SET", "rate", 1.24)
	}

	return &rp
}

func (rp *RedisProvider) Get(amount float32) float32 {

	conn := rp.redisPool.Get()
	defer conn.Close()

	v, err := redis.Float64(conn.Do("GET", "rate"))
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	return float32(v)
}

func (rp *RedisProvider) Set(amount float32, rate float32) {

	conn := rp.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", "rate", rate)
	if err != nil {
		panic(err)
	}
}
