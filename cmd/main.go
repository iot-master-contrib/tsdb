package main

import (
	"github.com/iot-master-contrib/tsdb"
	"github.com/iot-master-contrib/tsdb/config"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func main() {
	banner.Print("iot-master-plugin:tsdb")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}

	app := web.CreateEngine()

	//调用启动
	err = tsdb.Startup(app)
	if err != nil {
		log.Fatal(err)
	}

	err = tsdb.Register()
	if err != nil {
		log.Fatal(err)
	}

	//注册静态页面
	fs := app.FileSystem()
	tsdb.Static(fs)

	//监听HTTP
	app.Serve()
}
