package storage

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"fullSiteSpider/config"
)

var redisHandler *redisCache

func Redis() *redisCache {
	if redisHandler == nil {
		redisHandler = newRedis()
	}
	return redisHandler
}

func newRedis() *redisCache {

	redisHandler = new(redisCache)
	redisHandler.mutex = new(sync.RWMutex)
	redisHandler.conn()
	return redisHandler
}

// redisHandler is Redis redisHandler adapter.
type redisCache struct {
	pool     *redis.Pool // redis connection pool
	dbNum    int
	password string
	mutex    *sync.RWMutex
}

func (db *redisCache) Keys(pattern string) ([]string, error) {
	result, err := db.do("KEYS", pattern)
	return redis.Strings(result, err)
}

func (db *redisCache) Set(key string, data string) (err error) {
	_, err = db.do("SET", key, data)
	return err
}

func (db *redisCache) Expire(key string, timeoutSecond int) (err error) {
	_, err = db.do("EXPIRE", key, timeoutSecond)
	return err
}

func (db *redisCache) TTL(key string) (int, error) {
	result, err := db.do("TTL", key)
	return redis.Int(result, err)
}

func (db *redisCache) Get(key string) (string, error) {
	result, err := db.do("GET", key)
	return redis.String(result, err)
}

func (db *redisCache) ZAdd(key string, score int, value string) (err error) {
	_, err = db.do("ZADD", key, score, value)
	return err
}

func (db *redisCache) ZScore(key string, value string) (string, error) {
	result, err := db.do("ZSCORE", key, value)
	return redis.String(result, err)
}

func (db *redisCache) ZExist(key, field string) bool {
	_, err := db.ZScore(key, field)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *redisCache) ZRem(key string, value string) (err error) {
	_, err = db.do("ZREM", key, value)
	return err
}

func (db *redisCache) ZRangeByScore(key string, scoreMin string, scoreMax string) (values []string, err error) {

	return redis.Strings(db.do("ZRANGEBYSCORE", key, scoreMin, scoreMax))
}

func (db *redisCache) ZRangeWithScoreByScore(key string, scoreMin string, scoreMax string) (values map[string]string, err error) {

	return redis.StringMap(db.do("ZRANGEBYSCORE", key, scoreMin, scoreMax, "WITHSCORES"))
}

func (db *redisCache) RPOPLPUSH(fromList, toList string) (string, error) {

	result, err := db.do("RPOPLPUSH", fromList, toList)
	return redis.String(result, err)
}

func (db *redisCache) LPop(key string) (string, error) {

	result, err := db.do("LPOP", key)
	return redis.String(result, err)
}

func (db *redisCache) RPop(key string) (string, error) {

	result, err := db.do("RPOP", key)
	return redis.String(result, err)
}

func (db *redisCache) LPush(key string, value string) (err error) {

	_, err = db.do("LPUSH", key, value)
	return err
}

func (db *redisCache) RPush(key string, value string) (err error) {

	_, err = db.do("RPUSH", key, value)
	return err
}

func (db *redisCache) HSet(key, field, value string) (err error) {
	_, err = db.do("HSET", key, field, value)
	return err
}

func (db *redisCache) HMSet(key string, keyVal []string) (err error) {

	args := []interface{}{}
	args = append(args, key)

	for _, value := range keyVal {
		args = append(args, value)
	}

	_, err = db.do("HMSET", args...)
	return err
}

func (db *redisCache) HGet(key, field string) (string, error) {
	result, err := db.do("HGET", key, field)
	return redis.String(result, err)
}

func (db *redisCache) HDel(key, field string) (error) {
	_, err := db.do("HDEL", key, field)
	return err
}

func (db *redisCache) SAdd(key string, value string) (err error) {
	_, err = db.do("SAdd", key, value)
	return err
}

func (db *redisCache) SIsmember(key string, value string) (r interface{}, err error) {
	r, err = db.do("SISMEMBER", key, value)
	return r, err
}

func (db *redisCache) SExist(key string, value string) (bool) {

	result, err := redis.Int64(db.SIsmember(key, value))
	if err == nil && result == 1 {
		return true
	}
	return false
}

func (db *redisCache) do(cmd string, args ...interface{}) (interface{}, error) {
	db.mutex.Lock()
	client := db.pool.Get()

	defer func() {

		client.Close()

		db.mutex.Unlock()
	}()

	return client.Do(cmd, args...)
}

func (r *redisCache) conn() (*redisCache) {

	redisConfig := config.Redis()

	pool := &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: redisConfig.ConnectionMax(),

		Dial: func() (redis.Conn, error) {

			connect, err := redis.Dial("tcp", redisConfig.Link())
			if err != nil {
				//return nil, err
				panic(err.Error())
			}

			if redisConfig.Password != "" {
				if _, err := connect.Do("AUTH", redisConfig.Password); err != nil {
					connect.Close()
					return nil, err
				}
			}

			if redisConfig.DBNumber() != 0 {
				connect.Do("SELECT", redisConfig.DBNumber())
			}

			return connect, err
		},
	}
	r.pool = pool
	return r
}
