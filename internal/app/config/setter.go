package config

import (
	"github.com/simonkimi/minebangumi/pkg/hash"
	"github.com/simonkimi/minebangumi/tools/xstring"
)

func UpdateUser(c Config, username *string, password *string) string {
	if username != nil {
		c.SetString(UserUsername, *username)
	}
	if password != nil {
		c.SetString(UserPassword, *password)
	}

	var token string
	if xstring.IsEmptyOrWhitespace(c.GetString(UserPassword)) {
		c.SetString(UserApiToken, "")
		token = ""
	} else {
		token = hash.GenerateRandomKey(40)
		c.SetString(UserApiToken, token)
	}
	c.Save()
	return token
}
