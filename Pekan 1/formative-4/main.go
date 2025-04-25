package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Soal 1
	for i := 0; i < 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println("#")
	}

	//Soal 2
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Print("[ ")
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] != "aku" && kalimat[i] != "dan" {
			fmt.Print(kalimat[i] + " ")
		}
	}
	fmt.Print("]")

	//Soal 3
	var sayuran = [7]string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}
	for i := 0; i < len(sayuran); i++ {
		fmt.Println(strconv.Itoa(i) + ". " + sayuran[i])
	}

	//Soal 4
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Println(key + " = " + strconv.Itoa(value))
	}
	var volume = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Println("Volume Balok = " + strconv.Itoa(volume))

	//Soal 5
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume = volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

}
func luasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}
func kelilingPersegiPanjang(panjang int, lebar int) int {
	keliling := 2 * (panjang + lebar)
	return keliling
}
func volumeBalok(panjang int, lebar int, tinggi int) int {
	volume := panjang * lebar * tinggi
	return volume
}
