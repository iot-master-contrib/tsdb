package tsdb

import (
	"embed"
	"encoding/json"
	"github.com/iot-master-contrib/tsdb/api"
	_ "github.com/iot-master-contrib/tsdb/docs"
	"github.com/iot-master-contrib/tsdb/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

func App() *model.App {
	return &model.App{
		Id:   "tsdb",
		Name: "时序数据库",
		Icon: "/app/tsdb/assets/tsdb.svg",
		Entries: []model.AppEntry{{
			Path: "app/tsdb/history",
			Name: "历史",
		}, {
			Path: "app/tsdb/setting",
			Name: "配置",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	}
}

// go:embed all:app/tsdb
var wwwFiles embed.FS

// @title 历史数据库接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/tsdb/api/
// @query.collection.format multi
func main() {
}

func Startup(app *web.Engine) error {
	err := internal.Open()
	if err != nil {
		return err
	}

	internal.SubscribeProperty()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/tsdb/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/tsdb"), "tsdb")

	return nil
}

func Register() error {
	payload, _ := json.Marshal(App())
	token := mqtt.Publish("master/register", payload)
	token.Wait()
	return token.Error()
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/tsdb", http.FS(wwwFiles), "", "app/tsdb/index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
