package main

import "fmt"

func main() {
	fmt.Println("=====A.9.1. Deklarasi Variabel Beserta Tipe Data=====")
	var fistName string = "Ardin"
	fmt.Println("Hello, my name is", fistName)

	fmt.Println("=====A.9.3. Deklarasi Variabel Tanpa Tipe Data=====")
	var lastName = "Nugraha"
	middleName := "Middle"
	fmt.Println(lastName, middleName)

	fmt.Println("=====A.9.4. Deklarasi Multi Variabel=====")
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	empat, lima := 4, 5
	fmt.Println("Multi Variabel :", fourth, fifth, sixth, empat, lima)

	fmt.Println("=====A.9.6. Deklarasi Variabel Menggunakan Keyword new=====")
	name := new(string)
	fmt.Println(name) // 0x20818a220
	fmt.Println("isi keyword new", *name)
}
