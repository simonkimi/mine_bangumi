package errno

import "net/http"

const (
	Success             = 200
	BadRequest          = 400
	Unauthorized        = 401
	InternalServerError = 500

	// 10000 - 19999 用户相关错误

	ErrorUserPasswordWrong = 10001
	ErrorUserTotpWrong     = 10002

	// 20000 - 29999 业务相关错误

	ErrorCancel  = 20000
	ErrorTimeout = 20001

	// 50000 - 59999 三方API错误

	ErrorApiNetwork = 50000
	ErrorApiParse   = 50001
)

func GetErrorMessage(code int) string {
	// 400 http错误
	switch code {
	case http.StatusBadRequest:
		return "Bad request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not found"
	case http.StatusMethodNotAllowed:
		return "Method not allowed"
	case http.StatusNotAcceptable:
		return "Not acceptable"
	case http.StatusRequestTimeout:
		return "Request timeout"
	}

	switch code {
	case Success:
		return "ok"
	case InternalServerError:
		return "Internal server error"
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
