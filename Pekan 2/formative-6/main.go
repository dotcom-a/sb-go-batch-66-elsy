package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Soal 1
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//Soal 2
	var buah = []string{}

	addData(&buah, "Jeruk")
	addData(&buah, "Semangka")
	addData(&buah, "Mangga")
	addData(&buah, "Strawberry")
	addData(&buah, "Durian")
	addData(&buah, "Manggis")
	addData(&buah, "Alpukat")

	printAll(&buah)

	//Soal 3
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	for i, item := range dataFilm {
		fmt.Print(strconv.Itoa(i+1) + ". ")
		for key, value := range item {
			fmt.Println("\t" + key + " : " + value)
		}
	}

	//Soal 4
	var arr = []int{1, 3, 5, 7, 9}
	fmt.Println("Array Sebelum Fungsi : ")
	for _, num := range arr {
		fmt.Print(strconv.Itoa(num) + " ")
	}

	toEvenNum(&arr)
	fmt.Println("\nArray Setelah Fungsi : ")
	for _, num := range arr {
		fmt.Print(strconv.Itoa(num) + " ")
	}

	//Soal 5
	type buahStruct struct {
		nama       string
		warna      string
		adaBijinya bool
		harga      int
	}

	var Nanas buahStruct
	var Jeruk buahStruct
	var Semangka buahStruct
	var Pisang buahStruct

	Nanas.nama = "Nanas"
	Nanas.warna = "Kuning"
	Nanas.adaBijinya = false
	Nanas.harga = 9000

	Jeruk.nama = "Jeruk"
	Jeruk.warna = "Oranye"
	Jeruk.adaBijinya = true
	Jeruk.harga = 8000

	Pisang.nama = "Pisang"
	Pisang.warna = "Kuning"
	Pisang.adaBijinya = true
	Pisang.harga = 5000

	Semangka.nama = "Semangka"
	Semangka.warna = "Hijau dan Merah"
	Semangka.adaBijinya = false
	Semangka.harga = 10000

	fmt.Println("\nNama | Warna | Ada Bijinya | Harga")
	fmt.Println(Nanas)
	fmt.Println(Jeruk)
	fmt.Println(Semangka)
	fmt.Println(Pisang)

}

/*===== Func Soal 1 =====*/
func introduce(sentence *string, name string, gender string, occupation string, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak " + name + " adalah seorang " + occupation + " yang berusia " + age
	} else if gender == "perempuan" {
		*sentence = "Bu " + name + " adalah seorang " + occupation + " yang berusia " + age
	} else {
		//return empty strins
		*sentence = ""
	}
}

/*===== Func Soal 2 =====*/
func addData(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}
func printAll(buah *[]string) {
	for i := 0; i < len(*buah); i++ {
		println(strconv.Itoa(i+1) + ". " + (*buah)[i])
	}
}

/*===== Func Soal 3 =====*/
func tambahDataFilm(dataFilm *[]map[string]string, name string, duration string, genre string, year string) {
	*dataFilm = append(*dataFilm, map[string]string{"name": name, "duration": duration, "genre": genre, "year": year})
}

/*===== Func Soal 4 =====*/
func toEvenNum(listOfNum *[]int) {
	for i := range *listOfNum {
		(*listOfNum)[i] *= 2
	}
}
