package main

import "fmt"

func main() {
	var berat float64
	var uangDiterima int
	var biayaCuci int
	var totalBayar int
	var kembalian int

	const TARIF_PER_KG = 7500
	const BIAYA_ANTAR = 3000

	fmt.Print("Masukkan berat cucian (kg): ")
	fmt.Scan(&berat)

	fmt.Print("Masukkan uang yang diterima: Rp")
	fmt.Scan(&uangDiterima)

	biayaCuci = int(berat*float64(TARIF_PER_KG) + 0.5)
	totalKotor := biayaCuci + BIAYA_ANTAR
	totalBayar = ((totalKotor + 99) / 100) * 100
	kembalian = uangDiterima - totalBayar

	fmt.Println("Nama : Made Kevin Arta Wijaya NIM : 42630015")
	fmt.Println("\n====STRUK LAUNDRY==== ")
	fmt.Printf("Berat cucian      : %.1f kg\n", berat)
	fmt.Printf("Tarif per kg      : Rp%d\n", TARIF_PER_KG)
	fmt.Printf("Biaya cuci        : Rp%d\n", biayaCuci)
	fmt.Printf("Biaya antar       : Rp%d\n", BIAYA_ANTAR)
	fmt.Printf("Total bayar       : Rp%d\n", totalBayar)
	fmt.Printf("Uang diterima     : Rp%d\n", uangDiterima)
	fmt.Printf("Kembalian         : Rp%d\n", kembalian)
}
