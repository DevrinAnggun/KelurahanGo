package menu

import (
	"fmt"
	"time"
)

func TampilkanBeranda() {
	fmt.Println("\n====================================")
	fmt.Println("  Selamat Datang di KelurahanGo CLI ")
	fmt.Println("====================================")
	fmt.Println("Tanggal dan Waktu Saat Ini:", time.Now().Format("02 Jan 2006 15:04"))
	fmt.Println()
}
