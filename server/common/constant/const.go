package constant

const (
	LoggerKey = "logger"
	MenuAct   = "show_menu"
)
const (
	UserDefaultPassword = "123456"
	UserDefaultRole     = "default_role"
)

const (
	TokenPeriod    = 30           // token有效时间30分钟  分钟
	JwTPeriod      = 14 * 24 * 60 // Jwt最大有效期为14天，每14天一定要登录一次 分钟
	UserContextKey = "userId"
	TokenHeader    = "Access-Token"
)

const (
	CasbinTypeP = "p"
	CasbinTypeG = "g"
)
