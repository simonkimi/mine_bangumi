package api

type SystemInfo struct {
	Version    string `json:"version"`
	IsInitUser bool   `json:"is_init_user"`
	IsLogin    bool   `json:"is_login"`
}

type UserCredentialForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DownloaderForm struct {
	Type     string `json:"type" binding:"required"`
	Api      string `json:"api" binding:"required"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
