package protocol

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token    string    `json:"token"`
	UserInfo *UserInfo `json:"user_info"`
}
