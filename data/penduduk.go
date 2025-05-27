package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Penduduk struct {
	Nama      string `json:"nama"`
	NIK       string `json:"nik"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
}

func LoadPenduduk() []Penduduk {
	file, err := os.ReadFile("penduduk.json")
	if err != nil {
		return []Penduduk{}
	}
	var data []Penduduk
	json.Unmarshal(file, &data)
	return data
}

func SimpanPenduduk(pendudukList []Penduduk) {
	data, _ := json.MarshalIndent(pendudukList, "", "  ")
	os.WriteFile("penduduk.json", data, 0644)
}

func TambahPenduduk(p Penduduk) {
	penduduk := LoadPenduduk()
	penduduk = append(penduduk, p)
	SimpanPenduduk(penduduk)
	fmt.Println("Data penduduk berhasil ditambahkan.\n")
}
