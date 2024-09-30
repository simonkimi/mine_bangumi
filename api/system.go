package api

type SystemInfo struct {
	Version                string `json:"version"`
	AppDatabaseVersion     uint   `json:"app_database_version"`
	CurrentDatabaseVersion uint   `json:"current_database_version"`
	IsSystemInit           bool   `json:"is_system_init"`
	IsLogin                bool   `json:"is_login"`
	Username               string `json:"username"`
}
