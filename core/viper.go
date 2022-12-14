package core

import (
	"aixinge/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			// 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
				config = global.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", global.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//配置文件没有找到
			panic(fmt.Errorf("the config file does not exist: %s \n", err))
		} else {
			// 配置文件找到了,但是在这个过程有又出现别的什么error
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
