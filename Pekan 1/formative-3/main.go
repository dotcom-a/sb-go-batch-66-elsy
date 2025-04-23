package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Soal 1
	kalimat := "halo halo bandung"
	angka := 2021

	angkaStr := strconv.Itoa(angka)

	var newKalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	fmt.Println("\"" + newKalimat + "\" - " + angkaStr)

	//Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Print("Nilai John : ")
	hitung(nilaiJohn)
	fmt.Print("Nilai Doe : ")
	hitung(nilaiDoe)

	//Soal 3
	var tanggal = 17
	var bulan = 12
	var tahun = 2000

	var bulanString string
	switch bulan {
	case 1:
		bulanString = "Januari"
	case 2:
		bulanString = "Februari"
	case 3:
		bulanString = "Maret"
	case 4:
		bulanString = "April"
	case 5:
		bulanString = "Mei"
	case 6:
		bulanString = "Juni"
	case 7:
		bulanString = "Juli"
	case 8:
		bulanString = "Agustus"
	case 9:
		bulanString = "September"
	case 10:
		bulanString = "Oktober"
	case 11:
		bulanString = "November"
	case 12:
		bulanString = "Desember"
	}

	fmt.Println(strconv.Itoa(tanggal) + " " + bulanString + " " + strconv.Itoa(tahun))

	//Soal 4
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	}

	//Soal 5
	/*
	* Jika angka ganjil maka tampilkan Santai
	* Jika angka genap maka tampilkan Berkualitas
	* Jika angka yang sedang ditampilkan adalah kelipatan 3 DAN angka ganjil maka tampilkan I Love Coding.
	 */

	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println(strconv.Itoa(i) + " I love Coding")
		} else if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + " Berkualitas")
		} else {
			fmt.Println(strconv.Itoa(i) + " Santai")
		}
	}
}

func hitung(nilai int) {
	if nilai >= 80 {
		fmt.Println("A")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("B")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("C")
	} else if nilai >= 50 && nilai < 60 {
		fmt.Println("D")
	} else if nilai < 50 {
		fmt.Println("E")
	}
}
