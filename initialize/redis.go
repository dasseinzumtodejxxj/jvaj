package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gva/global"
)

func Redis() {
	redisCfg := global.Gxva_CONFIG.Redis
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Gxva_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.Gxva_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.Gxva_REDIS = client
	}
}
