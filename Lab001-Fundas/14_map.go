//!! Fungsi-fungsi yang ada di map
//?? len(map)						: Untuk mendapatkan jumlah data di map
//?? map[key]						: Mengambil data di map dengan key
//?? map[key] = value				: Mengubah data di map dengan key
//?? make(map[TypeKey]TypeValue)	: Membuat map baru
//?? delete(map, key)				: Menghapus data di map dengan key

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Bimo",
		"address": "Jatijajar",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//** beda penulisan deklarasi aja
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Bimo"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
