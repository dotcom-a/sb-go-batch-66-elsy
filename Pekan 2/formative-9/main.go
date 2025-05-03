package main

import (
	"fmt"
	. "ninth-project/formative-8"
	. "ninth-project/formative-8/lib"
	"strconv"
)

func main() {
	//Soal 1
	var bandunDatar HitungBangunDatar
	bandunDatar = SegitigaSamaSisi{3, 3}
	fmt.Print("Luas Segitiga : ")
	fmt.Println(bandunDatar.Luas())
	fmt.Print("Keliling Segitiga : ")
	fmt.Println(bandunDatar.Keliling())

	bandunDatar = PersegiPanjang{4, 8}
	fmt.Print("Luas Persegi Panjang : ")
	fmt.Println(bandunDatar.Luas())
	fmt.Print("Keliling Persegi Panjang : ")
	fmt.Println(bandunDatar.Keliling())

	var bangunRuang HitungBangunRuang
	bangunRuang = Tabung{5, 10}
	fmt.Print("Volume Tabung : ")
	fmt.Println(bangunRuang.Volume())
	fmt.Print("Luas Permukaan Tabung : ")
	fmt.Println(bangunRuang.LuasPermukaan())

	bangunRuang = Balok{3, 3, 3}
	fmt.Print("Volume Balok : ")
	fmt.Println(bangunRuang.Volume())
	fmt.Print("Luas Permukaan Balok : ")
	fmt.Println(bangunRuang.LuasPermukaan())

	/*=== Soal 2 ===*/
	var samsung PhoneDetails
	samsung = Phone{
		Name:   "X21",
		Brand:  "Samsung",
		Year:   2002,
		Colors: []string{"black", "white", "red"},
	}
	var oppo PhoneDetails
	oppo = Phone{
		Name:   "y21",
		Brand:  "Oppo",
		Year:   2001,
		Colors: []string{"Platinum black", "Berry white", "red"},
	}

	samsung.PrintDetails()
	oppo.PrintDetails()

	//Soal 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

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
