package auth

type LoginHandlerParams struct {
	Email    string
	Password string
}

type LoginHandlerResponse struct {
	Token string `json:"token"`
}
