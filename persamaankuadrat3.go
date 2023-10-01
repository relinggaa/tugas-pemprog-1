package main

import "fmt"

func main() {
	var x int
	var d1, d2, d3 int

	fmt.Println("Masukkan nilai x = ")

	fmt.Scanln(&x)
	// cek x apakah sudah berisi bilangan
	if x==0{
		fmt.Println("Anda belum memasukan angka")

	}

	// Untuk bilangan diluar jangkauan 0-999
	if x < 0 || x > 999 {
		fmt.Println("Angka di luar jangkauan")
	}

	// bilangan dibawah 10
	if x < 10 {
		fmt.Println(
			d1 == 0,
			d2 == 0,
			d3 == x,
		)
	}

	// Untuk bilangan 0-999
	if x >= 0 && x <= 999 {
		modulus := x // Simpan nilai x dalam variabel sementara
		d3 = modulus  % 10
		modulus /= 10
		d2 = modulus  % 10
		modulus /= 10
		d1 = modulus  % 10
		
	}
	fmt.Println("Jika x =", x, "Maka d1 =", d1, "d2 =", d2, "d3 =", d3)
	
}
