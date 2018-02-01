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
	viper.SetDefault("api.listen", ":8080")
	viper.SetDefault("btc.testnet.zmq", "tcp://127.0.0.1:28332")
	viper.SetDefault("btc.testnet.rpc", "http://127.0.0.1:18332")
	// 加载配置
	err := viper.ReadInConfig()
	if err != nil {
		stdlog.Fatalln("load config error:", err)
	}

	// 接口文档
	if docuri := viper.GetString("api.doc"); docuri != "" {
		var doc = api.Router.Group(docuri)
		doc.GET("/", api.BinFsHandler("index.html"))
		doc.GET("/redoc.min.js", api.BinFsHandler("redoc.min.js"))
		doc.StaticFile("/swagger.json", "swagger.json")
	}

	api.Init()
}

func main() {
	var apisrv = &http.Server{
		Addr:    viper.GetString("api.listen"),
		Handler: api.Router,
	}
	stdlog.Fatalln(apisrv.ListenAndServe())
}
