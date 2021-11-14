# gowebdome 是什么？
一个基于 Gin 的 Web Service 框架，封装了各种常用组件，包含数据库、配置、缓存、队列、TraceId 日志、crontab 命令等，让大伙可以以本项目为基础快速搭建Restful Web API。

这阵子学习 Go 语言之后，在搞 web 服务的时候，在网上找现成的 web 框架。但是每次 clone 下来之后，总发现有不如意的地方，可能是某个组件没有，可能是某些写法自己觉得太辣鸡，可能是自己就是想写个轮子......

总而言之，言而总之，干脆就自己搞一个吧，所以这玩意就诞生啦！

## 要干啥？

其实就是整合了许多 Web 开发 API 所必要的组件：

- [x] WEB 服务框架 Gin [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [x] MySQL 数据库 ORM 工具 GORM [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
- [x] 配置文件操作 viper [https://github.com/spf13/viper](https://github.com/spf13/viper)
- [x] Redis 和 RedisCluster go-redis [https://github.com/go-redis/redis](https://github.com/go-redis/redis)
- [ ] Kafka 队列 sarama [https://github.com/Shopify/sarama](https://github.com/Shopify/sarama)
- [x] log 日志 zap + lumberjack
  * zap [https://github.com/uber-go/zap](https://github.com/uber-go/zap)
  * lumberjack [https://github.com/natefinch/lumberjack](https://github.com/natefinch/lumberjack)
- [x] 优雅重启 endless [https://github.com/fvbock/endless](https://github.com/fvbock/endless)
- [ ] 命令行 cobra [https://github.com/spf13/cobra](https://github.com/spf13/cobra)


## 应用结构
```
[gowebdome]                 应用名称
├── conf                    配置文件
│   └── config.dev.yaml     开发环境配置文件
│   └── config.prod.yaml    生产环境配置文件
│   └── config.yaml         默认配置文件
├── controller              控制器
│   └── demo.go
├── core                    服务核心库
│   ├── config              配置解析和操作
│   ├── logger              日志模块
│   ├── middleware          中间件模块
│   ├── mysql               mysql配置
│   ├── response            接口返回体
├── data                    数据
│   └── log                 日志
├── library                 类库
├── middleware              中间件
├── model                   模型
├── router                  路由
│   └── router.go           统一定义路由文件
├── service                 业务逻辑处理
├── output                  打包后输出的部署文件目录
├── test                    单元测试
├── deploy.sh               打包部署脚本，输出在output目录下
├── go.mod                  go依赖
├── go.sum
├── LICENSE                 许可证
├── main.go                 主入口
└── README.md               说明
```

## 开发运行
本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```
cd gowebdome
go mod tidy
go run main.go
```

开发过程个人比较推荐使用 gowatch 进行热重启

gowatch [https://github.com/silenceper/gowatch](https://github.com/silenceper/gowatch)

```
cd gowebdome
go mod tidy
gowatch
```

## 编码规范
参考：
* Uber：[https://github.com/xxjwxc/uber_go_guide_cn](https://github.com/xxjwxc/uber_go_guide_cn)
* Go：[https://github.com/golang/go/wiki/CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

重点几个：
* 变量，函数：驼峰命名法
* 文件：全部小写，下划线分隔，尽量不要分隔
* 包名，目录名：全部小写，没有下划线


## TODO列表
- [x] 支持多数据库\多缓存
- [ ] 数据库自动读写分离
- [ ] 应用打包脚本
- [ ] dockerfile 文件编写
- [ ] 使用 docker-compose 把组件都编排上
- [ ] 服务发现
- [ ] 基于 OpenTracing 协议的 Jaeger 的分布式链路追踪
- [ ] 支持 gprc
