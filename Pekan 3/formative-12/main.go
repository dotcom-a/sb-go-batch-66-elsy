package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// getNilai func
func GetNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilai := nilaiNilaiMahasiswa
		dataNilai, err := json.Marshal(nilai)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// Post nilai
func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nilaiMahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&nilaiNilaiMahasiswa); err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(nilaiNilaiMahasiswa); i++ {
				nilaiNilaiMahasiswa[i].IndeksNilai = getIndeksNilai(int(nilaiNilaiMahasiswa[i].Nilai))
			}
		} else {
			// parse dari form
			nama := r.PostFormValue("nama")
			mata_kuliah := r.PostFormValue("mata_kuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			indeks_nilai := getIndeksNilai(nilai)

			nilaiMahasiswa = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mata_kuliah,
				IndeksNilai: indeks_nilai,
				Nilai:       uint(nilai),
			}
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
		}
		dataMahasiswa, _ := json.Marshal(nilaiNilaiMahasiswa) // to byte
		w.Write(dataMahasiswa)                                // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func main() {

	http.HandleFunc("/post_nilai", PostNilai)
	http.HandleFunc("/get_nilai", GetNilai)
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getIndeksNilai(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai < 80 && nilai >= 70 {
		return "B"
	} else if nilai < 70 && nilai >= 60 {
		return "C"
	} else if nilai < 60 && nilai >= 50 {
		return "D"
	} else if nilai < 50 {
		return "E"
	} else {
		//ignore
		return ""
	}

}
