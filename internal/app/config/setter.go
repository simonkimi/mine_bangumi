package config

import "github.com/simonkimi/minebangumi/pkg/hash"

func UpdateUser(c Config, username *string, password *string) string {
	if username != nil {
		c.SetString(UserUsername, *username)
	}
	if password != nil {
		c.SetString(UserPassword, *password)
	}
	token := hash.GenerateRandomKey(40)
	c.SetString(UserApiToken, token)
	c.Save()
	return token
}
