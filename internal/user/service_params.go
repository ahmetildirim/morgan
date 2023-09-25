package user

type CreateServiceParams struct {
	Email    string
	Password string
}

type AuthenticateServiceParams struct {
	Email    string
	Password string
}
