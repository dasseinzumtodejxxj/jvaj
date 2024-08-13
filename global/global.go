package global

import (
	"github.com/qiniu/qmgo"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"gva/config"
	"gva/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Gxva_DB     *gorm.DB
	Gxva_DBList map[string]*gorm.DB
	Gxva_REDIS  redis.UniversalClient
	Gxva_MONGO  *qmgo.QmgoClient
	Gxva_CONFIG config.Server
	Gxva_VP     *viper.Viper

	// GVA_LOG    *oplogging.Logger
	Gxva_LOG                 *zap.Logger
	Gxva_Timer               timer.Timer = timer.NewTimerTask()
	Gxva_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBname  通过名称获取db list中得db
func GetGlobalDBByDBname(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return Gxva_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := Gxva_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
