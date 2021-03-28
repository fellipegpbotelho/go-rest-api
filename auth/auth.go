package auth

type Auther interface {
	AuthenticateUser(email string, password string) bool
}

type Auth struct {
	email    string
	password string
}

func BuildAuth() *Auth {
	return &Auth{
		email:    "clark@dc.com",
		password: "iamsuperman",
	}
}

func (auth *Auth) AuthenticateUser(email string, password string) bool {
	return auth.email == email && auth.password == password
}
