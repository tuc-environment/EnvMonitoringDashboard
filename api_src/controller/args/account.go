package args

type AccountRegisterArgs struct {
	Username string `json:"username" example:"bob"`
	Password string `json:"password" example:"123456"`
}

type AccountLoginArgs struct {
	Username string `json:"username" example:"bob"`
	Password string `json:"password" example:"123456"`
}

type AccountChangePasswordArgs struct {
	NewPassword string `json:"new_password" example:"123456"`
}
