package config

var ServerConfig ServerConf

type ServerConf struct {
	AppPath               string             `mapstructure:"app_path"`
	AppConfigPath         string             `mapstructure:"app_config_path"`
	AppConfig             AppConf            `mapstructure:"app_config"`
	LogConfig             LogConf            `mapstructure:"log_config"`
	RedisConfigList       []RedisConf        `mapstructure:"redis_config_list"`
	RedisClusteConfigList []RedisClusterConf `mapstructure:"redis_cluster_config_list"`
	MysqlConfigList       []MysqlConf        `mapstructure:"mysql_config_list"`
}

type AppConf struct {
	AppName   string `mapstructure:"app_name"`   // 应用名称
	Debug     bool   `mapstructure:"debug"`      // 是否调试模式
	AccessLog bool   `mapstructure:"access_log"` // 是否记录访问日志
	HttpAddr  string `mapstructure:"http_addr"`  // 监听地址
	HttpPort  int    `mapstructure:"http_port"`  // 监听端口号
}

type LogConf struct {
	LogNameList []string `mapstructure:"log_name_list"` // 自定义日志文件
	LogDir      string   `mapstructure:"log_dir"`       // 日志路径
	Level       string   `mapstructure:"level"`         // 日志等级,支持 debug/info/warn/error/dpanic/panic/fatal 共7种日志级别,级别从左往右为从小到大
	MaxSize     int      `mapstructure:"max_size"`      // 日志轮转配置, 单位MB,表示最大文件大,超出则会新生成一个日志文件,默认为100MB
	MaxAge      int      `mapstructure:"max_age"`       // 日志轮转配置, 文件最多保存多少天,单位天,默认不移除
	MaxBackups  int      `mapstructure:"max_backups"`   // 日志轮转配置, 日志文件最多保存多少个备份,默认保留所有日志文件
	Compress    bool     `mapstructure:"compress"`      // 日志轮转配置, 决定是否压缩日志文件存放
}

type RedisConf struct {
	Name               string `mapstructure:"name"`                 // 名称
	Addr               string `mapstructure:"addr"`                 // 节点地址加端口，如127.0.0.1:6379
	Password           string `mapstructure:"password"`             // 密码,无则为空
	Username           string `mapstructure:"username"`             // 账号,无则填空
	Db                 int    `mapstructure:"db"`                   // 选择db
	MaxRetries         int    `mapstructure:"max_retries"`          // 命令执行失败时，最多重试多少次，默认为3次,-1表示不重试
	PoolSize           int    `mapstructure:"pool_size"`            // 连接池大小，默认值为10*CPU个数
	MinIdleConns       int    `mapstructure:"min_idle_conns"`       // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
	MaxConnAge         int    `mapstructure:"max_conn_age"`         // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	DialTimeout        int    `mapstructure:"dial_timeout"`         // 连接建立超时时间，默认5秒
	ReadTimeout        int    `mapstructure:"read_timeout"`         // socket读取超时时间，-1为不限制超时，0为默认值，默认为3s,单位为秒
	WriteTimeout       int    `mapstructure:"write_timeout"`        // socket写超时时间，默认值跟readtimeout一致
	PoolTimeout        int    `mapstructure:"pool_timeout"`         // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒
	IdleTimeout        int    `mapstructure:"idle_timeout"`         // 闲置超时，默认5分钟，-1表示取消闲置超时检查
	IdleCheckFrequency int    `mapstructure:"idle_check_frequency"` // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理
}

type RedisClusterConf struct {
	Name               string   `mapstructure:"name"`                 // 名称
	Addrs              []string `mapstructure:"addrs"`                //集群节点地址 ip:port
	Password           string   `mapstructure:"password"`             //集群密码
	Username           string   `mapstructure:"username"`             //集群账号
	MaxRetries         int      `mapstructure:"max_retries"`          //命令执行失败时，最多重试多少次，默认为0即不重试
	RouteByLatency     bool     `mapstructure:"route_by_latency"`     //默认false,为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中选取Ping()的响应时长最短的一个节点来读数据
	RouteRandomly      bool     `mapstructure:"route_randomly"`       //默认false,为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中随机挑选一个节点来读数据
	PoolSize           int      `mapstructure:"pool_size"`            //连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
	MinIdleConns       int      `mapstructure:"min_idle_conns"`       //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；
	MaxConnAge         int      `mapstructure:"max_conn_age"`         //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	DialTimeout        int      `mapstructure:"dial_timeout"`         //连接建立超时时间，默认5秒
	ReadTimeout        int      `mapstructure:"read_timeout"`         //读超时，默认3秒， -1表示取消读超时
	WriteTimeout       int      `mapstructure:"write_timeout"`        //写超时，默认等于读超时，-1表示取消读超时
	PoolTimeout        int      `mapstructure:"pool_timeout"`         //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒
	IdleTimeout        int      `mapstructure:"idle_timeout"`         //闲置超时，默认5分钟，-1表示取消闲置超时检查
	IdleCheckFrequency int      `mapstructure:"idle_check_frequency"` //闲置连接检查的周期，无默认值，由ClusterClient统一对所管理的redis.Client进行闲置连接检查。初始化时传递-1给redis.Client表示redis.Client自己不用做周期性检查，只在客户端获取连接时对闲置连接进行处理
}

type MysqlConf struct {
	Name          string   `mapstructure:"name"`
	Master        string   `mapstructure:"master"`
	SlaveList     []string `mapstructure:"slave_list"`
	LogLevel      string   `mapstructure:"log_level"`
	SlowThreshold int      `mapstructure:"slow_threshold"`
}
