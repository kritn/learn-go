package utils

import (
	"errors"
	"fmt"
	"go_cleanarc/configs"
)

func ConnectionUrlBuilder(stuff string, cfg *configs.DBConfig) (string, error) {
	var url string

	switch stuff {
	case "mariaDB":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
