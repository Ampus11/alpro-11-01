# <h1 align="center">Tugas Pendahuluan Modul 06 - “PERULANGAN WHILE”</h1>
<p align="center">[Rizky Fadlurrahman ramadhan] - [109092630002]</p>

### 1. lulus.go

```go
package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	for nilai != -1 {
		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}

		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)
	}

	fmt.Println("Program selesai.")
}
```

##### Output
![Screenshot Output Unguided](tp/lulus/output.png)

#### Deskripsi
Program ini membaca nilai ujian terus-menerus dan mencetak lulus kalau nilainya 60 ke atas. Kalau yang dimasukkan -1, program berhenti

### 2. deret.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	i := 1
	jumlah := 0
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}
```

##### Output
![Screenshot Output Unguided](tp/deret/output.png)

#### Deskripsi
Program ini menjumlahkan 1 sampai n. Angka i dimulai dari 1 lalu ditambahkan ke jumlah sampai i lebih besar dari n

### 3. duaterendah.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var nama1, nama2 string
	nilai1, nilai2 := math.MaxInt, math.MaxInt

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < nilai1 {
			nilai2, nama2 = nilai1, nama1
			nilai1, nama1 = nilai, nama
		} else if nilai < nilai2 {
			nilai2, nama2 = nilai, nama
		}
		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}
```

##### Output
![Screenshot Output Unguided](tp/duaterendah/output.png)

#### Deskripsi
Program ini mencari dua siswa dengan nilai paling kecil. Setiap kali ada nilai yang lebih kecil, data dua nilai terendah diperbarui, lalu namanya dicetak di akhir

### 4. biaya.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0
	tahun := 1
	for tahun < n {
		biaya = biaya * 1.06
		tahun++
	}

	fmt.Printf("%.0f\n", biaya)
}
```

##### Output
![Screenshot Output Unguided](tp/biaya/output.png)

#### Deskripsi
Program ini menghitung biaya kuliah tahun ke-n. Biayanya dikali 1,06 tiap tahun, lalu dicetak dibulatkan dengan `%.0f`

## Kesimpulan
Di Go tidak ada `while`, jadi perulangannya ditulis pakai `for` yang hanya berisi kondisi. Bentuk ini enak dipakai kalau jumlah putarannya belum pasti, misalnya membaca input sampai -1. Di keempat soal, perulangan ini dipakai untuk membaca nilai, menjumlah deret, mencari dua nilai terendah, dan menghitung kenaikan biaya
