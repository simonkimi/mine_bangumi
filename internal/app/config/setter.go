package config

func UpdateUserCredential(username string, password string) {
	config.keys.userUsername.setValue(username)
	config.keys.userPassword.setValue(password)
	saveConfig()
}
