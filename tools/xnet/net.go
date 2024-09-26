package xnet

import (
	"fmt"
	"github.com/simonkimi/minebangumi/tools/xstring"
)

func GetProxyUrl(scheme string, host string, port string, username string, password string) string {
	if !xstring.IsEmptyOrWhitespace(username) {
		return fmt.Sprintf("%s://%s:%s@%s:%s", scheme, username, password, host, port)
	}
	return fmt.Sprintf("%s://%s:%s", scheme, host, port)
}
