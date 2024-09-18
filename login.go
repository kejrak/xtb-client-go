package xtb

// Login describes JSON login command documented here: http://developers.xstore.pro/documentation/#login
type loginRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	AppName  string `json:"appName"`
}

type LoginResponse struct {
	Response
	StreamSessionID string `json:"streamSessionId"`
}

func NewLoginCommand(user, password string) *Command {
	return &Command{
		Command: "login",
		Arguments: &loginRequest{
			UserID:   user,
			Password: password,
			AppName:  "XTB-GO",
		},
	}
}
