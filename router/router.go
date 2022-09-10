package router

import (
	"fmt"
	"gowebdemo/core/logger"
	"gowebdemo/core/mysql"
	"gowebdemo/core/redis"
	"gowebdemo/core/rediscluster"
	"gowebdemo/core/response"
	"time"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
	logger.Log(ctx).Infof("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	logger.Log(ctx, "custom").Infof("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}
	count := int64(-1)
	uList := []user{}
	u := user{}
	var err error

	// 数据库
	defaultDbConn := mysql.GetSQLConn("default")
	err = defaultDbConn.Master(ctx).Table("user").Where("`age` = ?", 12).Count(&count).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, count)
	err = defaultDbConn.Master(ctx).Table("user").Where("`age` = ?", 12).Find(&uList).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, uList)
	err = defaultDbConn.Master(ctx).Table("user").Where("`age` = ?", 12).First(&u).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, u)
	// response.Success(ctx)
	// response.Success(ctx, uList)
	// response.Error(ctx, uconfig.ErrDemo)

	// redis
	defaultRedisConn := redis.GetRedisConn("default")
	r, e := defaultRedisConn.Set(ctx, "tt", "val", 60*time.Second).Result()
	fmt.Printf("--- %T, %#v\n", r, r)
	fmt.Printf("--- %T, %#v\n", e, e)

	// redis 集群
	defaultRedisClusterConn := rediscluster.GetRedisClusterConn("default")
	r, e = defaultRedisClusterConn.Set(ctx, "tt", "val", 60*time.Second).Result()
	fmt.Printf("=== %T, %#v\n", r, r)
	fmt.Printf("=== %T, %#v\n", e, e)

	response.Success(ctx)
	return

}

// Register 接口注册函数
func Register(r *gin.Engine) {
	r.GET("/ping", ping)
}
