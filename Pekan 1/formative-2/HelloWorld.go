package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var text1 = "Bootcamp "
	var text2 = "Digital "
	var text3 = "Skill "
	var text4 = "Sanbercode "
	var text5 = "Golang"

	//Soal 1
	fmt.Println(text1 + text2 + text3 + text4 + text5)

	//Soal 2
	halo := "Halo Dunia"
	var newhalo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(newhalo)

	//Soal 3
	var kataPertama = "saya "
	var kataKedua = "senang "
	var kataKetiga = "belajar "
	var kataKeempat = "golang"

	// kataKedua: capitalize first letter
	word2 := strings.Title(strings.ToLower(kataKedua))

	// kataKetiga: capitalize last letter
	word3 := strings.ToLower(kataKetiga[:len(kataKetiga)-1]) + strings.ToUpper(string(kataKetiga[len(kataKetiga)-1]))

	word4 := strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama + word2 + word3 + word4)

	//Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num1, err1 = strconv.Atoi(angkaPertama)
	var num2, err2 = strconv.Atoi(angkaKedua)
	var num3, err3 = strconv.Atoi(angkaKetiga)
	var num4, err4 = strconv.Atoi(angkaKeempat)

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		fmt.Println(num1 + num2 + num3 + num4)
	}

	//Soal 5
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangInt, err1 := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangInt, err2 := strconv.Atoi(lebarPersegiPanjang)

	alasSegitigaInt, err3 := strconv.Atoi(alasSegitiga)
	tinggiSegitigaInt, err4 := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		luasPersegiPanjang = panjangPersegiPanjangInt * lebarPersegiPanjangInt
		kelilingPersegiPanjang = 2*panjangPersegiPanjangInt + 2*lebarPersegiPanjangInt
		luasSegitiga = (alasSegitigaInt * tinggiSegitigaInt) / 2

		fmt.Println(luasPersegiPanjang)
		fmt.Println(kelilingPersegiPanjang)
		fmt.Println(luasSegitiga)
	}

}
