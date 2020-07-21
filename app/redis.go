package app

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gookit/ini/v2"

	"time"
)

var (
	Names   []string
	nameMap = make(map[string]bool, 0)
	pools   map[string]*redis.Pool
)

func loadServerNames() {
	Names = ini.Strings("servers", ",")
	for _, name := range Names {
		nameMap[name] = true
	}
}

// HasName
func HasName(name string) bool {
	_, ok := nameMap[name]
	return ok
}

// Rds get redis connection
func Rds(name string) *redis.Pool {
	if pl, ok := pools[name]; ok {
		return pl
	}

	if !HasName(name) {
		return nil
	}

	if conf := ini.StringMap(name); len(conf) > 0 {
		pl := newPool(conf["url"], conf["pwd"], 0)
		pools[name] = pl

		return pl
	}

	return nil
}

// create new pool
func newPool(url, password string, redisDb int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}

			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", redisDb)

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
