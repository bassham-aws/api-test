package environment

import (
	"os"
)

func AppName() string {
	return "api-test"
}

func AppEnv() string {
	return os.Getenv("APP_ENV")
}

func ListenAddress() string {
	return os.Getenv("LISTEN_ADDRESS")
}
