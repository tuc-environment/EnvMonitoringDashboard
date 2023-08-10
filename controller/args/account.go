package args

type AccountRegisterArgs struct {
	Username string `example:"bob"`
	Password string `example:"123456"`
}

type AccountLoginArgs struct {
	Username string `example:"bob"`
	Password string `example:"123456"`
}
