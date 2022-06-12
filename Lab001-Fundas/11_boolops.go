package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
