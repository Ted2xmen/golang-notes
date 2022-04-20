package main

import "fmt"

func main() {

	var name string = "Ted"
	// or
	var surName string
	surName = "dogru"
	// or

	// or := type atamıs oluyorsun,

	age := 40 // int tanımlamana gerek kalmıyo

	var first, last string = "TUGRUL", "ERDEM"
	fmt.Println(first, last)

	var isActive bool
	isActive = true

	var (
		isClient bool   = false
		tools    string = "Tool 1, tool 2" // coklu atama var parantezi ile
	)

	var isMarried, weight, height = false, 80, 190 // coklu atama
	// or
	// isMarried, weight, height := true, 80, 190 // birden cok değişkene değer ataması.

	// temel veri tipleri
	fmt.Printf("%T\n", age) // veri tipini gösterir. int verir

	// zero values. string için "", sayısal ifadeler için 0. eğer değişkene değer vermezsen numerik değerlerin zero values'i sıfır.
	// zero values, boolean için  = false olur

	fmt.Println("Hello", name+surName)
	fmt.Println(age)
	fmt.Println(isClient)
	fmt.Println(isActive)
	fmt.Println(tools)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

}
