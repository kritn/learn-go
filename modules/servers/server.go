package servers

import (
	"fmt"
	"go_cleanarc/configs"

	// "go_cleanarc/modules/servers"
	"go_cleanarc/pkg/databases"
	//_ "go_cleanarc/pkg/routes"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func StartServer() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("error .env file "))
	}
	cfg := new(configs.DBConfig)
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_User")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.Password = os.Getenv("DB_PASSWORD")

	// New Database
	db, err := databases.NewDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	//setup routes
	r := SetupRouter()

	// running
	r.Run(":8080")
}
