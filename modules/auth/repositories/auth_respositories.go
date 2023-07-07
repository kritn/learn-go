package respositories

//มีหน้าที่ ในการรับส่ง Entities เข้าออกจาก Database หรือพูดง่ายๆ ก็คือมีหน้าที่ Query ข้อมูลจาก Database นั่นแหละ
// SQL Query

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"go_cleanarc/modules/entities"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type authRepo struct {
	Db *sql.DB
}

func NewAuthRepository(db *sql.DB) entities.AuthRepository {
	return &authRepo{
		Db: db,
	}
}

func (r *authRepo) SignUsersAccessToken(req *entities.UsersPassport) (string, error) {
	claims := entities.UsersClaims{
		Id:       req.Id,
		Username: req.Username,
		Timeout:  (time.Now().Add(30 * time.Minute)).Format("2006-01-02 15:04:05"),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return ss, nil
}
