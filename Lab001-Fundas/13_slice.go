//!! Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
//?? Pointer adalah penunjuk data pertama di array para slice
//?? Length adalah panjang dari slice, dan
//?? Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

//!! Fungsi-fungsi pada slice
//?? array[low:high] 	: Membuat slice dari array dimulai index low sampai index sebelum high
//?? array[low:]		: Membuat slide dari array dimulai index low sampai index akhir di array
//?? array[:high]		: Membuat slice dari array dimulai index 0 sampai index sebelum high
//?? array[:]			: Membuat slice dari array dimulai index 0 sampai index akhir di array

package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bimo")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
	fmt.Println("=====================================================================")

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Sinatrio"
	newSlice[1] = "Bimo"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("=====================================================================")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println("=====================================================================")

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
