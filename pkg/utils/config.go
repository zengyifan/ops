package utils

import (
	"flag"
	mylog "github.com/rebirthmonkey/ops/pkg/log"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	// 定义一个命令行参数 -c 用于指定配置文件
	configPath := flag.String("c", "./configs/config.yaml", "config file path")
	flag.Parse()

	// 使用 Viper 读取配置文件
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	mylog.SetupLogger(viper.GetString("log.filePath"))
}
