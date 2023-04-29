package cache

import (
	"blogProject/pkg/config"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"time"
)
const (
	LoginUserKey             = "blog:login:user:" // 已登陆用户在 redis 中存放的 key, 该 key + username 构成完整的 key
	LoginUserExpiredDuration = time.Hour * 24    // 用户登录信息在缓存中存在的时长
)


var rdb *redis.Client

// New函数用于创建Redis客户端的新实例，并返回一个错误对象。

func New() error{
	cfg := config.Get().RedisConfig
	rdb = redis.NewClient(&redis.Options{
		Addr:               cfg.Addr,
		Password:           cfg.Password,
		DB:                 cfg.Database,
	})
	// 尝试向Redis服务器发送一个PING命令，检查连接是否正常。
	// 如果PING命令返回错误，则将其作为错误返回。
	_,err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}


//
// Set 将键值对存入 redis 中
//	@parameter	key 键
//	@parameter	value 值
//	@parameter	ttl 过期时间,这里是一个时间段,也就是这个键值对儿在 redis 中的存活时长
//	@return		error 错误信息
//
func Set(key, value string, ttl time.Duration) error {
	if err := rdb.Set(context.Background(), key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}
	return nil
}

//
// Get 从 redis 中获取 key 对应的 value
//	@parameter	key 键
//	@return		string key 在 redis 中对应的 value
//	@return		error 错误信息
//
func Get(key string) (string, error) {
	value, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}
	return value, nil
}

