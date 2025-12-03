package main

import (
	"KONVERSI_SUHU/temperature"
	"fmt"
)

func program() {
	var pilihan_input int
	var status_check bool = true
	var status_ulang string

	for status_check {

		fmt.Println("Masukan jenis yang ingin anda ubah skala pengukuran :")
		fmt.Println("1. Celcius")
		fmt.Println("2. Fahrenheit")
		fmt.Println("3. Kelvin")
		fmt.Println("4. Reamur")
		fmt.Scan(&pilihan_input)

		if pilihan_input == 1 {
			temperature.Celcius()
		} else if pilihan_input == 2 {
			temperature.Fahrenheit()
		} else if pilihan_input == 3 {
			temperature.Kelvin()
		} else if pilihan_input == 4 {
			temperature.Reamur()
		}

		fmt.Println("Apakah anda ingin menginput lagi ?")
		fmt.Println("iya | tidak")
		fmt.Scan(&status_ulang)

		if status_ulang == "iya" || status_ulang == "y" {
			status_check = true
		} else {
			status_check = false
			fmt.Println("Terima Kasih 🙏🏻")
		}
	}

}
