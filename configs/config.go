package configs

import (
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/ini.v1"
)

type ServerConfig struct {
	Host string `ini:"host"`
	Port uint   `ini:"port"`
}

type UploadConfig struct {
	UploadRootPath string `ini:"upload_root_path"`
}

type MysqlConfig struct {
	Host        string `ini:"host"`
	Username    string `ini:"username"`
	Password    string `ini:"password"`
	Port        uint   `ini:"port"`
	Database    string `ini:"database"`
	Charset     string `ini:"charset"`
	TablePrefix string `ini:"table_prefix"`
}

type JwtConfig struct {
	Secret          string `ini:"secret"`
	RemeberDuration uint64 `ini:"remeber_duration"`
}

// 逐个加载，方便使用
var (
	Server = &ServerConfig{}
	Upload = &UploadConfig{}
	Mysql  = &MysqlConfig{}
	Jwt    = &JwtConfig{}
)

var Once sync.Once

func init() {
	Once.Do(func() {
		Load()
	})
}

func Load() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		panic("加载配置文件失败: " + err.Error())
	}

	if err = cfg.Section("server").MapTo(Server); err != nil {
		panic("加载配置文件失败: " + err.Error())
	}

	if err = cfg.Section("mysql").MapTo(Mysql); err != nil {
		panic("加载配置文件失败: " + err.Error())
	}

	if err = cfg.Section("upload").MapTo(Upload); err != nil {
		panic("加载配置文件失败: " + err.Error())
	}

	// 如果没有配置上传文件的根目录， 默认存储在程序运行目录下的 /uploads下面
	// todo, 检查是否有该目录的写入权限
	if Upload.UploadRootPath == "" {
		currentPath, err := os.Getwd()
		if err != nil {
			panic("获取运行目录失败：" + err.Error())
		}
		Upload.UploadRootPath = filepath.Join(currentPath, "/uploads")
	}

	if err = cfg.Section("jwt").MapTo(Jwt); err != nil {
		panic("加载配置文件失败: " + err.Error())
	}
}
