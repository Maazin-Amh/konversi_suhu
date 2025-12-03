package temperature

import "fmt"

func Fahrenheit() {
	var pilihan_konversi int
	var angka_f, hasil float64

	fmt.Println("Anda memilih Fahrenheit, anda ingin mengubah fahrenheit ke ?")
	fmt.Println("1. Fahrenheit ke Celcius")
	fmt.Println("2. Fahrenheit ke Kelvin")
	fmt.Println("3. Fahrenheit ke Reamur")

	fmt.Scan(&pilihan_konversi)

	fmt.Println("Masukan angka Fahrenheit!")
	fmt.Scan(&angka_f)

	if pilihan_konversi == 1 {
		hasil = 5.0 / 9.0 * (angka_f - 32.0)
		fmt.Println("C = ", hasil)
	} else if pilihan_konversi == 2 {
		hasil = (angka_f - 32.0) * 5.0 / 9.0 + 273.15
		fmt.Println("K = ", hasil)
	} else if pilihan_konversi == 3 {
		hasil = 4.0 / 5.0 * (angka_f - 32.0)
		fmt.Println("R = ", hasil)
	}
}
