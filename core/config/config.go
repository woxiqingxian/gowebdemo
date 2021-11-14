package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetUp() {

	flag.StringVar(&ServerConfig.AppConfigPath, "config", "conf/config.yaml", "Configuration file")
	flag.Parse()

	var err error
	if ServerConfig.AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(fmt.Sprintf("Get App Path Error:%s", err))
	}

	viper.SetConfigFile(ServerConfig.AppConfigPath)
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Parse Config Get File Error: %s", err))
	}
	if err = viper.Unmarshal(&ServerConfig); err != nil {
		panic(fmt.Sprintf("Parse Config Unmarshal Error: %s", err))
	}

	return
}
