package main

import "fmt"

func main() {
    var x, y float64
    var hasil float64

    fmt.Print("Masukkan nilai x: ")
    fmt.Scanln(&x)

    fmt.Print("Masukkan nilai y: ")
    fmt.Scanln(&y)

    //cek apakah nilai x berisi bilangan

    if x == 0 {
        fmt.Println("Anda belum memasukkan angka X")
    //cek apakah nilai y berisi bilangan    
    }else if y == 0 {
        fmt.Println("Anda belum memasukan angka y")
     //cek kedua nilai
    }else if x == 0 || y == 0{
        fmt.Println("Anda belum memasukan angka x dan y ")
    }else{
        hasil = 1 / (3*x *x+10)+10*y+7
        fmt.Println("Jika X:",x,"Maka:",hasil)
    }              

  
}
