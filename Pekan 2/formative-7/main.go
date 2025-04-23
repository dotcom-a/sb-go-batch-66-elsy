package main

import (
	"fmt"
	"strconv"
)

// Struct Soal 1
type Mahasiswa struct {
	Nama string
	Nim  string
	Usia int
}

//Struct Soal 2

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) LuasSegitiga() {
	fmt.Print("\nLuas Segitiga : ")
	fmt.Println((s.alas * s.tinggi) / 2)
}

func (s persegi) LuasPersegi() {
	fmt.Print("\nLuas Persegi : ")
	fmt.Println(s.sisi * s.sisi)
}

func (s persegiPanjang) LuasPersegiPanjang() {
	fmt.Print("\nLuas Persegi Panjang : ")
	fmt.Println(s.panjang * s.lebar)
}

// Struct Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (s *phone) changeColor(color string) {
	s.colors = append(s.colors, color)
}

// struct soal 4
type movies struct {
	title, genre   string
	year, duration int
}

// struct soal 5
type cat struct {
}

type dog struct {
}

type hewan interface {
	suara()
}

func (c cat) SuaraKucing() {
	fmt.Println("Meow")
}

func (d dog) SuaraAnjing() {
	fmt.Println("Woof")
}

func main() {
	//Soal 1
	var Mahasiswa Mahasiswa
	Mahasiswa.Nama = "Kassandra"
	Mahasiswa.Nim = "339129301"
	Mahasiswa.Usia = 20

	fmt.Println("Nama  :", Mahasiswa.Nama)
	fmt.Println("Nim :", Mahasiswa.Nim)
	fmt.Println("Usia :", Mahasiswa.Usia)

	//Soal 2
	var segitiga = segitiga{2, 5}
	segitiga.LuasSegitiga()

	var persegi = persegi{3}
	persegi.LuasPersegi()

	var persegiPanjang = persegiPanjang{5, 10}
	persegiPanjang.LuasPersegiPanjang()

	//Soal 3
	var phone phone
	phone.name = "x5"
	phone.brand = "samsung"
	phone.year = 2000
	phone.changeColor("Blue")
	phone.changeColor("Black")

	fmt.Println(phone)

	//Soal 4
	var dataFilm = []map[string]string{}

	var movie1 = movies{"LOTR", "action", 1999, 2}
	var movie2 = movies{"avenger", "action", 2019, 2}
	var movie3 = movies{"spiderman", "action", 2004, 2}
	var movie4 = movies{"juon", "horror", 2004, 2}

	tambahDataFilm(&dataFilm, movie1)
	tambahDataFilm(&dataFilm, movie2)
	tambahDataFilm(&dataFilm, movie3)
	tambahDataFilm(&dataFilm, movie4)

	for i, item := range dataFilm {
		fmt.Print(strconv.Itoa(i+1) + ". ")
		for key, value := range item {
			fmt.Println("\t" + key + " : " + value)
		}
	}

	//Soal 5
	var kucing cat
	kucing.SuaraKucing()

	var anjing dog
	anjing.SuaraAnjing()
}

/*===== Func Soal 4 =====*/
func tambahDataFilm(dataFilm *[]map[string]string, m movies) {
	*dataFilm = append(*dataFilm, map[string]string{"title": m.title, "duration": strconv.Itoa(m.duration), "genre": m.genre, "year": strconv.Itoa(m.year)})
}
