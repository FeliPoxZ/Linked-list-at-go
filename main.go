package main

import (
	"fmt"
	List "modelo/impl"
)

func main() {
	// var number int = 17
	// pointer := &number

	// fmt.Println(&number)
	// fmt.Println(number)
	// fmt.Println(pointer)
	// fmt.Println(*pointer)
	// fmt.Println(&pointer)

	var lista *List.List
	fmt.Println(lista)

	lista = List.Create()

	List.Append(lista, 5)
	List.Append(lista, 7)
	List.Append(lista, 9)
	List.Append(lista, 4)
	fmt.Println(lista)
	List.Display(lista)

	value := List.Get(lista, 2)
	fmt.Println("valor retornado ->", value)

	List.Set(lista, 4, 12)
	List.Display(lista)

}
