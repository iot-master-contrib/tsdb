package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询数据库配置
// @Schemes
// @Description 查询数据库配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[influx.Options] 返回数据库配置
// @Router /config/tsdb [get]
func configGetInfluxdb(ctx *gin.Context) {
	curd.OK(ctx, nil)
}

// @Summary 修改数据库配置
// @Schemes
// @Description 修改数据库配置
// @Tags config
// @Param cfg body influx.Options true "数据库配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/tsdb [post]
func configSetInfluxdb(ctx *gin.Context) {
	curd.OK(ctx, nil)
}

func configRouter(app *gin.RouterGroup) {

	app.POST("/tsdb", configSetInfluxdb)
	app.GET("/tsdb", configGetInfluxdb)

}
