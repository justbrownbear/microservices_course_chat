package config

import (
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

// Load загружает конфигурацию из файла .env или переданного файла конфига
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func stringToUint16(s string) (uint16, error) {
	parsed, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to uint16: %v", err)
	}

	return uint16(parsed), nil
}
