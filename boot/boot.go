package boot

import (
	"io"
	"os"
)

func init() {
	copyIniConf()
}

func copyIniConf() {
	iniConfig := "configs/config.ini"
	iniExampleConfig := "configs/config.example.ini"

	fileInfo, err := os.Stat(iniConfig)

	if err == nil {
		if !fileInfo.IsDir() {
			return
		}
		panic("配置文件目录存在同名的文件夹，无法创建配置文件")
	}

	// 打开文件失败，并且返回的错误不是文件未找到
	if !os.IsNotExist(err) {
		panic("初始化失败: " + err.Error())
	}

	// 自动复制一份config.ini
	source, err := os.Open(iniExampleConfig)
	if err != nil {
		panic("创建配置文件失败，config.example.ini文件不存在: " + err.Error())
	}
	defer source.Close()

	dst, err := os.Create(iniConfig)
	if err != nil {
		panic("生成config.ini失败: " + err.Error())
	}
	defer dst.Close()

	_, err = io.Copy(dst, source)
	if err != nil {
		panic("写入config.ini失败: " + err.Error())
	}
}
