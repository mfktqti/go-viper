package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
viper 支持的配置很多
从JSON,TOML,YAML,HCL,envfile和java属性配置文件读取
实时监视和重新阅读配置文件
从环境变量中读取
从远程配置系统（etcd或consul）读取，并观察变化
从命令行标志读取
从缓冲区读取
*/
func main() {
	ReadIni()
	ReadToStruct()
}

type Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}

func ReadToStruct() {
	c := new(Config)
	v := viper.New()
	// path, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("path: %v\n", path)
	//v.AddConfigPath(path)
	// v.AddConfigPath("./conf") //路径//搞蒙了，有时能出来，有时出不来
	// v.SetConfigFile("config.yaml")
	// v.SetConfigType("yaml")

	v.SetConfigFile("./conf/config.yaml")

	err := v.ReadInConfig() //读取配置
	if err != nil {
		fmt.Println("读取配置文件出错：", err)
		return
	}

	if err := v.Unmarshal(c); err != nil {
		fmt.Println("读取配置出错：", err)
		return
	}

	fmt.Printf("c: %+v\n", c)
}

func ReadIni() {
	v := viper.New()
	v.SetConfigFile("./conf/config.ini")

	// v.AddConfigPath("./conf") //路径//搞蒙了，有时能出来，有时出不来
	// v.SetConfigName("config") //名称
	// v.SetConfigType("ini")    //类型

	err := v.ReadInConfig() //读取配置
	if err != nil {
		fmt.Println("读取配置文件出错：", err)
		return
	}
	//[section] 默认为default
	version := v.GetFloat64("default.version")
	fmt.Printf("version: %v\n", version)

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("文件被修改：")
		s := v.GetString("db.username")
		p := v.GetString("db.password")
		port := v.GetInt64("db.port")
		version := v.GetFloat64("default.version")
		fmt.Printf("version: %v\n", version)
		fmt.Printf("s: %v\n", s)
		fmt.Printf("p: %v\n", p)
		fmt.Printf("port: %v\n", port)
	})

	s := v.GetString("db.username")
	p := v.GetString("db.password")
	port := v.GetInt64("db.port")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("port: %v\n", port)

	time.Sleep(20 * time.Second)
}
