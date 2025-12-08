package temperature

import (
	"fmt"

	"golang.org/x/text/encoding/korean"
)

func Reamur() {
	var pilihan_konversi int
	var angka_r, hasil float64

	fmt.Println("Anda memilih Fahrenheit, anda ingin mengubah kelvin ke")
	fmt.Println("1. Reamur ke Celcius")
	fmt.Println("2. Reamur ke Fahrenheit")
	fmt.Println("3. Reamur ke Kelvin")

	if pilihan_konversi < 1 || pilihan_konversi > 3{
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	} 


	fmt.Scan(&pilihan_konversi)

	fmt.Println("Masukan angka Reamur")
	fmt.Scan(&angka_r)

	if pilihan_konversi == 1 {
		hasil = 5.0 / 4.0 * angka_r
		fmt.Println("C = ", hasil)
	} else if pilihan_konversi == 2 {
		hasil = 9.0/4.0*angka_r + 32.0
		fmt.Println("F = ", hasil)
	} else if pilihan_konversi == 3 {
		hasil = 5.0/4.0*angka_r + 273.15
		fmt.Println("R = ", hasil)
	}
}
