package webserver

import (
	"fmt"
	"net/http"
)

func VolumeTabung(r float64, t float64) float64 {
	return 3.14 * r * r * t
}

func LuasTabung(r float64) float64 {
	return 3.14 * r * r
}

func KelilingTabung(r float64) float64 {
	return 2 * 3.14 * r
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w)
}

func StartServer() {
	jariJari := 7.0
	tinggi := 10.0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w)
		fmt.Println("Jari-jari : ", jariJari)
		fmt.Println("Tinggi : ", tinggi)
		fmt.Println("Volume : ", VolumeTabung(jariJari, tinggi))
		fmt.Println("Luas : ", LuasTabung(jariJari))
		fmt.Println("Keliling : ", KelilingTabung(jariJari))
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
