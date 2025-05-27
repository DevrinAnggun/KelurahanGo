package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Surat struct {
	Nama    string `json:"nama"`
	NIK     string `json:"nik"`
	Alamat  string `json:"alamat"`
	Jenis   string `json:"jenis"`
}

func SimpanSurat(s Surat) {
	file, _ := os.ReadFile("surat.json")
	var suratList []Surat
	json.Unmarshal(file, &suratList)

	suratList = append(suratList, s)
	data, _ := json.MarshalIndent(suratList, "", "  ")
	os.WriteFile("surat.json", data, 0644)

	fmt.Println("Data surat berhasil disimpan ke file.\n")
}
