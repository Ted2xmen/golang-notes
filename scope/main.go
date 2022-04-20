package main

import "fmt"

var packVar = "package scope" // paketin içindeki tüm fonksiyonlar içinden ulasılabilir
// kısa değişken gösterme metodu paket değişkenlerinde kullanılamaz. anahtar kelimelerin belirgin olması sebebiyle..

func main() {

	if true {
		var blockVar = "block scope"
		fmt.Println(blockVar) // hata yok.
	}

	var funcVar = "Func scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)
	//	fmt.Println(blockVar) // hata: cünkü kendi scopu'unda tanımlı.
	myFunc()
}

func myFunc() {
	fmt.Println(packVar)
}
