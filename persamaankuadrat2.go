package main

import "fmt"

func main() {
var x float64
fmt.Print("Masukkan nilai x: ")
fmt.Scanln(&x)

// cek x apakah sudah berisi bilangan

hasil := (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
fmt.Println(hasil)




}
