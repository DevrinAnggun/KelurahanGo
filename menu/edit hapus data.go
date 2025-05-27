package menu

import (
	"fmt"
	"kelurahango/data"
	"kelurahango/utils"
)

func Tampilkanpenduduk() {
	pendudukList := data.LoadPenduduk()

	fmt.Println("\n--- Data Penduduk ---")
	for i, p := range pendudukList {
		fmt.Printf("%d. %s - %s\n", i+1, p.Nama, p.NIK)
	}
	fmt.Println("1. Cari berdasarkan NIK")
	fmt.Println("2. Edit Data Penduduk")
	fmt.Println("3. Hapus Data Penduduk")
	fmt.Println("4. Tambah Penduduk Baru")
	fmt.Println("5. Kembali")

	switch utils.InputInt() {
	case 1:
		fmt.Print("Masukkan NIK: ")
		nik := utils.InputString()
		for _, p := range pendudukList {
			if p.NIK == nik {
				fmt.Printf("Nama: %s, Alamat: %s, Pekerjaan: %s\n\n", p.Nama, p.Alamat, p.Pekerjaan)
				return
			}
		}
		fmt.Println("Data tidak ditemukan.\n")

	case 2:
		fmt.Print("Masukkan NIK yang ingin diedit: ")
		nik := utils.InputString()
		for i, p := range pendudukList {
			if p.NIK == nik {
				fmt.Print("Nama Baru: ")
				pendudukList[i].Nama = utils.InputString()
				fmt.Print("Alamat Baru: ")
				pendudukList[i].Alamat = utils.InputString()
				fmt.Print("Pekerjaan Baru: ")
				pendudukList[i].Pekerjaan = utils.InputString()
				data.SimpanPenduduk(pendudukList)
				fmt.Println("Data berhasil diedit.\n")
				return
			}
		}
		fmt.Println("NIK tidak ditemukan.\n")

	case 3:
		fmt.Print("Masukkan NIK yang ingin dihapus: ")
		nik := utils.InputString()
		for i, p := range pendudukList {
			if p.NIK == nik {
				pendudukList = append(pendudukList[:i], pendudukList[i+1:]...)
				data.SimpanPenduduk(pendudukList)
				fmt.Println("Data berhasil dihapus.\n")
				return
			}
		}
		fmt.Println("NIK tidak ditemukan.\n")

	case 4:
		var p data.Penduduk
		fmt.Print("Nama: ")
		p.Nama = utils.InputString()
		fmt.Print("NIK: ")
		p.NIK = utils.InputString()
		fmt.Print("Alamat: ")
		p.Alamat = utils.InputString()
		fmt.Print("Pekerjaan: ")
		p.Pekerjaan = utils.InputString()
		data.TambahPenduduk(p)

	case 5:
		return
	}
}
