package main

// type conversion
import "fmt"

func main() {

	x := 10
	y := 10.0

	// int ve float64 işlem yapmaya çalışırsan hata : mismatched types
	// fmt.Println(x + y)

	// hataya engel olmak için birini birine çevirmek gerekiyor.
	fmt.Println(x + int(y))

	// yalnız bu değişim kısa süreliğine gerçekleşiyor. bu işlem değişkenin asıl değerini değiştirmiyor. sadece gerekli işlem gerçekleşene kadar.

}
