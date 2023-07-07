package usecases

// มีหน้าที่ รับมือกับ Logic ต่างๆ ก่อนที่จะส่งข้อมูลเข้าออก Database เช่น Search, Sort, Hash

import (
	"fmt"
	"go_cleanarc/modules/entities"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo entities.UserRepository
}

func NewUserUseCase(repo entities.UserRepository) entities.UserUseCase {
	return &userUseCase{repo}
}

func (t *userUseCase) GetAllUsers() (user []entities.UsersRegisterRes, err error) {
	var users []entities.UsersRegisterRes
	handleErr := t.userRepo.GetAllUsers(&users)
	return users, handleErr
}

func (u *userUseCase) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.userRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
