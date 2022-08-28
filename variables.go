package main

import "fmt"

func main() {
	//variables
	//string
	var name string = "rodrigo"

	//numericos
	var age int8 = 1                 // (-128 to 127)
	var number_normal int16 = 30000  // (-32768 to 32767)
	var number_large int32 = 3000000 // (-2147483648 to 2147483647)
	var big_number int64 = 100000000

	//flotantes
	var float_32 float32 = 32.0000
	var float_64 float64 = 32.0000000

	//booleanos
	var is_true bool = true
	var is_false bool = false

	fmt.Println(name, age, number_large, number_normal, big_number)
	fmt.Println(float_32, float_64)
	fmt.Println(is_false, is_true)

	//asignacion de variables y su tipo de manera implicita
	name_user := "lucas" //asigna el tipo string al momento de compilado y antes del mismo

	fmt.Println(name_user)

	//también tenemos variables de tipo complejo pero es más uso para matematicas

	const PI float32 = 3.14 //constantes, no se pueden reasignar su valor
	fmt.Println(PI)

}
