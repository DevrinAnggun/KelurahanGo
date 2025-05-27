package menu

import "fmt"

type Penduduk struct {
	Nama    string
	NIK     string
	Alamat  string
	Pekerjaan string
}

func TampilkanPenduduk() {
	pendudukList := []Penduduk{
		{"Siti Aminah", "3471066901010001", "RT.01/RW.03", "Petani"},
		{"Budi Santoso", "3471066902020002", "RT.02/RW.03", "Pedagang"},
		{"Ani Lestari", "3471066903030003", "RT.03/RW.01", "Guru"},
	}

	fmt.Println("\n--- Data Penduduk Desa Karangtengah ---")
	for i, p := range pendudukList {
		fmt.Printf("%d. Nama: %s\n", i+1, p.Nama)
		fmt.Printf("   NIK : %s\n", p.NIK)
		fmt.Printf("   Alamat: %s\n", p.Alamat)
		fmt.Printf("   Pekerjaan: %s\n\n", p.Pekerjaan)
	}
}
