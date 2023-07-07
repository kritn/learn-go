package repositories

//มีหน้าที่ ในการรับส่ง Entities เข้าออกจาก Database หรือพูดง่ายๆ ก็คือมีหน้าที่ Query ข้อมูลจาก Database นั่นแหละ
// SQL Query

import (
	"database/sql"
	"fmt"
	"go_cleanarc/modules/entities"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(db *sql.DB) entities.UserRepository {
	return &userRepository{db}
}

func (t *userRepository) GetAllUsers(users *[]entities.UsersRegisterRes) (err error) {
	var (
		user entities.UsersRegisterRes
	)
	sql_query := `SELECT id, username,password FROM users`

	stmt, err := t.conn.Prepare(sql_query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	theRows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Query"))
	}
	for theRows.Next() {
		err := theRows.Scan(&user.Id, &user.Username, &user.Password)
		*users = append(*users, user)

		if err != nil {
			return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Scan"))
		}
	}
	err = theRows.Err()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when theRows"))
	}

	// defer stmt.Close()
	return nil
}
func (t *userRepository) FindOneUser(username string, password string) (*entities.UsersPassport, error) {
	res := new(entities.UsersPassport)

	query := `
	SELECT
	id,
	username,
	password
	FROM users
	WHERE username = ? ;
	`
	stmt, err := t.conn.Prepare(query)
	if err != nil {
		return res, fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	scan_err := stmt.QueryRow(username).Scan(&res.Id, &res.Username, &res.Password)
	if scan_err != nil {
		return res, fmt.Errorf(fmt.Sprintf("%s %s", scan_err, "user not found"))
	}

	bcrypt_ := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if bcrypt_ == nil {
		return res, nil
	}
	return res, fmt.Errorf("invalid password")
}

func (r *userRepository) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Initail a user object
	user := new(entities.UsersRegisterRes)

	var userExist entities.UsersRegisterRes

	theRows := r.conn.QueryRow("SELECT id,username FROM users WHERE username = ?;", req.Username)
	if theRows != nil {
		theRows.Scan(&userExist.Id, &userExist.Username)
	}

	if userExist.Id != 0 {
		return user, fmt.Errorf("user does not exists")
	}

	query := `INSERT INTO users(username, password) VALUES (?, ?);`

	// Query part
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return user, fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	rs, err := stmt.Exec(req.Username, req.Password)
	if err != nil {
		return user, fmt.Errorf(fmt.Sprintf("%s %s", err, "when Exec"))
	}
	id, _ := rs.LastInsertId()
	user.Id = id
	user.Username = req.Username

	// defer r.conn.Close()

	return user, nil
}
