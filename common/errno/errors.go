package errno

var (
	OK = NewResult(0, "OK")

	// 错误码规则
	// 总共五位数
	// 服务级错误码		模块级错误码		具体错误码
	// eg: 1				01				01
	// 服务级错误码: 1位数表示, 如: 1为系统级错误;		2为普通错误
	// 模块级错误码: 2位数表示, 如: 01为用户模块; 		02为订单模块
	// 具体错误码:   2位数表示, 如: 01为邮箱不合法; 	02为验证码错误
	ErrServer       = NewResult(10001, "服务异常, 请稍后重试")
	ErrParam        = NewResult(10002, "参数有误")
	ErrSignParam    = NewResult(10003, "签名参数有误")
	ErrSignMethod   = NewResult(10004, "签名算法有误")
	ErrInvalidToken = NewResult(10005, "Token有误")
	ErrExpiredToken = NewResult(10006, "Token已失效")

	ErrUserEmail   = NewResult(20101, "用户邮箱不合法")
	ErrUserUnlogin = NewResult(20102, "请先登录")
)
