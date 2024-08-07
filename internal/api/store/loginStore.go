package store

type User struct {
	Email       string
	Password    string
	AccessToken string
}

type LoginData struct {
	Error *string
	User  *User
}

var (
	LoginStore = LoginData{
		Error: nil,
		User:  nil,
	}
)
