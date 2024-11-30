package repo

import (
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/orm"
	"e-commerce/service/infra/redis"
	"go.uber.org/zap"
)

func Init() {
	// init db connections
	initDBConnection(consts.DB_RETAIL)

	//initRedisConnection()
}

func initDBConnection(name string) {
	log.Info("Connect to Mysql database", zap.String("name", name))
	// 通用参数
	addr := "127.0.0.1:3306"
	username := "root"
	password := ""
	maxOpenConns := 10
	maxIdleConns := 1

	// TODO 入参 与 此处hardcode 逻辑待优化
	dbName := "retail"

	options := []orm.Option{
		orm.WithDBname(dbName),
		orm.WithAddr(addr),
		orm.WithUsername(username),
		orm.WithPassword(password),
		orm.WithMaxOpenConns(maxOpenConns),
		orm.WithIdleConns(maxIdleConns),
	}

	orm.NewDB(name, options...)
}

func initRedisConnection() {
	host := "127.0.0.1"
	port := "6379"
	password := ""
	dbnum := 6

	log.Infof("Initialize cache, host: %s, port: %s, password: %s, dbnum: %d ", host, port, password, dbnum)
	redis.NewClient(consts.RETAIL_STOCK_LOCK, host, port, password, dbnum)
}

// defer close connection
func Close() {
	orm.GetORM(consts.DB_RETAIL).Close()

	redis.GetClient(consts.REDIS_RETAIL).Close()
	redis.GetClient(consts.RETAIL_STOCK_LOCK).Close()
}

func IsRecordNotFoundError(err error) bool {
	return orm.IsRecordNotFoundError(err)
}
