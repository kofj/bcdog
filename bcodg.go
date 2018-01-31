package main

import (
	stdlog "log"
	"net/http"
	"os"

	"github.com/kofj/bcdog/api"
	"github.com/spf13/viper"
)

func init() {
	// 设置配置查找路径与配置名称
	if len(os.Args) == 2 {
		viper.AddConfigPath(os.Args[1])
	} else {
		viper.AddConfigPath(".")
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/bcdog")
	viper.AddConfigPath("$HOME/.bcdog")
	// 设置配置缺省值
	viper.SetDefault("APIs.Listen", ":8080")
	// 加载配置
	err := viper.ReadInConfig()
	if err != nil {
		stdlog.Fatalln("load config error:", err)
	}

	api.Init()
}
func main() {
	var apisrv = &http.Server{
		Addr:    viper.GetString("APIs.Listen"),
		Handler: api.Engine,
	}
	stdlog.Fatalln(apisrv.ListenAndServe())
}
