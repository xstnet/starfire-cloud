package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/xstnet/starfire-cloud/configs"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() { initDb() })
}

func initDb() {

	logConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             0,           // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 是否启用彩色打印
		},
	)

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: configs.Mysql.TablePrefix, // 表名前缀
			// SingularTable: true,  // 使用单数表名
		},
		Logger: logConfig,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		configs.Mysql.Username,
		configs.Mysql.Password,
		configs.Mysql.Host,
		configs.Mysql.Port,
		configs.Mysql.Database,
		configs.Mysql.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	DB = db
}
