package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/tsdb/internal"
	"github.com/xuri/excelize/v2"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询历史数据
// @Schemes
// @Description 查询历史数据
// @Tags query
// @Param pid path string true "产品ID"
// @Param id path string true "设备ID"
// @Param name path string true "变量名称"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Param window query string false "窗口时间"
// @Param fn query string false "算法"
// @Produce json
// @Success 200 {object} ReplyData[[]influx.Point] 返回报警信息
// @Router /query/{pid}/{id}/{name} [get]
func noopQuery() {}

// @Summary 导出历史数据
// @Schemes
// @Description 导出历史数据
// @Tags query
// @Param pid path string true "产品ID"
// @Param id path string true "设备ID"
// @Param name path string true "变量名称"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Param window query string false "窗口时间"
// @Param fn query string false "算法"
// @Produce json
// @Success 200 {object} ReplyData[[]influx.Point] 返回报警信息
// @Router /query/{pid}/{id}/{name}/export [get]
func noopQueryExport() {}

func queryRouter(app *gin.RouterGroup) {

	app.GET("/:pid/:id/:name", func(ctx *gin.Context) {
		pid := ctx.Param("pid")
		id := ctx.Param("id")
		key := ctx.Param("name")

		start := ctx.DefaultQuery("start", "-5h")
		end := ctx.DefaultQuery("end", "0h")
		window := ctx.DefaultQuery("window", "10m")
		//fn := ctx.DefaultQuery("fn", "mean") //last

		//values, err := influx.Query(pid, id, key, start, end, window, fn)
		values, err := internal.Query(pid, id, key, start, end, window)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		curd.OK(ctx, values)
	})

	app.GET("/:pid/:id/:name/export", func(ctx *gin.Context) {
		pid := ctx.Param("pid")
		id := ctx.Param("id")
		key := ctx.Param("name")

		start := ctx.DefaultQuery("start", "-5h")
		end := ctx.DefaultQuery("end", "0h")
		window := ctx.DefaultQuery("window", "10m")
		//fn := ctx.DefaultQuery("fn", "mean") //last

		//values, err := influx.Query(pid, id, key, start, end, window, fn)
		values, err := internal.Query(pid, id, key, start, end, window)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		if len(values) == 0 {
			curd.Fail(ctx, "无记录")
			return
		}

		//创建文件
		excel := excelize.NewFile()
		defer excel.Close()

		index, err := excel.NewSheet(key)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Set value of a cell.
		_ = excel.SetCellValue(key, "A1", "time")
		_ = excel.SetCellValue(key, "B1", "value")

		for k, v := range values {
			_ = excel.SetCellValue(key, fmt.Sprintf("A%d", k+1), v.Time)
			_ = excel.SetCellValue(key, fmt.Sprintf("B%d", k+1), v.Value)
		}

		// Set active sheet of the workbook.
		excel.SetActiveSheet(index)

		filename := pid + "-" + key + "-" + values[0].Time.Format("20060102150405") + "-" + values[len(values)-1].Time.Format("20060102150405") + ".xlsx"

		//下载头
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+filename+".xlsx") // 用来指定下载下来的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")

		_ = excel.Write(ctx.Writer)
	})

}
