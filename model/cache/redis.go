package cache

import (
	"time"

	"github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

const (
	envKeyCacheRedisAddr         = "CACHE_REDIS_ADDR"
	envKeyCacheRedisPassword     = "CACHE_REDIS_PASSWORD"
	envKeyCacheRedisDB           = "CACHE_REDIS_DB"
	envKeyCacheRedisMaxRetries   = "CACHE_REDIS_MAX_RETRIES"
	envKeyCacheRedisDialTimeout  = "CACHE_REDIS_DIAL_TIMEOUT"
	envKeyCacheRedisReadTimeout  = "CACHE_REDIS_READ_TIMEOUT"
	envKeyCacheRedisWriteTimeout = "CACHE_REDIS_WRITE_TIMEOUT"
)

func parseEnvStringOrDefault(envKey, defaultStr string) string {
	var result string
	if viper.GetString(envKey) == "" {
		result = defaultStr
	} else {
		result = viper.GetString(envKey)
	}
	return result
}

func parseEnvIntOrDefault(envKey string, defaultInt int) int {
	var result int
	if viper.GetInt(envKey) == 0 {
		result = defaultInt
	} else {
		result = viper.GetInt(envKey)
	}
	return result
}

var DefaultRedis *redis.Client

// in viper yaml
//	Redis:
//	  addr: localhost:39006
//	  password:
//	  db: 0
//	  max_retries: 0 # Default is to not retry failed commands
//	  dial_timeout: 5 # Default is 5 seconds.
//	  read_timeout: 3 # Default is 3 seconds.
//	  write_timeout: 3 # Default is ReadTimeout
// and will cover by ENV
//	ENV_WEB_CACHE_REDIS_ADDR
//	ENV_WEB_CACHE_REDIS_PASSWORD
//	ENV_WEB_CACHE_REDIS_DB
//	ENV_WEB_CACHE_REDIS_MAX_RETRIES
//	ENV_WEB_CACHE_REDIS_DIAL_TIMEOUT
//	ENV_WEB_CACHE_REDIS_READ_TIMEOUT
//	ENV_WEB_CACHE_REDIS_WRITE_TIMEOUT
func InitRedis() (*redis.Client, error) {
	redisConf := config.Global().Redis

	redisAddr := parseEnvStringOrDefault(envKeyCacheRedisAddr, redisConf.Addr)
	redisPassword := parseEnvStringOrDefault(envKeyCacheRedisPassword, redisConf.Password)
	redisDB := parseEnvIntOrDefault(envKeyCacheRedisDB, redisConf.DB)
	redisMaxRetries := parseEnvIntOrDefault(envKeyCacheRedisMaxRetries, redisConf.MaxRetries)
	redisDialTimeout := parseEnvIntOrDefault(envKeyCacheRedisDialTimeout, redisConf.DialTimeout)
	redisReadTimeout := parseEnvIntOrDefault(envKeyCacheRedisReadTimeout, redisConf.ReadTimeout)
	redisWriteTimeout := parseEnvIntOrDefault(envKeyCacheRedisWriteTimeout, redisConf.WriteTimeout)

	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword, // no password set
		DB:           redisDB,       // use default DB
		MaxRetries:   redisMaxRetries,
		DialTimeout:  time.Duration(redisDialTimeout) * time.Second,
		ReadTimeout:  time.Duration(redisReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisWriteTimeout) * time.Second,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		zlog.S().Errorf("redis conn pong: %v , err: %v", pong, err)
		return nil, err
	} else {
		zlog.S().Infof("redis conn success pong: %v", pong)
		DefaultRedis = client
	}
	return client, nil
}

// scan keys instead redisCli.Keys()
//	redisCli *redis.Client
//	match string
//	maxCount int64
// return
//	error scan error
//	[]string removed repeated key
func RedisScanKeysMatch(redisCli *redis.Client, match string, maxCount int64) (error, []string) {
	var cursor uint64
	var scanFull []string
	for {
		keys, cursor, err := redisCli.Scan(cursor, match, maxCount).Result()
		if err != nil {
			return err, nil
		}
		if len(keys) > 0 {
			for _, v := range keys {
				scanFull = append(scanFull, v)
			}
		}
		if cursor == 0 {
			break
		}
	}
	scanRes := removeRepeatedElementString(scanFull)
	return nil, scanRes
}

func removeRepeatedElementString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
