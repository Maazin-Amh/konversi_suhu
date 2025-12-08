package temperature

import "fmt"

func Celcius() {
	var pilihan_konversi int
	var angka_c, angka_f, angka_k, angka_r, hasil float64

	fmt.Println("Anda memilih Celcius, anda ingin mengubah celcius ke ?")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Celcius ke Reamur")
	fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}
	fmt.Scan(pilihan_konversi)

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


	fmt.Println("Anda memilih Fahrenheit, anda ingin mengubah fahrenheit ke ?")
	fmt.Println("1. Fahrenheit ke Celcius")
	fmt.Println("2. Fahrenheit ke Kelvin")
	fmt.Println("3. Fahrenheit ke Reamur")
	fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}

	fmt.Println("Masukan angka Fahrenheit!")
	fmt.Scan(&angka_f)

	if pilihan_konversi == 1 {
		hasil = 5.0 / 9.0 * (angka_f - 32.0)
		fmt.Println("C = ", hasil)
	} else if pilihan_konversi == 2 {
		hasil = (angka_f-32.0)*5.0/9.0 + 273.15
		fmt.Println("K = ", hasil)
	} else if pilihan_konversi == 3 {
		hasil = 4.0 / 5.0 * (angka_f - 32.0)
		fmt.Println("R = ", hasil)
	}

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

	fmt.Println("Anda memilih Fahrenheit, anda ingin mengubah kelvin ke")
	fmt.Println("1. Reamur ke Celcius")
	fmt.Println("2. Reamur ke Fahrenheit")
	fmt.Println("3. Reamur ke Kelvin")
	fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}
	fmt.Println("pilihan_konversi")

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
	fmt.Println("Anda memilih Fahrenheit, anda ingin mengubah kelvin ke")
	fmt.Println("1. Reamur ke Celcius")
	fmt.Println("2. Reamur ke Fahrenheit")
	fmt.Println("3. Reamur ke Kelvin")
	fmt.Scan(&pilihan_konversi)

	if pilihan_konversi < 1 || pilihan_konversi > 3 {
		fmt.Println("input yang anda masukan tidak ada di pilihan")
		return
	}
	fmt.Println("pilihan_konversi")

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




