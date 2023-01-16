package global

type Server struct {
	ServerHost     string            `json:"server_host"`
	ServerPort     int               `json:"server_port"`
	Log            Log               `json:"log"`
	Mysql          Mysql             `json:"mysql"`
	Redis          Redis             `json:"redis"`
	Tasks          map[string]string `json:"tasks"`
	UploadFilePath string            `json:"upload_file_path"`
}

// Mysql配置
type Mysql struct {
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     int    `json:"mysql_port"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	DBName        string `json:"db_name"`
	MaxIdleConns  int    `json:"max_idle_conns"`
	MaxOpenConns  int    `json:"max_open_conns"`
}

type Redis struct {
	RedisHost      string `json:"redis_host"`
	RedisPort      int    `json:"redis_port"`
	RedisPassword  string `json:"redis_password"`
	PoolMaxIdle    int    `json:"pool_max_idle"`
	PoolMaxActive  int    `json:"pool_max_active"`
	RedisDB        int    `json:"redis_db"`
	ConnectTimeOut int    `json:"connect_time_out"`
	IdleTimeout    int    `json:"idle_timeout"`
}
type Log struct {
	LogLevel string `json:"log_level"`
	LogPath  string `json:"log_path"`
}
