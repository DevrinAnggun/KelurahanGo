package auth

import (
	"fmt"
	"kelurahango/utils"
)

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Dummy hardcoded admin
var admins = []Admin{
	{"admin", "1234"},
}

func LoginAdmin() bool {
	fmt.Print("Username: ")
	username := utils.InputString()
	fmt.Print("Password: ")
	password := utils.InputString()

	for _, admin := range admins {
		if admin.Username == username && admin.Password == password {
			fmt.Println("Login berhasil!\n")
			return true
		}
	}
	fmt.Println("Login gagal. Username atau password salah.\n")
	return false
}
