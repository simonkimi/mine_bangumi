package api

type SystemInfo struct {
	Version    string `json:"version"`
	IsFirstRun bool   `json:"is_first_run"`
	IsLogin    bool   `json:"is_login"`
}

type InitUserForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
