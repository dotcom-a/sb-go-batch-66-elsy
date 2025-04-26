package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Soal 1
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//Soal 2
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//Soal 3
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(name string, duration string, genre string, year string) {
		dataFilm = append(dataFilm, map[string]string{"name": name, "duration": duration, "genre": genre, "year": year})
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	//Soal 4
	var num5 = factorial(5)
	var num7 = factorial(7)
	fmt.Println(strconv.Itoa(num5))
	fmt.Println(strconv.Itoa(num7))

	//Soal 5
	var luasLigkaran float64
	var kelilingLingkaran float64

	fmt.Println("Before Pointer : ")

	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

	LuasLingkaran(&luasLigkaran)
	KelilingLingkaran(&kelilingLingkaran)

	fmt.Println("After Pointer : ")
	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

}

func introduce(name string, gender string, occupation string, age string) (saySmthg string) {
	if gender == "laki-laki" {
		saySmthg = "Pak " + name + " adalah seorang " + occupation + " yang berusia " + age
	} else if gender == "perempuan" {
		saySmthg = "Bu " + name + " adalah seorang " + occupation + " yang berusia " + age
	} else {
		//return empty strins
		saySmthg = ""
	}
	return saySmthg
}

func buahFavorit(name string, fruits ...string) string {
	var buahFavoritJohn string
	for _, fruit := range fruits {
		buahFavoritJohn += "\"" + fruit + "\", "
	}
	return "halo nama saya " + name + " dan buah favorit saya adalah " + buahFavoritJohn
}

func factorial(num int) int {
	total := num
	for i := num - 1; i > 0; i-- {
		total = total * i
	}
	return total
}

func LuasLingkaran(num *float64) {
	*num = 10.1
}

func KelilingLingkaran(num *float64) {
	*num = 40.5
}
