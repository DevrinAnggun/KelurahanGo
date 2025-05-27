package menu

import (
	"fmt"
	"kelurahango/utils"
)

func PelayananSurat() {
	fmt.Println("\n--- Pelayanan Surat Kelurahan ---")
	fmt.Println("1. Surat Keterangan Domisili")
	fmt.Println("2. Surat Pengantar KTP, SIM, SKCK, Akta Kelahiran")
	fmt.Println("3. Surat Keterangan Usaha")
	fmt.Println("4. Surat Pengantar Beasiswa")
	fmt.Println("5. Surat Pengantar Ijin Menikah")
	fmt.Println("6. Kembali ke Menu Utama")

	fmt.Print("Pilih jenis surat (1-4): ")
	choice := utils.InputInt()

	var nama, nik, alamat string

	if choice >= 1 && choice <= 3 {
		fmt.Print("Masukkan Nama Anda   : ")
		nama = utils.InputString()
		fmt.Print("Masukkan NIK         : ")
		nik = utils.InputString()
		fmt.Print("Masukkan Alamat Anda : ")
		alamat = utils.InputString()

		fmt.Println("\n--- Konfirmasi Permintaan Surat ---")
		fmt.Printf("Nama      : %s\n", nama)
		fmt.Printf("NIK       : %s\n", nik)
		fmt.Printf("Alamat    : %s\n", alamat)

		switch choice {
		case 1:
			fmt.Println("Jenis Surat: Surat Keterangan Domisili")
		case 2:
			fmt.Println("Jenis Surat: Surat Pengantar KTP, SIM, SKCK, Akta Kelahiran")
		case 3:
			fmt.Println("Jenis Surat: Surat Keterangan Usaha")
		case 4:
			fmt.Println("Jenis Surat: Surat Pengantar Beasiswa")
		case 5:
			fmt.Println("Jenis Surat: Surat Pengantar Ijin Menikah")
		case 6:
			fmt.Println("Kembali ke Menu Utama")	
		}

		fmt.Println("Status    : BERHASIL diajukan. Silakan ambil di kantor kelurahan.\n")

	} else if choice == 4 {
		return
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.\n")
	}
}
