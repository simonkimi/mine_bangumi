package errno

const (
	Success             = 200
	BadRequest          = 400
	InternalServerError = 500

	// 10000 - 19999 用户相关错误

	ErrorUserUnauthorized  = 10000
	ErrorUserPasswordWrong = 10001
	ErrorUserTotpWrong     = 10002

	// 20000 - 29999 业务相关错误

	ErrorCancel  = 20000
	ErrorTimeout = 20001

	// 50000 - 59999 三方API错误

	ErrorApiNetwork = 50000
	ErrorApiParse   = 50001
)

func getErrorMessage(code int) string {
	switch code {
	case Success:
		return "ok"
	case InternalServerError:
		return "Internal server error"
	case ErrorUserUnauthorized:
		return "User unauthorized"
	case ErrorUserPasswordWrong:
		return "Password incorrect"
	case ErrorUserTotpWrong:
		return "TOTP incorrect"
	case ErrorCancel:
		return "Operation canceled"
	case ErrorTimeout:
		return "Operation timeout"
	case ErrorApiNetwork:
		return "Third-party API network error"
	case ErrorApiParse:
		return "Third-party API parsing error"
	default:
		return "Unknown error"
	}
}
