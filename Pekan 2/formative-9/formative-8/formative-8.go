package formative8

import (
	"fmt"
	"math"
	"strconv"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (s PersegiPanjang) Luas() int {
	return s.Lebar * s.Panjang
}

func (s PersegiPanjang) Keliling() int {
	return (s.Lebar + s.Panjang) * 2
}

func (s Tabung) Volume() float64 {
	return math.Pi * s.Tinggi * math.Pow(s.JariJari, 2)
}

func (s Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * s.JariJari * (s.Tinggi + s.JariJari)
}

func (s Balok) Volume() float64 {
	return float64(s.Lebar) * float64(s.Panjang) * float64(s.Tinggi)
}

func (s Balok) LuasPermukaan() float64 {
	return float64(2 * ((s.Panjang * s.Lebar) + (s.Panjang * s.Tinggi) + (s.Tinggi * s.Lebar)))
}

/* ===  struct 2 ===*/
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneDetails interface {
	PrintDetails()
}

func (s Phone) PrintDetails() {
	fmt.Println("Name : ", s.Name)
	fmt.Println("Brand : ", s.Brand)
	fmt.Println("Year : ", strconv.Itoa(s.Year))
	fmt.Println("Colors : ", s.Colors)

}

// Struct 3
type Persegi struct {
	Sisi int
}

// Fungsi 3
func LuasPersegi(x int, enable bool) interface{} {
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
