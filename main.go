package main

import (
	"fmt"
	"kelurahango/menu"
	"kelurahango/utils"
	"os"
)


func main() {
	for {
		menu.TampilkanBeranda()
		fmt.Println("=== Menu Utama KelurahanGo ===")
		fmt.Println("1. Informasi Desa")
		fmt.Println("2. Pelayanan Surat")
		fmt.Println("3. Data Penduduk")
		fmt.Println("4. Informasi Kontak")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")

		choice := utils.InputInt()

		switch choice {
		case 1:
			menu.TampilkanInformasiDesa()
		case 2:
			menu.PelayananSurat()
		case 3:
			menu.TampilkanKontak()
		case 4:
			menu.TampilkanPenduduk()
		case 5:
			fmt.Println("Terima kasih telah menggunakan KelurahanGo.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
