package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// InputInt membaca input dan mengubahnya menjadi integer
func InputInt() int {
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan input. Coba lagi.")
			continue
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}
		return num
	}
}

// InputString membaca input string dari pengguna
func InputString() string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi kesalahan input. Coba lagi.")
		return InputString()
	}
	return strings.TrimSpace(input)
}
