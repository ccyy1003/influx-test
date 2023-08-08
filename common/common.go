package common

import (
	"os"
)

var (
	Addr     string
	Username string
	Password string
)

func CheckEnv() {

	Addr = os.Getenv("INFLUX_TEST_ADDR")
	if Addr == "" {
		Addr = "http://127.0.0.1:8086"
	}
	Username = os.Getenv("INFLUX_TEST_USER")
	Password = os.Getenv("INFLUX_TEST_PWD")
}
