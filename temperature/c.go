package temperature

import "fmt"

func Celcius() {
	var pilihan_konversi int
	var angka_c, hasil float64

	fmt.Println("Anda memilih Celcius, anda ingin mengubah celcius ke ?")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Celcius ke Reamur")
	fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}

	fmt.Println("Masukan angka celciusnya!")
	fmt.Scan(&angka_c)

	if pilihan_konversi == 1 {
		hasil = (9.0 / 5.0 * angka_c) + 32.0
		fmt.Println("F = ", hasil)
	} else if pilihan_konversi == 2 {
		hasil = angka_c + 273.15
		fmt.Println("K = ", hasil)
	} else if pilihan_konversi == 3 {
		hasil = 4.0 / 5.0 * angka_c
		fmt.Println("R = ", hasil)
	}
}
