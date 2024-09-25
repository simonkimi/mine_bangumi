package xnet

import "fmt"

func GetProxyUrl(scheme string, host string, port string, useAuth bool, username string, password string) string {
	if useAuth {
		return fmt.Sprintf("%s://%s:%s@%s:%s", scheme, username, password, host, port)
	}
	return fmt.Sprintf("%s://%s:%s", scheme, host, port)
}
