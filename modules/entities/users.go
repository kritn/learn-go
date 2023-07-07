package entities

/*interface*/
type UserUseCase interface {
	GetAllUsers() (t []UsersRegisterRes, err error)
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UserRepository interface {
	GetAllUsers(t *[]UsersRegisterRes) (err error)
	FindOneUser(username string, password string) (*UsersPassport, error)
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

/*model*/
type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
