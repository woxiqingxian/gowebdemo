package router

import (
	"fmt"
	"gowebdemo/core/logger"
	"gowebdemo/core/mysql"
	"gowebdemo/core/response"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
	logger.Log(ctx).Infof("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	type user struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}
	count := int64(-1)
	uList := []user{}

	defaultDb := mysql.InitSQL("default")
	err := defaultDb.Master(ctx).Table("user").Where("`age` = ?", 12).Find(&uList).Error
	fmt.Printf("mysql 查询--------%#v %#v %#v \n", count, err, uList)

	// ctx.JSON(200, map[string]string{"aa": "111", "bbb": "2222"})
	// ctx.Abort()
	// response.Success(ctx)
	response.Success(ctx, uList)
	// response.Error(ctx, uconfig.ErrParam)
	return

}

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
}
