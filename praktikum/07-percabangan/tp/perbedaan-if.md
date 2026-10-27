if : Blok kode hanya dijalankan kalau kondisinya true. Kalau false, blok dilewati

if-else : Ada dua pilihan jalur. Kondisi true menjalankan blok if, kondisi false menjalankan blok else

else if : Dipakai untuk lebih dari dua pilihan. Kondisi dicek dari atas ke bawah dan berhenti di kondisi pertama yang true. Bagian else di akhir jadi pilihan terakhir kalau semua kondisi false

switch : Memilih jalur berdasarkan satu nilai (switch x { case 1: ... }) atau berdasarkan kondisi boolean (switch { case x > 5: ... }). Hasilnya lebih rapi daripada else if yang panjang


Dua perbedaan switch di Go dan Java:
1. Soal break. Di Go, switch langsung berhenti setelah satu case cocok, jadi tidak perlu break. Di Java, tanpa break program terus lanjut ke case berikutnya (fall-through). Di Go, kalau memang ingin lanjut, harus ditulis fallthrough

2. Bentuk case. Di Go, switch boleh tanpa ekspresi dan tiap case berisi kondisi boolean (misalnya case nilai >= 80:). Satu case juga boleh memuat beberapa nilai (case 1, 3, 5:). Di Java, case hanya mencocokkan satu nilai tetap dengan ekspresi switch