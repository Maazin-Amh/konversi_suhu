package temperature

import (
	"fmt"
)

func Kelvin() {
	var pilihan_konversi int
	var angka_k, hasil float64

	fmt.Println("Anda memilih Kelvin, anda ingin mengubah kelvin ke")
	fmt.Println("1. Kelvin ke Celcius")
	fmt.Println("2. Kelvin ke Fahrenheit")
	fmt.Println("3. Kelvin ke Reamur")
		fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}

	fmt.Println("Masukan angka Kelvin")
	fmt.Scan(&angka_k)

	if pilihan_konversi == 1 {
		hasil = angka_k - 273.15
		fmt.Println("C = ", hasil)
	} else if pilihan_konversi == 2 {
		hasil = 5.0/9.0*(angka_k-273.15) + 32.0
		fmt.Println("F = ", hasil)
	} else if pilihan_konversi == 3 {
		hasil = 4.0 / 5.0 * (angka_k - 273.15)
		fmt.Println("R = ", hasil)
	}
}
