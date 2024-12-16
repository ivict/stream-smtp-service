package tests

import (
	"os"
)

func init() {
	if _, err := os.Stat(".env"); err != nil {
		os.Chdir("..") // .env file location
	}
}
