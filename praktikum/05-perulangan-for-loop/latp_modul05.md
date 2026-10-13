# <h1 align="center">Tugas Pendahuluan Modul 05 - “PERULANGAN FOR” </h1>
<p align="center">[Rizky Fadlurrahman ramadhan] - [109092630002]</p>

### 1. ganjil.go

```.go
package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jumlah += i
	}
	fmt.Println()
	fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/ganjil/output.png)


#### Deskripsi
Program ini mencetak bilangan ganjil dari 1 sampai n beserta jumlahnya. Variabel i mulai dari 1 dan naik 2 setiap putaran, sehingga hanya bilangan ganjil yang dicetak dan dijumlahkan

### 2. perkalian.go

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/perkalian/output.png)


#### Deskripsi
Program ini mencetak tabel perkalian bilangan a dari a x 1 sampai a x b. Variabel i berjalan dari 1 sampai b, dan tiap putaran mencetak satu baris berformat a x i = hasil


## Kesimpulan
Perulangan for dipakai untuk menjalankan perintah yang sama berulang kali tanpa menulis kode berkali-kali. Di Go, for adalah satu-satunya perulangan, dan bisa ditulis dengan tiga komponen (jumlah putaran pasti) atau hanya kondisi (seperti while). Pada Soal 3 dan 4, for dipakai untuk mencetak bilangan ganjil beserta jumlahnya dan untuk membuat tabel perkalian