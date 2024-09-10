package exception

const (
	Success             = 200
	InternalServerError = 500

	// 10000 - 19999 用户相关错误

	ErrorUserUnauthorized  = 10000
	ErrorUserPasswordWrong = 10001
	ErrorUserTotpWrong     = 10002
)

func getErrorMessage(code int) string {
	switch code {
	case Success:
		return "ok"
	case InternalServerError:
		return "服务器内部错误"
	case ErrorUserUnauthorized:
		return "用户未授权"
	case ErrorUserPasswordWrong:
		return "密码错误"
	case ErrorUserTotpWrong:
		return "TOTP错误"
	default:
		return "未知错误"
	}
}
