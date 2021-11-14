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
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}
	count := int64(-1)
	uList := []user{}
	u := user{}
	var err error

	// 数据库
	defaultDb := mysql.InitSQL("default")
	err = defaultDb.Master(ctx).Table("user").Where("`age` = ?", 12).Count(&count).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, count)
	err = defaultDb.Master(ctx).Table("user").Where("`age` = ?", 12).Find(&uList).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, uList)
	err = defaultDb.Master(ctx).Table("user").Where("`age` = ?", 12).First(&u).Error
	fmt.Printf("mysql 查询--------%#v %#v \n", err, u)
	// response.Success(ctx)
	// response.Success(ctx, uList)
	// response.Error(ctx, uconfig.ErrDemo)

	// redis
	defaultRedis := redis.InitRedis("default")
	r, e := defaultRedis.Set(ctx, "tt", "val", 60*time.Second).Result()
	fmt.Printf("--- %T, %#v\n", r, r)
	fmt.Printf("--- %T, %#v\n", e, e)

	// redis 集群
	defaultRedisCluster := rediscluster.InitRedisCluster("default")
	r, e = defaultRedisCluster.Set(ctx, "tt", "val", 60*time.Second).Result()
	fmt.Printf("=== %T, %#v\n", r, r)
	fmt.Printf("=== %T, %#v\n", e, e)

	response.Success(ctx)
	return

}

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
}
