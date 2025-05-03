package main

import (
	"fmt"
	"math"
	"strconv"

	. "eigth-project/lib"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (s persegiPanjang) luas() int {
	return s.lebar * s.panjang
}

func (s persegiPanjang) keliling() int {
	return (s.lebar + s.panjang) * 2
}

func (s tabung) volume() float64 {
	return math.Pi * s.tinggi * math.Pow(s.jariJari, 2)
}

func (s tabung) luasPermukaan() float64 {
	return 2 * math.Pi * s.jariJari * (s.tinggi + s.jariJari)
}

func (s balok) volume() float64 {
	return float64(s.lebar) * float64(s.panjang) * float64(s.tinggi)
}

func (s balok) luasPermukaan() float64 {
	return float64(2 * ((s.panjang * s.lebar) + (s.panjang * s.tinggi) + (s.tinggi * s.lebar)))
}

/* ===  struct 2 ===*/
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneDetails interface {
	printDetails()
}

func (s phone) printDetails() {
	fmt.Println("Name : ", s.name)
	fmt.Println("Brand : ", s.brand)
	fmt.Println("Year : ", strconv.Itoa(s.year))
	fmt.Println("Colors : ", s.colors)

}

// Struct 3
type persegi struct {
	sisi int
}

func main() {
	//Soal 1
	var bandunDatar hitungBangunDatar
	bandunDatar = segitigaSamaSisi{3, 3}
	fmt.Print("Luas Segitiga : ")
	fmt.Println(bandunDatar.luas())
	fmt.Print("Keliling Segitiga : ")
	fmt.Println(bandunDatar.keliling())

	bandunDatar = persegiPanjang{4, 8}
	fmt.Print("Luas Persegi Panjang : ")
	fmt.Println(bandunDatar.luas())
	fmt.Print("Keliling Persegi Panjang : ")
	fmt.Println(bandunDatar.keliling())

	var bangunRuang hitungBangunRuang
	bangunRuang = tabung{5, 10}
	fmt.Print("Volume Tabung : ")
	fmt.Println(bangunRuang.volume())
	fmt.Print("Luas Permukaan Tabung : ")
	fmt.Println(bangunRuang.luasPermukaan())

	bangunRuang = balok{3, 3, 3}
	fmt.Print("Volume Balok : ")
	fmt.Println(bangunRuang.volume())
	fmt.Print("Luas Permukaan Balok : ")
	fmt.Println(bangunRuang.luasPermukaan())

	/*=== Soal 2 ===*/
	var samsung phoneDetails
	samsung = phone{
		name:   "X21",
		brand:  "Samsung",
		year:   2002,
		colors: []string{"black", "white", "red"},
	}
	var oppo phoneDetails
	oppo = phone{
		name:   "y21",
		brand:  "Oppo",
		year:   2001,
		colors: []string{"Platinum black", "Berry white", "red"},
	}

	samsung.printDetails()
	oppo.printDetails()

	//Soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	var k0 = prefix.(string)
	var k1 = kumpulanAngkaPertama.([]int)
	var k2 = kumpulanAngkaKedua.([]int)
	fmt.Println(k0 + strconv.Itoa(k1[0]) + " + " + strconv.Itoa(k1[1]) + " + " + strconv.Itoa(k2[0]) + " + " + strconv.Itoa(k2[1]) + " = " + strconv.Itoa(k1[0]+k1[1]+k2[0]+k2[1]))

	//Soal 5
	fmt.Println(Tambah(5, 6))
	fmt.Println(Kali(5, 6))

}

// Fungsi 3
func luasPersegi(x int, enable bool) interface{} {
	var temp interface{}
	if x != 0 && enable == true {
		temp = "Luas Persergi dengan Sisi " + strconv.Itoa(x) + " cm adalah " + strconv.Itoa(x*x) + " cm"
		return temp
	} else if x != 0 && enable == false {
		temp = "Sisi Persegi adalah " + strconv.Itoa(x) + " cm"
		return temp
	} else if x == 0 && enable == true {
		temp = "maaf anda belum menginputkan sisi persegi"
		return temp
	} else if x == 0 && enable == false {
		return nil
	}
	return nil
}
