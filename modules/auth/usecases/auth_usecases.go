package usecases

// มีหน้าที่ รับมือกับ Logic ต่างๆ ก่อนที่จะส่งข้อมูลเข้าออก Database เช่น Search, Sort, Hash

import (
	"go_cleanarc/modules/entities"
)

type authUse struct {
	AuthRepo  entities.AuthRepository
	UsersRepo entities.UserRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository, usersRepo entities.UserRepository) entities.AuthUsecase {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUse) Login(req *entities.UsersCredentials) (*entities.UsersLoginRes, error) {
	user, err := u.UsersRepo.FindOneUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, errors.New("error, password is invalid")
	// }

	token, err := u.AuthRepo.SignUsersAccessToken(user)
	if err != nil {
		return nil, err
	}
	res := &entities.UsersLoginRes{
		AccessToken: token,
	}
	return res, nil
}
