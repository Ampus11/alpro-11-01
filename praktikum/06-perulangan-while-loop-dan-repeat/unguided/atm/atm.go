package main

import "fmt"

func main() {
	benar := "1234"
	saldo := 10000000
	kesempatan := 0
	login := false

	for kesempatan < 3 && !login {
		var password string
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == benar {
			login = true
		} else {
			fmt.Println("Password salah. Silakan coba lagi.")
			kesempatan++
		}
	}

	if !login {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Println("Login berhasil!")

	pilihan := 0
	for pilihan != 3 {
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("Saldo Anda: Rp", saldo)
		} else if pilihan == 2 {
			var jumlah int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= saldo {
				saldo -= jumlah
				fmt.Println("Anda berhasil menarik uang sebesar Rp", jumlah)
				fmt.Println("Saldo tersisa: Rp", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		}
	}
}