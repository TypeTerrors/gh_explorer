package services

import (
	"os/user"
)

func GetUser() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.Username
}
