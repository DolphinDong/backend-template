package redis

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type RedisDao struct {
	*redis.Pool
}

func NewRedisDao() *RedisDao {
	return &RedisDao{
		global.RedisPool,
	}
}
func (rd *RedisDao) GetConnect() redis.Conn {
	return rd.Get()
}

func (rd *RedisDao) SetKeyExpiration(key string, expiration int) error {
	connect := rd.GetConnect()
	defer connect.Close()
	_, err := connect.Do("EXPIRE", key, expiration)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (rd *RedisDao) SetKeyWithExpiration(key, value string, expiration int) error {
	connect := rd.GetConnect()
	defer connect.Close()
	_, err := connect.Do("SETEX", key, expiration, value)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (rd *RedisDao) GetKey(key string) (interface{}, error) {
	connect := rd.GetConnect()
	defer connect.Close()
	value, err := connect.Do("GET", key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return value, nil
}

func (rd *RedisDao) DeleteKey(key string) error {
	connect := rd.GetConnect()
	defer connect.Close()
	_, err := connect.Do("DEL", key)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
