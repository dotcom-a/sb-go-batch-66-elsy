package main

import (
	"fmt"
	"slices"
	"strconv"
	"sync"
	"time"
)

func main() {
	angka := 1

	//Soal 1
	defer deferFunc("Golang Backend Development", 2021)

	//Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//Soal 3
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//Soal 4
	var phones = []string{}
	addData(&phones, "Xiaomi")
	addData(&phones, "Asus")
	addData(&phones, "IPhone")
	addData(&phones, "Samsung")
	addData(&phones, "Oppo")
	addData(&phones, "Realmi")
	addData(&phones, "Vivo")

	for i := 0; i < len(phones); i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(strconv.Itoa(i+1) + ". " + phones[i])
	}

	//Soal 5
	var wg sync.WaitGroup
	wg.Add(1)
	go displayProduct(&wg)
	wg.Wait()

}

// Soal 1
func deferFunc(kalimat string, year int) {
	fmt.Println(kalimat + " " + strconv.Itoa(year))
}

// Soal 2
func kelilingSegitigaSamaSisi(x int, enable bool) string {
	var temp string
	if x != 0 && enable == true {
		temp = "Keliling Segitiga Sama Sisi dengan Sisi " + strconv.Itoa(x) + " cm adalah " + strconv.Itoa(3*x) + " cm"
		return temp
	} else if x != 0 && enable == false {
		temp = "Sisi Segitiga adalah " + strconv.Itoa(x) + " cm"
		return temp
	} else if x == 0 && enable == true {
		temp = "maaf anda belum menginputkan sisi dari segitiga sama sisi"
		return temp
	} else if x == 0 && enable == false {
		defer logError()
		temp = "maaf anda belum menginputkan sisi dari segitiga sama sisi"
		panic(temp)
	}
	return ""
}

func logError() {
	message := recover()
	fmt.Println("Error : ", message)
}

// Soal 3
func tambahAngka(x int, total *int) {
	(*total) += x
}

func cetakAngka(total *int) {
	fmt.Println("Total : " + strconv.Itoa(*total))
}

// Soal 4
func addData(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

// Soal 5
func displayProduct(wg *sync.WaitGroup) {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	slices.Sort(phones)

	for i := 0; i < len(phones); i++ {
		fmt.Println(strconv.Itoa(i+1) + ". " + phones[i])
	}

	wg.Done()
}
