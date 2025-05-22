package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== SISTEM KELURAHAN ===")
		fmt.Println("1. Tambah Penduduk")
		fmt.Println("2. Lihat Penduduk")
		fmt.Println("3. Buat Surat Keterangan")
		fmt.Println("4. Lihat Surat")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			tambahPenduduk()
		case "2":
			lihatPenduduk()
		case "3":
			buatSurat()
		case "4":
			lihatSurat()
		case "5":
			fmt.Println("Keluar dari sistem.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
